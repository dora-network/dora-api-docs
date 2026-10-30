# Auto-rebalance API reference

GET routes require the owning user's key and remain self-scoped. PUT requires Admin or Integrator access and can target a CFD user with `user_id`; an Integrator is restricted to its own tenant. Authenticate with `Authorization: ApiKey <key>`. Decimal values are JSON strings.

See [configuration](./configure-auto-rebalance.md) and [behavior](./how-auto-rebalance-works.md).

## PUT /v2/accounts/auto_rebalance/config

Create or update a batch of isolated-account settings atomically. Example:

```json
{
  "user_id": "11111111-1111-4111-8111-111111111111",
  "configs": [{
    "account_id": "22222222-2222-4222-8222-222222222222",
    "enabled": true,
    "order_rebalance_enabled": true,
    "lowest_threshold": "0.30",
    "min_utilization": "0.60",
    "max_utilization": "0.80"
  }]
}
```

| Field | Required | Meaning |
|---|---|---|
| `user_id` | No | Target user; defaults to the authenticated operator. |
| `configs` | Yes | Non-empty list with unique account IDs. |
| `configs[].account_id` | Yes | Isolated account owned by the selected user. |
| `configs[].enabled` | Yes | Existing risk-rebalance participation; required for funding/returns too. |
| `configs[].order_rebalance_enabled` | No | Shared funding and auto-return opt-in. Inserts default false; omitted updates preserve the stored value. |
| `configs[].lowest_threshold` | Conditional | Required in the effective enabled shared opt-in. Must be positive and below `min_utilization`; omission retains a stored threshold. |
| `configs[].min_utilization` | Yes | Risk restore target/donor floor; positive and below `max_utilization`. |
| `configs[].max_utilization` | Yes | Risk trigger threshold; at most 1. |

There is no independent return toggle. With the shared opt-in off, the client manages its isolated balances. Both funding and returns also require DORA's environment rollout gate. Resting LIMIT orders prevent activation.

`200` returns `data.configs` with persisted settings. Errors use the normal `error` and `metadata.status_code` envelope:

| Status | Meaning |
|---|---|
| `400` | Invalid body, missing required fields, duplicate IDs, invalid thresholds, missing return threshold for an enabled shared opt-in, or resting limits preventing activation. |
| `401` | Missing or invalid authentication. |
| `403` | Writer lacks Admin/Integrator access, or enabling configuration for a non-CFD user. |
| `404` | Missing/out-of-tenant target, or an account not found, not owned by that target, or not isolated. |

## GET /v2/accounts/auto_rebalance/config

Return the authenticated user's persisted settings in `data.configs`, including disabled configurations. No live balance/capacity fields are added to this configuration-only list. An empty list means nothing is configured.

## GET /v2/accounts/auto_rebalance/config/{account_id}

Return one self-owned configured isolated account's settings and status in `data`. An invalid UUID gives `400`; no matching owned configuration gives `404`.

## GET /v2/accounts/auto_rebalance/status

Return settings and status for all of the authenticated user's configured isolated accounts in `data.configs`, including disabled accounts.

## Status fields

| Field | Meaning |
|---|---|
| `account_id`, `enabled`, `order_rebalance_enabled` | Persisted account ID and settings. |
| `lowest_threshold` | Return trigger when configured. Zero effective debt returns all free USD. |
| `min_utilization`, `max_utilization` | Risk restore/donor threshold and risk trigger threshold. |
| `available_usd` | Available USD in the isolated account. |
| `borrowed_value`, `borrow_limit` | Effective USD debt and weighted collateral limit. |
| `current_utilization` | Debt divided by borrow limit; null/omitted when undefined. An empty account can still be refilled; zero effective debt permits returning all free USD. |
| `above_max` | Present when utilization is defined; true above the risk trigger. |
| `in_active_liquidation` | Whether liquidation is active. |
| `last_rebalance_outcome` | Latest risk-rebalance result, excluding order funding and returns. |
| `order_funding_capacity_usd` | Optional read-only upper estimate of safe fundable USD. Omitted while inactive/ineligible; it reserves nothing and cannot include a specific future order's additional reserve. Submission replans. |

## Order submission

CFD users submit ordinary MARKET orders with explicit quantities. LIMIT and Max Buy requests are rejected in both single and bulk submission. Automatic funding applies to existing selected USD-quoted isolated accounts; global/new-account funding retains its existing behavior.

A funded `POST /v1/orders` response includes optional `data.funding`. Bulk responses include the same object on each funded result:

```json
{
  "amount_usd": "10.01",
  "sources": [{
    "account_id": "33333333-3333-4333-8333-333333333333",
    "amount_usd": "10.01"
  }]
}
```

Unfunded results omit `funding`. Each bulk item keeps its own transaction/outcome. Funding, audit rows and order insertion commit together or roll back together; failed orders leave no partial funding. Funding and routine returns send no new emails. Existing risk alerts keep their delivery settings.

An empty opted-in account keeps its ID and configuration. Later explicit-quantity orders can refill it if safe sources cover the shortage. Unsettled orders, liquidation, invalid debt prices and reserved/conditional cash prevent unsafe returns.
