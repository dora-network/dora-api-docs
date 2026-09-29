# Configure collateral auto-rebalance for your accounts

This guide shows how to enable, tune, monitor, and disable collateral auto-rebalance on your isolated margin accounts through the DORA API. It assumes you have a DORA API key and are comfortable making authenticated REST calls. For what the service does once enabled, see [How auto-rebalance works](./how-auto-rebalance-works.md); for full endpoint details, see the [API reference](./auto-rebalance-api.md).

## Prerequisites

- A DORA API key for the user whose accounts you are managing. Authenticate every request with:

  ```bash
  curl -L -H "Authorization: ApiKey $DORA_API_KEY" "$DORA_BASE_URL/..."
  ```

- One or more **isolated margin accounts** for that user. Auto-rebalance applies to isolated accounts only; global cash accounts participate automatically as a last-resort funding source and are never configured directly.
- Threshold values in mind per account: `min_utilization` and `max_utilization` as decimal strings satisfying `0 < min_utilization < max_utilization <= 1` (for example `"0.60"` and `"0.80"`).

> NOTE: All decimal values in auto-rebalance requests and responses are JSON **strings**, not numbers.

## Step 1 — Find your isolated account IDs

List the user's accounts:

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
  "$DORA_BASE_URL/v2/user/self/accounts"
```

From the response, note the `account_id` of each **isolated** account you want to manage.

**Check:** you have one or more isolated account IDs.

## Step 2 — Enable auto-rebalance with thresholds

Upsert configurations for one or more accounts in a single batch request:

```bash
curl -L -X PUT -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
        "configs": [
          {
            "account_id": "6f1d2b66-....-....-....-............",
            "enabled": true,
            "min_utilization": "0.60",
            "max_utilization": "0.80"
          },
          {
            "account_id": "8a3c94d1-....-....-....-............",
            "enabled": true,
            "min_utilization": "0.50",
            "max_utilization": "0.75"
          }
        ]
      }' \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/config"
```

The same route creates missing configurations and updates existing ones. Each item in `configs` must set all four fields. The request is rejected (`400`) if:

- any `min_utilization` is not greater than 0, or not less than that account's `max_utilization`
- any `max_utilization` is greater than 1
- `enabled` is missing, an `account_id` is missing or duplicated, or `configs` is empty

and with `404` if any account does not exist, does not belong to you, or is not an isolated account.

**Check:** the response echoes each account's persisted configuration:

```json
{
  "data": {
    "configs": [
      {
        "account_id": "6f1d2b66-....-....-....-............",
        "enabled": true,
        "min_utilization": "0.6",
        "max_utilization": "0.8"
      }
    ]
  },
  "metadata": { "status_code": 200 }
}
```

From this point the service is live for these accounts: as soon as an account's utilization rises above its `max_utilization`, DORA's daemon evaluates and executes a rebalance. No restart or further action is needed — configuration changes take effect immediately, including for accounts enabled later.

## Step 3 — Verify status

Check configuration plus live utilization for all managed accounts (including disabled ones):

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/status"
```

or for a single account:

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/config/6f1d2b66-....-....-....-............"
```

**Check:** each entry shows `current_utilization` and `above_max`, so you can confirm whether an account is currently in breach. After the first rebalance attempt, `last_rebalance_outcome` shows `completed`, `partial`, or `unsuccessful`.

## Step 4 — Tune thresholds or disable an account

To change thresholds or turn auto-rebalance off for specific accounts, PUT a new batch — the upsert overwrites the previous settings for every `account_id` included:

```bash
curl -L -X PUT -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
        "configs": [
          {
            "account_id": "8a3c94d1-....-....-....-............",
            "enabled": false,
            "min_utilization": "0.50",
            "max_utilization": "0.75"
          }
        ]
      }' \
  "$DORA_BASE_URL/v2/accounts/auto_rebalance/config"
```

Setting `enabled` to `false` disables the account immediately for both roles — it stops being a rebalance target *and* stops donating to other accounts. Thresholds must remain valid even for disabled accounts. Accounts not included in the request keep their existing configuration.

## Expected result

Once enabled, no further calls are required to keep the feature running:

- every account above its `max_utilization` is automatically rebalanced back toward its `min_utilization` using your other opted-in accounts and — last — your global cash account
- every attempt (completed, partial, or unsuccessful) triggers an email notification with the transfer details
- accounts still above their maximum after the daily sweep receive a daily risk report email

> **IMPORTANT:** Auto-rebalance is a best-effort risk-management assist, not a guarantee against liquidation. Keep enough transferable USD across your opted-in and global accounts, and treat `partial`/`unsuccessful` outcomes and daily risk reports as signals to add funds.

## Choosing thresholds

- `max_utilization` is how much utilization you tolerate before wanting automatic intervention — set it low enough that a rebalance can complete before the account approaches liquidation territory.
- `min_utilization` is both the restore target and the floor protecting donor accounts — an account at or above its own min never donates.
- A common starting point is `min_utilization = "0.60"` / `max_utilization = "0.80"`; tighten the gap for more frequent, smaller transfers or widen it for fewer, larger ones.
