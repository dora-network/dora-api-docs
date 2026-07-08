# Calculating open positions — Go

Fetches the authenticated user's account/position portfolio from the DORA REST API and prints one line per (account, asset) entry with a non-zero `available` or `borrowed` balance — i.e. an open position.

## Setup

```bash
cd examples/calculating_open_positions
cp example.env .env
# edit .env: set DORA_BASE_URL and DORA_API_KEY
cd go
go mod tidy
go run .
```

Requires Go 1.22+. The only third-party dep is `github.com/joho/godotenv`, declared in `go.mod`.

## What it calls

`GET {DORA_BASE_URL}/v2/ledger/accounts/self` with `Authorization: ApiKey <key>`.

Response shape (per [`understanding-positions.md`](../../../understanding-positions.md)):

```json
{
  "data": {
    "portfolio": {
      "accounts": {
        "<account-uuid>": { "<asset-uuid>": AccountV2 }
      }
    }
  }
}
```

## Output

One row per (account, asset) entry where `available > 0` or `borrowed > 0`. Reads:

```
account_id                           asset_id       is_global       available         borrowed
------------------------------------------------------------------------------------------------
019c4d37-…                            019c3401…      true              5000.0              0.0
019e1234-…                            019c5678…      true               100.0              0.0
019e9abc-…                            019c5678…      false             1000.0              0.0
019e9abc-…                            019c9999…      false                0.0            400.0
```

A bond's `available` in an isolated account = long position; `borrowed` in an isolated account = short or quote-asset borrow. Aggregate `available` across the global account and any isolated accounts to get total long; aggregate `borrowed` across isolated accounts only to get total short.

If you have no open positions, the script prints `No open positions.`
