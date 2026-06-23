---
name: dora-api
description: Interact with Dora Network API and multiplexed WebSocket (wsplex) using environment-configured endpoints
version: 1.10.0
author: Dora Network
platforms: [linux, macos, windows]
---

# Dora Network API Skill

This skill provides guidance on interacting with the Dora Network API using environment-specific base URLs and API keys.

## Configuration

### Authentication

The Dora Network API supports two auth schemes. Both are accepted on any endpoint that requires authentication:

**API Key** — `Authorization: ApiKey <key>`
- Works for: all authenticated endpoints (`/v1/trades`, `/v1/assets`, `/v1/orderbooks/*`, `/v1/charts/*`, `/v1/user`, `/v1/pl/self`, `/v1/ledger/*`, `/v1/realized_pnl_settlements`)
- Set via `DORA_*_API_KEY` env vars

**JWT Bearer** — `Authorization: Bearer <token>`
- Accepted on all authenticated endpoints as an alternative to API Key
- ⚠️ **Not currently available to integration partners**

Set environment variables in your environment for each environment:
- `DORA_STAGING_BASE_URL` and `DORA_STAGING_API_KEY` for staging
- `DORA_PROD_BASE_URL` and `DORA_PROD_API_KEY` for production

Example:
```
DORA_STAGING_BASE_URL=https://staging.dora.co/
DORA_STAGING_API_KEY=your_dora_staging_key
```

These environment variables **override** any base URLs defined in the OpenAPI spec. If the `*_BASE_URL` variables are not set, the base URLs from the OpenAPI spec are used as defaults.

### WebSocket Endpoints
WebSocket endpoints are derived from the base URL by replacing `https://` with `wss://`.
For WebSocket authentication, pass the API key as the `x-api-key` query parameter on the endpoint path.
For example, if your base URL is `$DORA_BASE_URL`:
- REST API: `$DORA_BASE_URL`
- WebSocket API: `${DORA_BASE_URL/https:/wss:}` (replace https with wss)
- WebSocket with auth: `${DORA_BASE_URL/https:/wss:}/v1/orderbooks/{id}/stream?x-api-key=$DORA_API_KEY`

## Usage Pattern

Before making API calls, set your environment variables:
```bash
# For staging
export DORA_BASE_URL=$DORA_STAGING_BASE_URL
export DORA_API_KEY=$DORA_STAGING_API_KEY

# For production
export DORA_BASE_URL=$DORA_PROD_BASE_URL
export DORA_API_KEY=$DORA_PROD_API_KEY
```

Or use them directly in commands:
```bash
curl -L -H "Authorization: ApiKey $DORA_STAGING_API_KEY" "$DORA_STAGING_BASE_URL/v1/assets?limit=10"
```

## Common Operations

### List Assets
Get a paginated list of assets with filtering options.

```bash
# List assets (first 100 assets)
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets?page=1&limit=100"

# Filter by asset kind (BOND, CURRENCY, INTEREST, POOL_SHARE)
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets?asset_kind=BOND&limit=50"

# Filter by date range
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets?created_after=2024-01-01T00:00:00Z&created_before=2024-12-31T23:59:59Z"

# Filter by capabilities
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets?can_trade=true&can_add_liquidity=true"
```

### Get Asset Details
Retrieve details for a specific asset by its UUID.

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets/123e4567-e89b-12d3-a456-426614174000"
```

### Get Asset YTM (Yield to Maturity)
Get annualized yield to maturity for a bond asset.

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets/123e4567-e89b-12d3-a456-426614174000/ytm"
```

### Get Coupon Payments
Get coupon payments for a bond asset.

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/assets/123e4567-e89b-12d3-a456-426614174000/coupon_payments"
```

### User Management

List users (admin only):

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/user?limit=50&offset=0"
```

Get specific user:

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/user/123e4567-e89b-12d3-a456-426614174000"
```

### Market Data
Get orderbook (L3 market depth) data:

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/orderbooks/123e4567-e89b-12d3-a456-426614174000/orders"
```

Get candlestick data for an orderbook (resolution: `1m`, `5m`, `15m`, `1h`, `4h`, `1d` — lowercase only):

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/charts/123e4567-e89b-12d3-a456-426614174000/candle?start=2024-01-01T00:00:00Z&end=2024-12-31T23:59:59Z&resolution=1h"
```

Get yield chart data for an asset:

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/charts/123e4567-e89b-12d3-a456-426614174000/yield?start=2024-01-01T00:00:00Z&end=2024-12-31T23:59:59Z&resolution=1d"
```

### Trade Data

Fetch paginated trades for a user+orderbook+date range. Each trade includes price, quantity, side, aggressor indicator, and timestamps.

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/trades?user_ids=USER_UUID&order_book_ids=ORDERBOOK_UUID&start=2026-05-26T00:00:00Z&end=2026-05-27T00:00:00Z&page=1&limit=100"
```

**Pagination**: iterate `page=N` until response has fewer than `limit` items. Max per-page is 100.

**Trade object fields**:
- `transaction_id` — UUID  
- `created_at` — ISO timestamp  
- `price`, `quantity_0` — string, parse as float  
- `side` — `BUY` or `SELL`  
- `aggressor_indicator` — bool (true=taker, false=maker)  
- `order_id` — unique per trade (each leg gets its own ID)  
- `order_book_id`, `user_id`, `asset_0` — entity UUIDs  
- `fee_quantity` — fee charged ("0" if fees settled separately)  

### Orderbook Stats

```bash
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
     "$DORA_BASE_URL/v1/orderbooks/019c3420-5cd7-7a88-8fe6-a5a622e01ad9/stats"
```

Returns current price, 24h high/low, change %, volume base and USD.

## WebSocket Connections

### Orderbook stats stream
```bash
# Derive WS URL from base URL
WS_URL="${DORA_BASE_URL/https:/wss:}"
wscat -c "$WS_URL/v1/orderbooks/123e4567-e89b-12d3-a456-426614174000/stats/stream?x-api-key=$DORA_API_KEY" \
     -H "User-Agent: DoraAPIClient/1.0"
```

### Multiplexed WebSocket (`wsplex`) — preferred for new integrations

DORA's multiplexed WebSocket protocol, `wsplex`, lets a single connection carry requests, responses, and server-pushed notifications for multiple endpoints (currently `/prices` and `/trades`) over one socket. Prefer it over the legacy per-stream endpoints above for new integrations.

**Endpoint:** `wss://<environment_base_url>/plex` (e.g. `wss://staging.dora.co/plex`).

**Auth:** the API key is sent as the standard HTTP header on the WebSocket upgrade request:
```
Authorization: ApiKey <key>
```
`Authorization: Bearer <token>` works identically.

**Protocol shape (one connection, many paths):**
- Request: `{"id": "<uuidv7>", "path": "/prices", "data": {...}}`. The `data` field is **required** — omitting it returns an error response and still consumes the request id.
- Response: `{"id": ..., "kind": "response", "path": ..., "data": ...}` or `{"id": ..., "kind": "response", "path": ..., "error": "..."}`. Exactly one response per request, matched by `id`; an error response keeps the same envelope shape and swaps `data` for `error`.
- Notification (server-pushed): `{"id": ..., "kind": "notification", "path": ..., "data": ...}`.
- Request ids are **single-use per connection** — generate a fresh UUIDv7 for every request, including retries. Reusing an id returns a duplicate-request error.

**Quick test with `wscat`:**
```bash
# Subscribe to /prices for a specific asset
WS_URL="${DORA_BASE_URL/https:/wss:}"
wscat -c "$WS_URL/plex" \
    -H "Authorization: ApiKey $DORA_API_KEY" \
    -x '{"id":"019ed20f-cfcb-7167-a318-4b42d0582517","path":"/prices","data":{"subscribe":["<asset-id>"]}}'
```
You should receive one response (matching `id`) followed by a stream of notifications on `/prices`.

**Subscriptions on `/prices` follow a two-state model:**
- A **subscribed list** of asset ids, mutated by `subscribe` (add) and `unsubscribe` (remove).
- A `subscribed_to_all` **bypass flag** — when true, the list is ignored and every asset's prices stream.
The list and the bypass flag are independent: `subscribe` mutates the list regardless of the flag, and `unsubscribe` is only allowed when the flag is `false` (calling it in `subscribed_to_all` mode returns an error response). `unsubscribed_to_all` clears the flag (not the list). To fully tear down a `/prices` subscription, send an explicit `unsubscribe` for every id you originally added — but only while `subscribed_to_all` is `false`.

**Subscriptions on `/trades` follow the same two-state model per axis** (order-book and user). The all-mode toggles (`order_books_all`, `users_all`) are **connection-scoped** — once set to `true` they remain `true` for the lifetime of the connection. To fully clear `/trades` subscription state, close the connection.

**Runnable examples in Go, Python, and TypeScript** are in the repo at `multiplex-websocket/examples/{go,python,typescript}/` (each has its own README). Prefer these over rolling your own client.

For the full protocol reference, see `multiplex-websocket/README.md` in this repo.

## Error Handling
The API returns standard HTTP status codes:
- 200: Success
- 201: Created
- 204: No Content
- 400: Bad Request (validation errors)
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 429: Rate Limited
- 500: Internal Server Error
- 503: Service Unavailable

Error responses follow the `ResponseEnvelope` schema.

## Pagination
List endpoints use page-based pagination:
- `page`: Page number (starts at 1)
- `limit`: Items per page (default 100, check spec for max)

Check response metadata for total pages/items if available.

## Rate Limiting
Be aware of rate limits. If you receive 429 responses, implement exponential backoff and retry-after handling.

## Using with Terminal
You can interact with the Dora Network API using curl from your terminal:

```bash
# List assets
curl -L -s -H "Authorization: ApiKey $DORA_API_KEY" "$DORA_BASE_URL/v1/assets?limit=10"

# Get asset details
curl -L -s -H "Authorization: ApiKey $DORA_API_KEY" "$DORA_BASE_URL/v1/assets/$(uuidgen)"
```

## Reference

### OpenAPI Specification
The full OpenAPI specification is bundled as a linked file within this skill. Load it with:

```
skill_view(name='dora-api', file_path='references/openapi.json')
```

### Exploring the Spec
To explore the spec manually using `jq`, first resolve the skill directory path via `skill_view(name='dora-api')` (which returns `skill_dir`), then:

```bash
# List all operation IDs
jq -r '.paths[] | .[] | .operationId' <skill_dir>/references/openapi.json | sort

# Get schema for a specific component
jq '.components.schemas.UserCreatedResponse' <skill_dir>/references/openapi.json
```

Replace `<skill_dir>` with the resolved path from `skill_view`.

### Additional References
- `skill_view(name='dora-api', file_path='references/openapi-exploration.md')` — guided exploration notes
- `skill_view(name='dora-api', file_path='references/trade-analysis.md')` — methodology for investigating blowout/drawdown trading days
