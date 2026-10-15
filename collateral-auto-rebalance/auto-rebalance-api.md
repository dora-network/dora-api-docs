# Auto-rebalance API reference

Reference for the `/v2/accounts/auto_rebalance` routes. All routes require the API key of the owning user (`Authorization: ApiKey <key>`) and operate only on that user's own isolated margin accounts. For behavior and semantics see [How auto-rebalance works](./how-auto-rebalance-works.md); for a guided walkthrough see [Configure auto-rebalance](./configure-auto-rebalance.md).

All decimal values are JSON strings (for example `"0.60"`).

Success responses use the standard envelope:

```json
{ "data": { ... }, "metadata": { "status_code": 200, "trace_id": "...", "request_id": "..." } }
```

Errors use:

```json
{ "error": "human-readable message", "metadata": { "status_code": 400, "trace_id": "...", "request_id": "..." } }
```

---

## PUT /v2/accounts/auto_rebalance/config

Create or update auto-rebalance configuration for one or more isolated accounts. Missing configurations are inserted; existing ones are updated. Configuration changes take effect immediately.

### Request body

```json
{
  "configs": [
    {
      "account_id": "6f1d2b66-....-....-....-............",
      "enabled": true,
      "min_utilization": "0.60",
      "max_utilization": "0.80"
    }
  ]
}
```

| Field | Type | Required | Description |
|---|---|---|---|
| `configs` | array | yes | At least one configuration item; `account_id` values must be unique within the request |
| `configs[].account_id` | string (uuid) | yes | Isolated margin account belonging to the authenticated user |
| `configs[].enabled` | boolean | yes | Whether auto-rebalance is active for the account |
| `configs[].min_utilization` | string (decimal) | yes | Lower utilization threshold; must be `> 0` and `< max_utilization` |
| `configs[].max_utilization` | string (decimal) | yes | Upper utilization threshold; must be `> min_utilization` and `<= 1` |

### Response — `200 OK`

`data.configs`: array of persisted configurations (fields `account_id`, `enabled`, `min_utilization`, `max_utilization`).

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

### Errors

| Status | When |
|---|---|
| `400` | Malformed JSON; `configs` empty; missing or duplicated `account_id`; missing `enabled`; thresholds violate `0 < min_utilization < max_utilization <= 1` |
| `404` | Any `account_id` was not found, does not belong to the authenticated user, or is not an isolated account |
| `401` | Missing or invalid API key |

---

## GET /v2/accounts/auto_rebalance/config

List the persisted auto-rebalance configuration for the authenticated user's isolated accounts. Returns configuration fields only, without live status.

### Response — `200 OK`

`data.configs`: array as returned by `PUT` above (fields `account_id`, `enabled`, `min_utilization`, `max_utilization`). Empty array when nothing is configured.

### Errors

| Status | When |
|---|---|
| `401` | Missing or invalid API key |

---

## GET /v2/accounts/auto_rebalance/config/{account_id}

Fetch configuration plus live status for a single isolated account.

### Path parameters

| Field | Type | Description |
|---|---|---|
| `account_id` | string (uuid) | Isolated account ID with an existing auto-rebalance configuration |

### Response — `200 OK`

A single status object; fields are described in [Status object](#status-object) below.

```json
{
  "data": {
    "account_id": "6f1d2b66-....-....-....-............",
    "enabled": true,
    "min_utilization": "0.6",
    "max_utilization": "0.8",
    "available_usd": "100",
    "borrowed_value": "92",
    "borrow_limit": "100",
    "current_utilization": "0.92",
    "above_max": true,
    "in_active_liquidation": false,
    "last_rebalance_outcome": "partial"
  },
  "metadata": { "status_code": 200 }
}
```

### Errors

| Status | When |
|---|---|
| `400` | `account_id` path parameter is not a valid UUID |
| `404` | No auto-rebalance configuration exists for the account |
| `401` | Missing or invalid API key |

---

## GET /v2/accounts/auto_rebalance/status

List configuration plus live status for **all** of the authenticated user's accounts with an auto-rebalance configuration, including disabled ones.

### Response — `200 OK`

`data.configs`: array of [status objects](#status-object). Empty array when nothing is configured.

### Errors

| Status | When |
|---|---|
| `401` | Missing or invalid API key |

---

## Status object

| Field | Type | Description |
|---|---|---|
| `account_id` | string (uuid) | Isolated account the configuration applies to |
| `enabled` | boolean | Whether auto-rebalance is active for the account |
| `min_utilization` | string (decimal) | Lower threshold; restore target and donor floor |
| `max_utilization` | string (decimal) | Upper threshold; trigger line |
| `available_usd` | string (decimal) | Available USD balance in the account |
| `borrowed_value` | string (decimal) | Borrowed value in USD used by utilization math |
| `borrow_limit` | string (decimal) | Borrow limit in USD used by utilization math |
| `current_utilization` | string (decimal), nullable | `borrowed_value / borrow_limit`; `null` when `borrow_limit <= 0` (account is skipped by auto-rebalance) |
| `above_max` | boolean | `true` when `current_utilization > max_utilization` (present only when `current_utilization` is set) |
| `in_active_liquidation` | boolean | `true` when the account is in an active liquidation flow (such accounts are never rebalance targets) |
| `last_rebalance_outcome` | string | Most recent persisted attempt outcome: `completed`, `partial`, or `unsuccessful`; present only after the account has had at least one attempt |
