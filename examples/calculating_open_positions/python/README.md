# Calculating open positions — Python

Fetches the authenticated user's account/position portfolio from the DORA REST API and prints one line per (account, asset) entry with a non-zero `available` or `borrowed` balance — i.e. an open position.

## Setup

```bash
cd examples/calculating_open_positions
cp example.env .env
# edit .env: set DORA_BASE_URL and DORA_API_KEY
cd python
python3 calculate_positions.py
```

Uses only the Python standard library (`urllib`, `json`). A small `.env` reader is included so no third-party packages are required. If you prefer the full dotenv behaviour, install `python-dotenv` and replace the `load_dotenv` function with `from dotenv import load_dotenv; load_dotenv()`.

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

One row per (account, asset) entry where `available > 0` or `borrowed > 0`. The script then labels each row with `is_global` so you can see at a glance whether the position is in your main balance or in an isolated account.

A bond's `available` in an isolated account = long position; `borrowed` in an isolated account = short or quote-asset borrow. Aggregate `available` across the global account and any isolated accounts to get total long; aggregate `borrowed` across isolated accounts only to get total short.

If you have no open positions, the script prints `No open positions.`
