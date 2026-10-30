# Configure collateral auto-rebalance

An Admin or Integrator configures auto-rebalance for a CFD user. The client reads its own settings and status; it does not change configuration itself. See [how it works](./how-auto-rebalance-works.md) and the [API reference](./auto-rebalance-api.md).

## Prerequisites

- A DORA Admin has assigned the institutional user the `CFD` tag through the existing user-tag route.
- The operator has an Admin or Integrator API key. An Integrator can configure users in its own tenant; an Admin can configure authorized CFD users across tenants.
- The user has existing isolated accounts to select. A first order still uses the existing isolated-account creation flow.
- Resting LIMIT orders have been cancelled before enabling automatic funding and returns. CFD users submit ordinary MARKET orders with explicit quantities; LIMIT and Max Buy requests are rejected.
- DORA has activated the paired feature in the environment. Account settings do not override that rollout gate.

All decimal values are JSON strings.

## Find the isolated accounts

The CFD client can list its own accounts using its own key:

```bash
curl -L -H "Authorization: ApiKey $DORA_CLIENT_API_KEY" \
  "$DORA_BASE_URL/v2/user/self/accounts"
```

Give the operator the user ID and the isolated account IDs to configure.

## Enable risk rebalancing, order funding and returns

The operator sends a batch with the target `user_id`:

```bash
curl -L -X PUT -H "Authorization: ApiKey $DORA_OPERATOR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "11111111-1111-4111-8111-111111111111",
    "configs": [{
      "account_id": "22222222-2222-4222-8222-222222222222",
      "enabled": true,
      "order_rebalance_enabled": true,
      "lowest_threshold": "0.30",
      "min_utilization": "0.60",
      "max_utilization": "0.80"
    }]
  }' \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/config"
```

Replace the example IDs with the selected user's IDs. Omitting `user_id` defaults to the authenticated operator, not a selected customer.

`enabled` controls existing risk rebalancing. `order_rebalance_enabled` is one shared opt-in for order funding and auto-return: enabling it enables both; disabling it disables both. There is no separate return switch. An active shared opt-in requires:

```text
0 < lowest_threshold < min_utilization < max_utilization <= 1
```

`lowest_threshold` triggers returns; `max_utilization` triggers risk funding toward `min_utilization`. With zero effective debt, all free USD returns instead of targeting a positive utilization. Effective debt includes interest and applicable coupon liabilities.

The batch is atomic. Invalid thresholds or resting limits reject activation with `400`; unauthorized writers or enabling a non-CFD user receive `403`. A missing/out-of-tenant target, foreign account or non-isolated account receives `404`.

## Read status

Using the CFD user's own key:

```bash
curl -L -H "Authorization: ApiKey $DORA_CLIENT_API_KEY" \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/status"
```

For one account, use `/v2/accounts/auto_rebalance/config/{account_id}`. GET routes remain self-scoped even though the operator can target a customer on PUT.

When present, `order_funding_capacity_usd` is an upper estimate of safe donor funding. It reserves nothing and cannot include a particular future order's additional reserve requirement. Submission always recalculates. `last_rebalance_outcome` reports risk rebalancing only, excluding order-funding and return audit entries.

## Update or disable

Every item still requires `account_id`, `enabled`, `min_utilization` and `max_utilization`. A missing shared opt-in defaults to false on insert; omitted updates retain its stored value. Omitting `lowest_threshold` retains a stored threshold. An existing enabled shared opt-in cannot be updated into an invalid threshold configuration.

- Set `order_rebalance_enabled` to false to stop both order funding and returns while keeping risk rebalancing enabled. The client then manages its isolated balances itself.
- Set `enabled` to false to disable the account's participation in risk rebalancing and the new funding/return behaviors. Authorized operators can disable existing configuration after tag removal.
- Accounts omitted from the batch keep their settings.

## What the client should expect

An explicit-quantity MARKET order on an existing opted-in USD-quoted isolated account can receive funding from eligible isolated donors and then Global. Existing cash-reserve, leverage, borrowing and position checks still apply. Transfers and order insertion commit together or roll back together. Successful single and per-order bulk responses include optional `funding` amount/source details; unfunded results omit them.

Returns retain locked and conditional commitments and defer while accepted MARKET orders remain unsettled, liquidation is active, or debt cannot be priced safely. Incoming USD can replenish Global even when it is below its reserve. Returning all free USD keeps the isolated account ID and settings, so a later explicit-amount order can refill the same account.

Funding and routine returns add no emails. Existing risk-attempt notifications and daily reports keep their delivery settings. Risk rebalancing remains best effort; maintain sufficient safe USD and monitor unresolved risk outcomes.
