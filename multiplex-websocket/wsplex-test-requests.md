# wsplex Manual Test Requests

Example subscribe/unsubscribe requests for every documented wsplex path.
All entity IDs are placeholders — replace them with real IDs from your
environment before sending. See [Fetching entity IDs](#fetching-entity-ids)
below for example API requests to obtain them.

## Connection

```
wss://<DORA_BASE_URL>/plex
Authorization: ApiKey <DORA_API_KEY>
User-Agent: wsplex-test/1.0
```

Send each request as a single JSON message on the socket. Each request
gets exactly one response (matched by `id`), and subscription requests
produce streaming notifications on the same `path`.

> **Request IDs are single-use per connection.** Generate a fresh UUIDv7
> for every request you actually send. The IDs below are examples —
> reuse them and you'll get a duplicate-request error.

## Placeholders

Replace these with real values before sending. In an HTTP client (e.g.
Bruno, Postman, Insomnia), define them as environment variables.

| Placeholder           | Description                                                                                               |
| --------------------- | --------------------------------------------------------------------------------------------------------- |
| `{{DORA_BASE_URL}}`   | Your environment base URL, e.g. `https://staging.dora.co`                                                 |
| `{{DORA_API_KEY}}`    | Your Dora API key                                                                                         |
| `{{ASSET_ID}}`        | A bond or currency asset UUID (see [Fetching entity IDs](#fetching-entity-ids))                           |
| `{{ASSET_ID_2}}`      | A second asset UUID                                                                                       |
| `{{ORDER_BOOK_ID}}`   | An order book UUID (see [Fetching entity IDs](#fetching-entity-ids))                                      |
| `{{ORDER_BOOK_ID_2}}` | A second order book UUID                                                                                  |
| `{{USER_ID}}`         | Your own Dora user UUID (all roles can access their own data)                                             |
| `{{USER_ID_2}}`       | Another user UUID (requires integrator or admin role — see [User ID access notes](#user-id-access-notes)) |

1 minute resolution = `"1m"`

## Fetching entity IDs

Before sending wsplex requests that take asset or order book IDs, fetch
real IDs from the Dora REST API. All requests require
`Authorization: ApiKey {{DORA_API_KEY}}`.

### List assets (bonds and currencies)

```http
GET {{DORA_BASE_URL}}/v1/assets?limit=100
Authorization: ApiKey {{DORA_API_KEY}}
```

Filter by kind to get only bonds or currencies:

```http
GET {{DORA_BASE_URL}}/v1/assets?asset_kind=BOND&limit=100
Authorization: ApiKey {{DORA_API_KEY}}
```

```http
GET {{DORA_BASE_URL}}/v1/assets?asset_kind=CURRENCY&limit=100
Authorization: ApiKey {{DORA_API_KEY}}
```

The response `data` array contains objects with `id` (the asset UUID),
`name`, `symbol`, and `kind`. Use any `id` where `kind` is `BOND` or
`CURRENCY` as `{{ASSET_ID}}` in `/prices` requests.

### List order books

```http
GET {{DORA_BASE_URL}}/v1/orderbooks?status=OPEN&limit=100
Authorization: ApiKey {{DORA_API_KEY}}
```

The response `data` array contains objects with `order_book_id` (the UUID),
`display_name`, `base_asset_id`, `quote_asset_id`, and `status`. Use any
`order_book_id` where `status` is `OPEN` as `{{ORDER_BOOK_ID}}` in
`/trades`, `/orderbook/stats`, `/charts/candles`, and `/pools/balance`
requests.

### Get your user ID

```http
GET {{DORA_BASE_URL}}/v1/user/self
Authorization: ApiKey {{DORA_API_KEY}}
```

The response `data.id` is your user UUID — use it as `{{USER_ID}}`.

> You can also find your user UUID in the Dora UI under **Settings → Profile**,
> or contact your Dora integration lead.

> **Integrator** roled users can list all their users (users who share the same tenant id)
> with `GET {{DORA_BASE_URL}}/v1/user?limit=50`

> **Admin users** can list all users with `GET {{DORA_BASE_URL}}/v1/user?limit=50`
> to find `{{USER_ID_2}}` values within their tenant.

## User ID access notes

User-scoped channels (`/accounts/balance`, `/orders/byuser`, and `/trades` with
`user_ids`) enforce role-based access:

- **Trader** — can only request their own user ID. Use `{{USER_ID}}` with your
  own UUID; `{{USER_ID_2}}` will return an error.
- **Integrator** — can request any user within the same tenant ID.
- **Dora admin** — can request any user ID.

## Path: `/`

### 1. List all available routes

```json
{
  "id": "019f3c41-b16e-7609-22fa-35a6acf15df8",
  "path": "/",
  "data": {}
}
```

## Path: `/prices`

### 2. /prices: subscribe to two assets

```json
{
  "id": "019f3c41-b16e-7da6-0518-1542f1244ca7",
  "path": "/prices",
  "data": {
    "subscribe": ["{{ASSET_ID}}", "{{ASSET_ID_2}}"]
  }
}
```

### 3. /prices: subscribe_all (all assets)

```json
{
  "id": "019f3c41-b16e-7c09-1924-ecf24cac7d18",
  "path": "/prices",
  "data": {
    "subscribe_all": true
  }
}
```

### 4. /prices: unsubscribe one asset

```json
{
  "id": "019f3c41-b16e-78bb-0d54-828345ff7a12",
  "path": "/prices",
  "data": {
    "unsubscribe": ["{{ASSET_ID}}"]
  }
}
```

### 5. /prices: unsubscribe_all

```json
{
  "id": "019f3c41-b16e-73b7-3a21-9fa570858599",
  "path": "/prices",
  "data": {
    "unsubscribe_all": true
  }
}
```

## Path: `/trades`

### 6. /trades: subscribe to one orderbook (all users)

```json
{
  "id": "019f3c41-b16e-766f-0c53-a81bdabb3d0b",
  "path": "/trades",
  "data": {
    "subscribe": [
      {
        "order_book_ids": ["{{ORDER_BOOK_ID}}"]
      }
    ]
  }
}
```

### 7. /trades: subscribe to all orderbooks, specific user

```json
{
  "id": "019f3c41-b16e-759f-098f-026c552219ab",
  "path": "/trades",
  "data": {
    "subscribe": [
      {
        "order_books_all": true,
        "user_ids": ["{{USER_ID}}"]
      }
    ]
  }
}
```

### 8. /trades: subscribe to all orderbooks, all users

```json
{
  "id": "019f3c41-b16e-72b2-1ccf-0a0e94043da5",
  "path": "/trades",
  "data": {
    "subscribe": [
      {
        "order_books_all": true,
        "users_all": true
      }
    ]
  }
}
```

### 9. /trades: unsubscribe specific orderbook + user pair

```json
{
  "id": "019f3c41-b16e-7cbc-3bbd-48aedabda3c3",
  "path": "/trades",
  "data": {
    "unsubscribe": [
      {
        "order_book_ids": ["{{ORDER_BOOK_ID}}"],
        "user_ids": ["{{USER_ID}}"]
      }
    ]
  }
}
```

### 10. /trades: unsubscribe all (clear everything)

```json
{
  "id": "019f3c41-b16e-7185-3c3c-77adb83b2abd",
  "path": "/trades",
  "data": {
    "unsubscribe": [
      {
        "order_books_all": true,
        "users_all": true
      }
    ]
  }
}
```

## Path: `/assets`

### 11. /assets: subscribe (start stream)

```json
{
  "id": "019f3c41-b16e-75df-34d3-b1f116c0f72d",
  "path": "/assets",
  "data": {
    "subscribe": true
  }
}
```

### 12. /assets: unsubscribe (stop stream)

```json
{
  "id": "019f3c41-b16e-7cbe-3807-e58de0c717f4",
  "path": "/assets",
  "data": {
    "subscribe": false
  }
}
```

## Path: `/orderbook/stats`

### 13. /orderbook/stats: subscribe to two orderbooks

```json
{
  "id": "019f3c41-b16e-718e-3f0a-56286fc4f2ad",
  "path": "/orderbook/stats",
  "data": {
    "subscribe": ["{{ORDER_BOOK_ID}}", "{{ORDER_BOOK_ID_2}}"]
  }
}
```

### 14. /orderbook/stats: subscribe_all

```json
{
  "id": "019f3c41-b16e-72b4-17f1-39abfb278525",
  "path": "/orderbook/stats",
  "data": {
    "subscribe_all": true
  }
}
```

### 15. /orderbook/stats: unsubscribe one orderbook

```json
{
  "id": "019f3c41-b16e-7679-2531-a0fc9d5328f5",
  "path": "/orderbook/stats",
  "data": {
    "unsubscribe": ["{{ORDER_BOOK_ID}}"]
  }
}
```

### 16. /orderbook/stats: unsubscribe_all

```json
{
  "id": "019f3c41-b16e-7ff6-1647-ee0bc588fa1f",
  "path": "/orderbook/stats",
  "data": {
    "unsubscribe_all": true
  }
}
```

## Path: `/charts/candles`

### 17. /charts/candles: subscribe 1m candles for one orderbook

```json
{
  "id": "019f3c41-b16e-7960-3a03-a90f4cbdd636",
  "path": "/charts/candles",
  "data": {
    "subscribe": {
      "orderbook_ids": ["{{ORDER_BOOK_ID}}"],
      "resolution": "1m"
    }
  }
}
```

### 18. /charts/candles: subscribe 1m candles for two orderbooks

```json
{
  "id": "019f3c41-b16e-7a14-193c-b98cadbd3e52",
  "path": "/charts/candles",
  "data": {
    "subscribe": {
      "orderbook_ids": ["{{ORDER_BOOK_ID}}", "{{ORDER_BOOK_ID_2}}"],
      "resolution": "1m"
    }
  }
}
```

### 19. /charts/candles: unsubscribe one orderbook at 1m

```json
{
  "id": "019f3c41-b16e-7897-03f1-8e596b8db3b1",
  "path": "/charts/candles",
  "data": {
    "unsubscribe": {
      "orderbook_ids": ["{{ORDER_BOOK_ID}}"],
      "resolution": "1m"
    }
  }
}
```

### 20. /charts/candles: unsubscribe all

```json
{
  "id": "019f3c41-b16e-7c8a-096a-ad2487cc8f1a",
  "path": "/charts/candles",
  "data": {
    "unsubscribe": {
      "all": true
    }
  }
}
```

## Path: `/accounts/balance`

### 21. /accounts/balance: subscribe to one user

```json
{
  "id": "019f3c41-b16e-706e-3bf4-149df9b9058b",
  "path": "/accounts/balance",
  "data": {
    "subscribe": ["{{USER_ID}}"]
  }
}
```

### 22. /accounts/balance: subscribe to two users

```json
{
  "id": "019f3c41-b16e-76d5-3a91-31d0ae273549",
  "path": "/accounts/balance",
  "data": {
    "subscribe": ["{{USER_ID}}", "{{USER_ID_2}}"]
  }
}
```

### 23. /accounts/balance: unsubscribe one user

```json
{
  "id": "019f3c41-b16e-7e83-0f6f-a022b587e431",
  "path": "/accounts/balance",
  "data": {
    "unsubscribe": ["{{USER_ID}}"]
  }
}
```

## Path: `/pools/balance`

### 24. /pools/balance: subscribe to two pools

```json
{
  "id": "019f3c41-b16e-79ce-03ee-4cccd0725985",
  "path": "/pools/balance",
  "data": {
    "subscribe": ["{{ORDER_BOOK_ID}}", "{{ORDER_BOOK_ID_2}}"]
  }
}
```

### 25. /pools/balance: subscribe_all

```json
{
  "id": "019f3c41-b16e-733c-32b6-1a3a46454e72",
  "path": "/pools/balance",
  "data": {
    "subscribe_all": true
  }
}
```

### 26. /pools/balance: unsubscribe one pool

```json
{
  "id": "019f3c41-b16e-7830-3a17-a0b5e5f344db",
  "path": "/pools/balance",
  "data": {
    "unsubscribe": ["{{ORDER_BOOK_ID}}"]
  }
}
```

### 27. /pools/balance: unsubscribe_all

```json
{
  "id": "019f3c41-b16e-7e7b-16e8-61e30594835e",
  "path": "/pools/balance",
  "data": {
    "unsubscribe_all": true
  }
}
```

## Path: `/orders/byuser`

### 28. /orders/byuser: subscribe for one user, all orderbooks

```json
{
  "id": "019f3c41-b16e-7cc6-1ec0-8cae0dd0c9ab",
  "path": "/orders/byuser",
  "data": {
    "user_id": "{{USER_ID}}",
    "subscribe_all_orderbooks": true
  }
}
```

### 29. /orders/byuser: subscribe for one user, specific orderbooks

```json
{
  "id": "019f3c41-b16e-78cd-00de-ad6374d09dd6",
  "path": "/orders/byuser",
  "data": {
    "user_id": "{{USER_ID}}",
    "subscribe_orderbooks": ["{{ORDER_BOOK_ID}}", "{{ORDER_BOOK_ID_2}}"]
  }
}
```

### 30. /orders/byuser: unsubscribe one orderbook for one user

```json
{
  "id": "019f3c41-b16e-75c2-1ca7-1beb1935a5cd",
  "path": "/orders/byuser",
  "data": {
    "user_id": "{{USER_ID}}",
    "unsubscribe_orderbooks": ["{{ORDER_BOOK_ID}}"]
  }
}
```

### 31. /orders/byuser: unsubscribe all orderbooks for one user

```json
{
  "id": "019f3c41-b16e-7c76-0f5f-909dadae9845",
  "path": "/orders/byuser",
  "data": {
    "user_id": "{{USER_ID}}",
    "unsubscribe_all_orderbooks": true
  }
}
```

## Path: `/debug/notify`

### 32. /debug/notify: echo after 100ms

```json
{
  "id": "019f3c41-b16e-7445-3dd2-9b82a70182ea",
  "path": "/debug/notify",
  "data": {
    "delay": 100000000,
    "data": {
      "ping": "pong"
    }
  }
}
```

### 33. /debug/notify: echo after 2s

```json
{
  "id": "019f3c41-b16e-731b-36f4-5c0449ac7e89",
  "path": "/debug/notify",
  "data": {
    "delay": 2000000000,
    "data": {
      "test": true,
      "order_book_id": "{{ORDER_BOOK_ID}}"
    }
  }
}
```
