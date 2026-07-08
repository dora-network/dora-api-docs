# Obtaining Open Positions

DORA tracks every trader's balances in **accounts** — one **global** account (the main cash balance) plus zero or more **isolated** accounts (created automatically when a leveraged position is opened). An isolated account holds positions in exactly one bond asset and the pair's quote asset; the asset pair corresponds to the order book the leveraged order was placed on (e.g. an order on `GOOGL_5.25_2055-USD` lands in an isolated account with positions in both `GOOGL_5.25_2055` and `USD`). The **global** account cannot use leverage but can hold positions in any bond or currency asset.

> **Note:** the legacy endpoint `GET /v1/ledger/positions/self` is deprecated. Use the V2 endpoints described below. The non-multiplexed ledger stream at `/v1/user/{user_id}/ledger/stream` is the V1 streaming equivalent of the V2 REST snapshot; it is **not** deprecated.

## One-shot snapshot — `GET /v2/ledger/accounts/self`

Returns the full account/position portfolio for the authenticated user. No parameters required.

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v2/ledger/accounts/self"
```

Response shape (envelope elided, `data.portfolio` shown):

```json
{
  "user_id": "<user-uuid>",
  "accounts": {
    "<account-uuid>": {
      "<asset-uuid>": {
        "id": "<account-uuid>",
        "asset_id": "<asset-uuid>",
        "seq": 1,
        "is_global": true,
        "available": "5000.0",
        "locked": "0.0",
        "supplied": "2000.0",
        "borrowed": "0.0",
        "impending_borrows": "0.0",
        "avg_entry_price": "0.0",
        "borrow_limit": "0.0",
        "liquidation_threshold": "0.0",
        "created_at": "2026-01-01T00:00:00Z",
        "pending_withdrawal": "0.0",
        "account_name": "global"
      }
    }
  },
  "net_stablecoin_equivalence": {
    "gained": { "USD": "0.0" },
    "lost":   { "USD": "0.0" }
  }
}
```

- `accounts` is keyed first by `account_id`, then by `asset_id`. The OpenAPI schema declares both levels as unconstrained `additionalProperties`, so the spec itself does not dictate the key order — this matches the live response shape.

- `is_global: true` marks the single global account. `account_name` is `"global"` for the global account and a generated label for isolated accounts.
- `is_global: false` indicates an isolated account scoped to one bond/quote-asset pair.
- `available` is what you can trade with right now; `locked` is reserved by open orders; `supplied` is the balance lent to the leverage module; `borrowed` is the outstanding debt on the account. On a leveraged order inside an **isolated** account:
  - **Leveraged buy:** the quote asset (e.g. USD) borrowed to fund the purchase is reflected in `borrowed` on the quote-asset position; the bond asset purchased is reflected in full in `available` on the bond-asset position.
  - **Leveraged sell (short):** the bond asset borrowed to cover the short is reflected in `borrowed` on the bond-asset position; the quote asset received from the sale is added to `available` on the quote-asset position.
- For an at-a-glance list of accounts (id, name, global flag) without the per-asset position breakdown, use `GET /v2/user/self/accounts`.

## Real-time updates — `GET /v1/user/{user_id}/ledger/stream`

Opens a WebSocket that first replays the user's positions as a snapshot, then pushes one message per position change. Auth is via the `x-api-key` query parameter.

```bash
WS_URL="${DORA_BASE_URL/https:/wss:}"
wscat -c "$WS_URL/v1/user/<USER_UUID>/ledger/stream?x-api-key=$DORA_API_KEY" \
      -H "User-Agent: DoraAPIClient/1.0"
```

Each message is a JSON array of entries shaped like:

```json
[
  {
    "Val": {
      "id": "<position-uuid>",
      "asset_id": "<asset-uuid>",
      "seq": 1,
      "is_global": true,
      "available": "5000.0",
      "locked": "0.0",
      "supplied": "2000.0",
      "borrowed": "0.0",
      "impending_borrows": "0.0",
      "avg_entry_price": "0.0",
      "borrow_limit": "0.0",
      "liquidation_threshold": "0.0",
      "created_at": "2026-01-01T00:00:00Z",
      "position_name": "global",
      "pending_withdrawal": "0.0"
    },
    "Time": "2026-01-01T00:00:00Z"
  }
]
```

- The stream summary in the OpenAPI spec references fetching updates "since a specific time", but the `parameters` array only documents `user_id` and does not list a `since` query parameter. The spec is ambiguous on this point — treat the behaviour as **not confirmed by the spec** and verify empirically (e.g. by calling the endpoint with `?since=<ISO timestamp>`) before relying on it in production.
- The stream stays open for the life of the connection. On reconnect, re-apply your last-seen `Time` to a fresh `GET /v2/ledger/accounts/self` snapshot and discard any stream entries with `Time` older than that cutoff to de-duplicate.

## Related V2 endpoints

For working with accounts once you have the snapshot:

| Action | Endpoint |
|---|---|
| List accounts (id, name, global flag) | `GET /v2/user/self/accounts` |
| Get full account/position portfolio | `GET /v2/ledger/accounts/self` |
| Transfer balance between accounts | `POST /v2/accounts/transfer_balances` |
| Close a position (generates a market order) | `POST /v2/accounts/close` |

Replace any reference to the V1 endpoints `/v1/user/self/position_accounts`, `/v1/positions/transfer_balances`, or `/v1/positions/close` with the V2 equivalents above.

## Calculating open positions

An **open position** on a bond asset is a non-zero `available` balance (a long position, in either the global account or an isolated account) or a non-zero `borrowed` balance (a short position, in an isolated account only — the global account cannot be borrowed against). Iterate through the accounts returned by the snapshot, look at each per-asset entry, and aggregate the bond-asset `available` (across global + all isolated accounts) and bond-asset `borrowed` (across isolated accounts only) by asset. Quote-asset entries (e.g. USD, USDC) are funding balances, not positions.

The simple rule:

- The bond-asset `available` in an isolated account is the long position held in that account. If the user also has a non-zero `available` for the same asset in their global account, that is an **additional** long position on top of the isolated long.
- The bond-asset `borrowed` in an isolated account is the short position held in that account. (The global account never holds a borrowed bond position — borrowing is only possible inside an isolated account.)
- For the quote asset, `borrowed` inside an isolated account reflects currency borrowed to fund a leveraged buy; `available` after a leveraged sell is the proceeds held inside the isolated account. Neither is a "position" in the directional sense.

### Algorithm

1. `GET /v2/ledger/accounts/self`.
2. Walk `data.portfolio.accounts` — outer key `account_id`, inner key `asset_id`, value `AccountV2`.
3. For each (account, asset) entry, if `available == 0` and `borrowed == 0`, skip — no position there.
4. Otherwise, classify by the `is_global` flag on the parent account:
   - `is_global: true` → this is the main balance. Treat the entry as one source of truth for that asset's long (the global account cannot be borrowed against for a directional position).
   - `is_global: false` → isolated account scoped to one bond/quote pair. Aggregate `available` and `borrowed` for the bond leg of the pair to that account's contribution to the user's position.
5. Sum across all accounts by asset_id.
6. For each asset where the user has any non-zero `available` or `borrowed`, the open position is: `long = sum(available)`, `short = sum(borrowed, isolated only)`, `net = long - short`.

### Worked example

Consider a user with the following account state. The example numbers are synthetic (not from a live account) and chosen so the math is easy to verify by hand.

```json
{
  "data": {
    "portfolio": {
      "accounts": {
        "<global-acct>": {
          "<USD-asset>":   { "is_global": true,  "available": "5000", "borrowed": "0" },
          "<GOOGL-asset>": { "is_global": true,  "available": "100",  "borrowed": "0" }
        },
        "<isol-GOOGL-acct>": {
          "<USD-asset>":    { "is_global": false, "available": "250",  "borrowed": "750" },
          "<GOOGL-asset>":  { "is_global": false, "available": "1000", "borrowed": "0" }
        },
        "<isol-TSLA-acct>": {
          "<USD-asset>":   { "is_global": false, "available": "500",  "borrowed": "0" },
          "<TSLA-asset>":  { "is_global": false, "available": "0",    "borrowed": "2000" }
        }
      }
    }
  }
}
```

The leveraged-short account holds the user's own USD as collateral: the user supplies $500 USD into the isolated account and, with 4× leverage on that collateral, the system lends them $2000 worth of TSLA bonds (2000 bonds at $1 each) to short. The user does not borrow USD for the short — the USD leg is the user's own balance, sitting in the isolated account as collateral. The leveraged-long account uses the same model in mirror image: the user supplies $250 of their own USD, the system lends $750 USD to top it up, and the $1000 total buys 1000 GOOGL at $1 each.

Walking each non-zero entry:

| Account          | Asset     | is_global | available | borrowed | What it means                                                    |
|------------------|-----------|-----------|-----------|----------|------------------------------------------------------------------|
| global           | USD       | true      | 5000      | 0        | Cash balance — not a position.                                   |
| global           | GOOGL     | true      | 100       | 0        | Long 100 GOOGL held in the main balance.                         |
| isolated GOOGL   | USD       | false     | 250       | 750      | User's $250 USD + $750 system-lent (4× leverage on the collateral), used to fund the leveraged long. |
| isolated GOOGL   | GOOGL     | false     | 1000      | 0        | Long 1000 GOOGL held in the isolated account.                    |
| isolated TSLA    | USD       | false     | 500       | 0        | User-supplied USD collateral in the isolated account.            |
| isolated TSLA    | TSLA      | false     | 0         | 2000     | Short 2000 TSLA (borrowed via 4× leverage on the $500 collateral). |

Aggregating per bond asset:

| Asset  | Long (sum available)               | Short (sum borrowed, isolated only) | Net position |
|--------|------------------------------------|-------------------------------------|--------------|
| GOOGL  | 100 (global) + 1000 (isolated) = 1100 | 0                                   | **long 1100**  |
| TSLA   | 0                                  | 2000 (isolated)                     | **short 2000**   |
These are tracked as separate per-account entries in the API response; the combined net is derived by summing across accounts.

When placing an order, DORA validates the balance of only **one** account at a time — the global account or a specific isolated account. Balances are not pooled across accounts. For example, with 100 GOOGL in the global account and 1000 GOOGL in an isolated account, a single sell of 1100 GOOGL will be rejected. Instead, create separate orders via `POST /v1/orders`, setting the `from_global_position` field to `true` for the global account or `false` for an isolated account.

### Runnable examples

Working scripts in Go, Python, and TypeScript live at [`examples/calculating_open_positions/`](./examples/calculating_open_positions/). Each one fetches the snapshot, walks the accounts, and prints one row per (account, asset) entry with a non-zero balance — the same iteration this section describes, just done for you. Copy [`examples/calculating_open_positions/example.env`](./examples/calculating_open_positions/example.env) to `.env` inside the example folder, set `DORA_BASE_URL` and `DORA_API_KEY`, then follow the per-language README to run it.
