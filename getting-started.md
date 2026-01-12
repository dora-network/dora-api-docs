# Getting started with the DORA API

## Introduction

The DORA platform provides a digital marketplace for trading fractionalized bonds
opening a new era of financial inclusion.

The DORA API allows developers to interact with the platform programmatically,
enabling them to build applications that can trade, make markets, manage positions
and analyse market data on the DORA platform.

## Authentication

At present, DORA's API requires authentication via a Bearer token or an API key.

### Bearer Token

To obtain a Bearer token, you need to log in to the DORA UI using your credentials.

Once logged in, you can find your authentication token in the user menu by clicking
on your user avatar in the top right corner.

Click on the `COPY TOKEN` button to copy your authentication token to the clipboard.

You can then use this token to authenticate your API requests by including it in the
`Authorization` header as follows:

```
Authorization: Bearer <your-authentication-token>
```

Authentication tokens are valid for a limited time and will need to be refreshed
periodically using the method above.

### API Key

#### API Key Management via DORA UI

Log into DORA as usual, using the user menu in the top right corner,

![User Menu](./images/dora-ui-top-right.png)

choose the account settings option from the menu.

![Full User Menu](./images/dora-ui-user-menu.png)

This will load the account settings page where you will be able to generate
your API keys.

![API Key Generator](./images/dora-ui-api-key-generation.png)

In the `Label` box, give a name for your api, this is to help you remember what the key is for.

You can optionally set an expiration date for the key. If an expiration date is set
the key will automatically become invalid at the expiration time.

Click the `Generate Key` button to generate the key, you will be presented with a
dialog presenting you with the key id, the label you selected, and the generated key.

You will be warned to save the generated key as it cannot be displayed again after
you close the dialog.

The keys you have generated will be listed below the API Key Generation section. To revoke
any key, choose the key you want to revoke and press the `Revoke` button.

#### API Key Management via API

To create an API key for the first time, you will need to use your Bearer token
to authenticate the request. Follow the steps above to obtain your authentication token.

> Note: This authentication token is refreshed periodically for security reasons.
> if you encounter authentication issues while using the API, please repeat the
> steps above to obtain a new token, then try again.

Once you have your authentication token, you can use the following `curl` command
to request a new API key:

```bash
curl -L -X POST -H "Authorizatio: Bearer <YOUR TOKEN>" -H "Content-Type: application/json" --data '{"label": "your-label-for-identifying-this-key"}' https://dora-staging.fly.dev/v1/user/apikey
```

If you receive an error response like this:

```json
{
  "error":"unauthorized: token has invalid claims: token is expired",
  "metadata":{
    "status_code":401,
    "trace_id":"staging-019b9916-a6cb-7f12-a897-f1a7ca750ee7",
    "request_id":"staging-019b9916-a6cb-7f15-89d5-0a7023e32dac"
  }
}
```

The authorization key you provided is invalid or expired. See note above for details.
Obtain a new token from the DORA UI and try again.

If the request is successful, you will receive a response like this:

```json
{
  "data": {
    "key_id": "unique identifier for your key",
    "api_key": "your-generated-api-key",
    "label": "your-label-for-identifying-this-key"
  },
  "metadata": {
    ... metadata fields ...
  }
}
```

The `key_id` is the unique identifier for your api key and will be required when you
want to delete the key. When requesting your list of keys, this identifier is provided
in the response.

The `label` is the identifying label you provided when creating the key.

The `api_key` is the key you will use to authenticate your API requests to DORA. Please
store this value securely, as it will not be shown again. If you lose this value, you
will need to delete the key and create a new one.

### Using the API Key for Authentication

Once you have generated your API key, you can use it to authenticate your requests to
DORA's API. To do this, replace the `Bearer <token>` part of your Authorization header
with `ApiKey <your-generated-api-key>`

**Example:**

Replace this:

```bash
curl -L -X GET -H "Authorization: Bearer <token> ...

```

With this:

```bash
curl -L -X GET -H "Authorization: ApiKey <your-generated-api-key> ...
```

For streaming API requests, you need to provide the api key as a query parameter instead of in the header.
The query parameter to pass is `x-api-key`.

**Example:**

```bash
wscat -c "wss://dora-staging.fly.dev:8085/v1/orderbooks/<order_book_id>/open/stream?since=2026-01-01T09:28:23.687804Z&x-api-key=<your-generated-api-key>
```

### Listing your API Keys

Once you have your generated API keys, you can start using them to authenticate your requests.
You can test that your key is working by listing your existing API keys.

```bash
curl -L -H "Authorization: ApiKey <your-generated-api-key>" -H "Content-Type: application/json" https://dora-staging.fly.dev/v1/user/apikey
```

You should receive a response like this:

```json
{
  "data": {
    "api_keys": [
      {
        "user_id": "your-user-id",
        "key_id": "unique identifier for your key",
        "label": "your-label-for-identifying-this-key",
        "expires_at": null,
        "is_active": true
      },
    ]
  },
  "metadata": {
    ... metadata fields ...
  }
}
```

### Revoking an API Key

To revoke an API key, you will need the `key_id` of the key you want to delete.
It is recommended to create a new API key before revoking an existing one to avoid any disruption in service.
Once you have created a new key, you can use the key in the request to revoke the old key.

If you revoked an API key that is currently in use, and you don't have another valid API key,
you will need to login to the DORA UI to generate a new authentication token to create a new API key.

To revoke an API key, you can make a `PUT` request to the following endpoint: `https://dora-staging.fly.dev/v1/user/apikey/{key-id}/revoke`

**Example:**

```bash
curl -L -X PUT -H "Authorization: ApiKey <your-generated-api-key>" -H "Content-Type: application/json" https://dora-staging.fly.dev/v1/user/apikey/{key_id}/revoke
```

### Streaming APIs

DORA provides several websocket endpoints for real-time data streaming. These endpoints allow you to receive updates on
market data, user orders, transactions, and ledgers.

For streaming endpoints that require authentication, you need to provide your API key or bearer token as a query parameter.

For Bearer token authentication, use the `token` query parameter:

`GET /v1/user/{user_id}/orders/all/stream?token=<your-authentication-token>`

For API key authentication, use the `x-api-key` query parameter:

`GET /v1/user/{user_id}/orders/all/stream?x-api-key=<your-generated-api-key>`

All streaming endpoints allow users to fetch historical data before receiving the real-time updates. This is useful for
cases where you want to start processing data from a specific point in time. For example, if you disconnected from DORA
for a while, you can fetch the data you missed when not connected, before subscribing to the real-time updates.

### Market Data

The following endpoints are available for market data streaming:

`GET /v1/orderbook/{orderbook_id}/open/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides L3 order book updates for the given order book.

`GET /v1/charts/{orderbook_id}/candle/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]&[resolution=XX]` provides candlestick bars
for the chosen resolution (e.g., 1 minute, 5 minutes, etc.) for the given order book. If no resolution is specified,
the default is 1 minute.

`GET v1/trades/{orderbook_id}/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides trades that have been executed on the given order book,
i.e. tick data for the order book. This includes information about the price, volume, and time of each trade.
This endpoint is useful for tracking the history of trades and understanding market activity.

`GET /v1/prices/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides the latest price updates for all assets on the DORA platform.
This endpoint is useful for tracking the current price of assets and understanding market trends.

`GET /v1/orderbook/{orderbook_id}/balances/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides updates on the balances of the base and
quote asset of the order book. The balances provide a view of the current liquidity available in the AMM pool associated with
the order book.

### User Specific Data

`GET /v1/user/{user_id}/orders/all/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides updates on all orders placed by a specific user.
This endpoint allows you to track the status of orders, including when they are placed, filled, or cancelled.

`GET /v1/user/{user_id}/ledger/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides updates on the user's ledger, which includes positions
and balances available for trading etc.

`GET /v1/user/{user_id}/transactions/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]` provides updates on the user's transactions, generated
on the DORA platform. This includes deposits, withdrawals, coupon payments, and other financial activities.

## Using the API

### User configuration

DORA provides limited support for updating user account data. The endpoint:

`PUT /v1/user/config/self` is a shortcut for updating your user avatar url, or timezone. Any other user data changes must be done
via the DORA web application or through the user interface provided by the platform.

To retrieve your user configuration data, you can use the endpoint:

`GET /v1/user/self`

### Assets

An asset is a tradable item on the DORA platform. It can be a currency, a bond, or other financial instrument. Each order book
is associated with two assets: a base asset and a quote asset. The base asset is the asset being traded, while the quote asset
is the currency or equivalent used to price the base asset.

A list of all available assets can be retrieved using the endpoint:

`GET /v1/assets`

If the asset is a bond, it will have additional bond specific information, such as the coupon rate, maturity date,
and other relevant details.

This endpoint also provides a filter allowing you to query for specific asset types, such as currencies or bonds, and
specific properties, such as assets created before or after a certain date.

To retrieve detailed information about a specific asset, you can use the endpoint:

`GET /v1/assets/{asset_id}`

The asset ID is a UUID v7 unique identifier that is automatically generated by the system when the asset is created.

#### Coupons

For assets that are bonds, you can retrieve information about the coupons associated with the bond using the endpoint:

`GET /v1/assets/{asset_id}/coupons`

### Order books

On DORA, each order book is a market for trading a specific pair of assets, such as a bond and a currency.
The order book is matched with an Automated Market Maker (AMM) pool, which provides liquidity for the order book.

Market orders are automatically executed against the AMM pool, while limit orders are added to the order book and
matched with other orders when the price conditions are met.

To retrieve a list of all order books available on the DORA platform, you can use the endpoint:

`GET /v1/orderbooks`

To filter the order books, you can use query parameters such as `base_asset_id` and `quote_asset_id` to filter by
specific asset pairs, or `status` to filter by the status of the order book (e.g., open, closed, suspended).

To retrieve detailed information about a specific order book, you can use the endpoint:

`GET /v1/orderbooks/{orderbook_id}`

The order book ID is a UUID v7 unique identifier that is automatically generated by the system when the order book
is created.

The API also provides 3 levels of market depth for order books:

- **Level 1 (L1)**: Best bid and ask prices, along with their respective volumes.
- **Level 2 (L2)**: Full order book with all bids and asks, including their prices and volumes.
- **Level 3 (L3)**: Full order book with individual orders, including their prices, volumes, and order IDs.

These can be accessed via the 

- `GET /v1/orderbook/{orderbook_id}/L1` 
- `GET /v1/orderbook/{orderbook_id}/L2` 
- `GET /v1/orderbook/{orderbook_id}/L3` 

endpoints. Additionally, you can subscribe to real-time updates for these levels via the streaming endpoints
as described in the "Streaming APIs" section.

A summary of the order book information can be retrieved using the endpoint:

`GET /v1/orderbooks/{orderbook_id}/summary` which provides a high-level overview of the order book,
including the best bid and ask prices, mid price and percentage spread.

### Orders

#### Submitting orders

To place an order on the DORA platform, you can use the endpoint:

`POST /v1/orders`

This endpoint allows you to submit a new order to the specified order book. The create order request looks like this:

```json
{
  "quantity": "100",  // number of units to buy/sell expressed as a string
  "inverse_leverage": 1.0, // the inverse leverage for the order, this is a decimal value between 0 and 1.0
  "price": "100.0", // required for limit orders, decimal value expressed as a string
  "kind": "limit", // the kind of order, can be "market" or "limit"
  "side": "buy", // the side of the order, can be "buy" or "sell"
  "from_global_position": true, // whether to use the global position for the order or an isolated position
  "order_book_id": "some-orderbook-id", // this must be a uuid v7 id of the order book you want to place the order on
  "order_modifiers": [
    "MAX_BUY"
  ] // optional order modifiers, can be used to modify the behavior of the order. Currently MAX_BUY is the only supported modifier.
  good_till_date: "2024-12-31T23:59:59Z", // optional good till date for the order, expressed in ISO 8601 format
  "trigger_price": "95.0", // optional trigger price for conditional orders, decimal value expressed as a string
  "trigger_type": "STOP_LOSS" // optional trigger type for conditional orders, can be "STOP_LOSS" or "TAKE_PROFIT"
}
```

For market orders, the `price` field is not required and should be omitted. Setting the `MAX_BUY` modifier allows the order to be
executed at the best available price, up to the specified quantity, without exceeding the user's available balance.

#### Cancelling orders

You can cancel an open limit order using the endpoint:

`DELETE /v1/orders/{order_id}`

The order ID is a UUID v7 unique identifier that is automatically generated by the system when the order is created.
Market orders will be automatically executed against the AMM pool and cannot be cancelled once submitted.

You can also cancel all your open orders at once with the endpoint:

`DELETE /v1/orders`

This will use the `user_id` from the authentication token to identify the user and cancel all their open orders.

#### Updating orders

It is not possible to update an existing order on the DORA platform. If you need to change an order, you must cancel it
and then place a new order with the desired changes.

#### Querying orders

To retrieve a list of your orders, you can use the endpoint:

`GET /v1/orders`

This endpoint allows you to filter orders by various parameters, such as `status`, `order_book_id`, `kind`, and `side`.
You can also provide a date range using the `from` and `to` query parameters to filter orders by their creation date.
If only a `from` date is provided, the API will return all orders created after that date. If only a `to` date is provided,
the API will return all orders created before that date. If both dates are provided, the API will return all orders
created within the specified date range.

To retrieve detailed information about a specific order, you can use the endpoint:

`GET /v1/orders/{order_id}`

The order ID is a UUID v7 unique identifier automatically generated by the system when the order is created.

To receive updates on your orders in real-time, you can subscribe to the streaming endpoint:

`GET /v1/user/{user_id}/orders/all/stream?[since=YYYY-MM-DDTHH:mm:ssZZZ]`

This will allow you to receive updates on all your orders, including when they are placed, filled, or cancelled.
It is recommended to use the `since` parameter to specify a starting point for the updates, so you can
catch up on any missed updates since your last connection. Once connected, you will receive real-time updates
on your orders as you place them, and they are processed, filled, or cancelled.

### Ledgers

The DORA platform maintains a ledger for each user, which tracks their positions, balances, and other relevant information.
To retrieve your ledger, you can use the endpoint:

`GET /v1/ledger/positions/self`

Use your DORA user ID to retrieve your ledger. The user ID is a UUID v7 unique identifier that is automatically generated by
the system when your user account is created.

The position response looks like this:

```json
{
  "portfolio": {
    "user_id": "some-user-uuid",
    "positon": {
      "some-asset-uuid": {
        "some-position-uuid": {
          "id": "some-position-uuid",
          "asset_id": "some-asset-uuid",
          "seq": 1,
          "is_global": true,
          "available": "5000.0",
          "locked": "0.0",
          "supplied": "2000.0",
          "collateral": "0.0",
          "supplied_collateral": "0.0",
          "borrowed": "0.0",
          "impending_borrows": "0.0"
        },
        "some-other-position-uuid": {
          "id": "some-other-position-uuid",
          "asset_id": "some-other-asset-uuid",
          "seq": 2,
          "is_global": false,
          "available": "1000.0",
          "locked": "0.0",
          "supplied": "0.0",
          "collateral": "0.0",
          "supplied_collateral": "0.0",
          "borrowed": "0.0",
          "impending_borrows": "0.0"
        }
      },
      "some-other-asset-uuid": {
        ...
        // more positions for other assets
      },
      ... // more asset positions
    },
    "net_stablecoin_equivalence": {
      "gained": "1000.0", // the total amount of stablecoin equivalent gained
      "lost": "0" // the total amount of stablecoin equivalent lost
    }
  }
}
```

The `portfolio` object contains a map of asset IDs to their respective positions.
Each portfolio contains a single `global` position, which is a special position that represents the user's main account balance.
A position that is not flagged as `global` is an `isolated` position, which is similar to a sub-account or a specific trading
balance that has been created for a leveraged position.

To view just a list of all your global positions, you can use the endpoint:

`GET /v1/ledger/balances/self`

instead. You can also filter the positions by asset ID using the `asset_id` query parameter.

### Leverage operations

The DORA platform allows users to supply assets to a leverage pool, borrow assets from the pool, and perform other associated operations.
The leverage pool is a shared pool of assets that provides liquidity for leveraged trading, and is separate from the liquidity pool used for
the AMM order book.

To supply assets to the leverage pool, you can use the endpoint:

`POST /v1/leverage/supply`

This will transfer the specified amount of assets from your chosen position to the leverage pool, the balance will be reflected in your position,
and your available balance will be updated accordingly. The request body should look like this:

```json
{
  "position_id": "some-user-uuid", // the user ID of the user supplying the asset
  "asset_id": "some-asset-uuid", // the asset you want to supply
  "amount": "1000.0" // the amount of the asset to supply, expressed as a string
}
```

To withdraw assets from the leverage pool, you can use the endpoint:

`POST /v1/leverage/withdraw`

> TODO: Explain the borrowing and repaying operations, as well as the collateral and position isolation features.

### Reporting

DORA provides several endpoints for retrieving information for reporting and analysis purposes. In addition to the
orders API which allows you to retrieve historical information about your orders, you can also retrieve
information about the trades you have executed, and transactions that have occurred on the platform.

#### Trades

To get a list of trades that you have executed on the DORA platform, you can use the endpoint:

`GET /v1/trades`

> TODO: Describe the usage of this endpoint once the implementation has been completed.

#### Transactions

The DORA platform records all transactions, including orders, trades, and leverage operations.
To retrieve a list of transactions, you can use the endpoint:

`GET /v1/transactions`

This endpoint allows you to filter transactions by various parameters, such as `pools`, `user_ids`, `tx_kind`, and date ranges.

To retrieve information about a specific transaction, you can use the endpoint:

`GET /v1/transactions/{transaction_id}`

### Historical Data

DORA provides limited support for retrieving historical trade, order book, and candlestick data from the platform. In addition to the streaming endpoints
described earlier that allows you to fetch historical data before receiving real-time updates, you can also use the following endpoints:

- `GET /v1/trade?[pools=poolid1,poolid2,...]&[user_ids=user_id1,user_id2,...]&[tx_kind=kind1,kind2,...]&[from=YYYY-MM-DDTHH:mm:ssZZZ]&[to=YYYY-MM-DDTHH:mm:ssZZZ]`
- `GET /v1/charts/candle/{orderbook_id}/L3?[start=YYYY-MM-DDTHH:mm:ssZZZ]&[end=YYYY-MM-DDTHH:mm:ssZZZ]&[resolution=XX]`
