# DefaultApi

All URIs are relative to *https://localhost:8084*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**cancelAllOpenOrders**](DefaultApi.md#cancelAllOpenOrders) | **DELETE** /v1/orders | Cancel all open orders, if user passes orderbook on query param it will cancel all orders on specific orderbook, admin can cancel user&#39;s orders on specific orderbook |
| [**cancelOrderById**](DefaultApi.md#cancelOrderById) | **DELETE** /v1/orders/{order_id} | Cancel an order by ID |
| [**checkUserEmailExists**](DefaultApi.md#checkUserEmailExists) | **GET** /v1/user/exists | Check whether a user email exists |
| [**claimLeverageGetAccruedInterest**](DefaultApi.md#claimLeverageGetAccruedInterest) | **POST** /v1/leverage/accrued_interest/claim | Claim current accrued leverage interest for a specific user |
| [**closeIsolatedPosition**](DefaultApi.md#closeIsolatedPosition) | **POST** /v1/positions/close | Close isolated positions, repaying the borrowed |
| [**createAPIKeyForUser**](DefaultApi.md#createAPIKeyForUser) | **POST** /v1/user/apikey | Create apikey for a user |
| [**createAPIKeyForUserID**](DefaultApi.md#createAPIKeyForUserID) | **POST** /v1/user/{user_id}/apikey | Create apikey for a user |
| [**createOrder**](DefaultApi.md#createOrder) | **POST** /v1/orders | Create a new order |
| [**createUser**](DefaultApi.md#createUser) | **POST** /v1/integrators/user | Create a new user |
| [**deleteUser**](DefaultApi.md#deleteUser) | **DELETE** /v1/user/{user_id} | Delete user by ID |
| [**getAPIKeysForUserID**](DefaultApi.md#getAPIKeysForUserID) | **GET** /v1/user/{user_id}/apikey | Get user&#39;s api keys: admin or integrator only |
| [**getAllAssetPrices**](DefaultApi.md#getAllAssetPrices) | **GET** /v1/price | Get the current price of all assets |
| [**getAssetById**](DefaultApi.md#getAssetById) | **GET** /v1/assets/{asset_id} | Get asset by ID |
| [**getAssetPrice**](DefaultApi.md#getAssetPrice) | **GET** /v1/price/asset/{asset_id} | Get the current price of an asset |
| [**getAssetYTMById**](DefaultApi.md#getAssetYTMById) | **GET** /v1/assets/{asset_id}/ytm | Get annualized yield to maturity for a bond asset |
| [**getAssetsStream**](DefaultApi.md#getAssetsStream) | **GET** /v1/assets/stream | Get all inserts or updates for assets |
| [**getCandleData**](DefaultApi.md#getCandleData) | **GET** /v1/charts/{order_book_id}/candle | Get candlestick data for an orderbook |
| [**getCouponPaymentsByAssetId**](DefaultApi.md#getCouponPaymentsByAssetId) | **GET** /v1/assets/{asset_id}/coupon_payments | Get coupon payments for a bond asset |
| [**getL1Depth**](DefaultApi.md#getL1Depth) | **GET** /v1/orderbooks/{order_book_id}/L1 | Get the top price levels for a specific orderbook (L1 market depth) |
| [**getL2Depth**](DefaultApi.md#getL2Depth) | **GET** /v1/orderbooks/{order_book_id}/L2 | Get the aggregated price levels for a specific orderbook (L2 market depth) |
| [**getL3Depth**](DefaultApi.md#getL3Depth) | **GET** /v1/orderbooks/{order_book_id}/L3 | Get all open orders for a specific orderbook (L3 market depth) |
| [**getLedgerBalancesSelf**](DefaultApi.md#getLedgerBalancesSelf) | **GET** /v1/ledger/balances/self | Get your own available, locked, and borrowed assets |
| [**getLedgerInterestSelf**](DefaultApi.md#getLedgerInterestSelf) | **GET** /v1/ledger/interest/self | Get your own interest |
| [**getLedgerModule**](DefaultApi.md#getLedgerModule) | **GET** /v1/ledger/module | Get the entire module object, including unborrowed leverage assets and total leverage trackers |
| [**getLedgerModuleByAsset**](DefaultApi.md#getLedgerModuleByAsset) | **GET** /v1/ledger/module/{asset_id} | Get the module object for a single asset ID |
| [**getLedgerPositionsSelf**](DefaultApi.md#getLedgerPositionsSelf) | **GET** /v1/ledger/positions/self | Get your own positions |
| [**getLedgerValueSelf**](DefaultApi.md#getLedgerValueSelf) | **GET** /v1/ledger/value/self | Get your own available, locked, and borrowed USD value; and realized and unrealized PnL |
| [**getOrderById**](DefaultApi.md#getOrderById) | **GET** /v1/orders/{order_id} | Get order by ID |
| [**getOrderbookById**](DefaultApi.md#getOrderbookById) | **GET** /v1/orderbooks/{order_book_id} | Get orderbook by ID |
| [**getOrderbookDepth**](DefaultApi.md#getOrderbookDepth) | **GET** /v1/orderbooks/{order_book_id}/depth | Get the aggregated price levels for a specific orderbook (L2 market depth) |
| [**getOrderbookOrders**](DefaultApi.md#getOrderbookOrders) | **GET** /v1/orderbooks/{order_book_id}/orders | Get all open orders for a specific orderbook (L3 market depth) |
| [**getOrderbookStats**](DefaultApi.md#getOrderbookStats) | **GET** /v1/orderbooks/{order_book_id}/stats | Get orderbook stats |
| [**getOrderbookStatsStream**](DefaultApi.md#getOrderbookStatsStream) | **GET** /v1/orderbooks/{order_book_id}/stats/stream | Orderbook stats stream |
| [**getOrderbookSummary**](DefaultApi.md#getOrderbookSummary) | **GET** /v1/orderbooks/{order_book_id}/summary | Get summary of an orderbook |
| [**getOrderbookTop**](DefaultApi.md#getOrderbookTop) | **GET** /v1/orderbooks/{order_book_id}/top | Get the top price levels for a specific orderbook (L1 market depth) |
| [**getPLForSelfByAccount**](DefaultApi.md#getPLForSelfByAccount) | **GET** /v1/pl/self | Get account-by-account PL breakdown for the logged in user |
| [**getPoolPrice**](DefaultApi.md#getPoolPrice) | **GET** /v1/price/pool/{pool_id} | Get the current price of a pool |
| [**getTradeById**](DefaultApi.md#getTradeById) | **GET** /v1/trades/{trade_id} | Get a trade by ID |
| [**getTrades**](DefaultApi.md#getTrades) | **GET** /v1/trades | Get a filtered, paginated list of trades |
| [**getTransactionById**](DefaultApi.md#getTransactionById) | **GET** /v1/transactions/{transaction_id} | Get a transaction by ID |
| [**getTransactions**](DefaultApi.md#getTransactions) | **GET** /v1/transactions | Get a filtered, paginated list of transactions |
| [**getUserById**](DefaultApi.md#getUserById) | **GET** /v1/user/{user_id} | Get user by ID (admin only) |
| [**getUserCouponPaymentsStream**](DefaultApi.md#getUserCouponPaymentsStream) | **GET** /v1/user/{user_id}/coupon_payments/stream | Stream user&#39;s coupon payment accruals in real time |
| [**getUserLedgerStream**](DefaultApi.md#getUserLedgerStream) | **GET** /v1/user/{user_id}/ledger/stream | Get a snapshot of user&#39;s ledger updates since a specific time, and opens a stream for further updates |
| [**getUserOrderUpdatesStream**](DefaultApi.md#getUserOrderUpdatesStream) | **GET** /v1/user/{user_id}/orders/{order_book_id}/updates/stream | Get a snapshot of user&#39;s order updates for the given order book since a specific time, and opens a stream for further updates |
| [**getUserOrdersUpdatesStreamAll**](DefaultApi.md#getUserOrdersUpdatesStreamAll) | **GET** /v1/user/{user_id}/orders/all/updates/stream | Get a snapshot of user&#39;s order updates across all order books since a specific time, and opens a stream for further updates |
| [**getUserSelf**](DefaultApi.md#getUserSelf) | **GET** /v1/user/self | Get user details for the authenticated user |
| [**getUserTransactionsStream**](DefaultApi.md#getUserTransactionsStream) | **GET** /v1/user/{user_id}/transactions/stream | Get a snapshot of user&#39;s executed transactions since a specific time, and opens a stream for further updates |
| [**getUsersAPIKeys**](DefaultApi.md#getUsersAPIKeys) | **GET** /v1/user/apikey | Get user&#39;s api keys |
| [**ledgerDeposit**](DefaultApi.md#ledgerDeposit) | **POST** /v1/ledger/deposit/{user_id} | Deposit assets into this user&#39;s account from the outside world |
| [**ledgerWithdraw**](DefaultApi.md#ledgerWithdraw) | **POST** /v1/ledger/withdraw/{user_id} | Withdraw assets from this user to the outside world |
| [**leverageGetAccruedInterestByUser**](DefaultApi.md#leverageGetAccruedInterestByUser) | **GET** /v1/leverage/accrued_interest/self | Get current accrued leverage interest for the user |
| [**leverageIsolateCollateral**](DefaultApi.md#leverageIsolateCollateral) | **POST** /v1/leverage/isolate_collateral | Create an isolated position by transferring collateral to the position from the user&#39;s global collateral |
| [**leverageSupply**](DefaultApi.md#leverageSupply) | **POST** /v1/leverage/supply | Supply leverage for a specific asset |
| [**leverageUnite**](DefaultApi.md#leverageUnite) | **POST** /v1/leverage/unite | Combines all isolated positions into a single global position |
| [**leverageWithdraw**](DefaultApi.md#leverageWithdraw) | **POST** /v1/leverage/withdraw | Withdraw leverage for a specific asset |
| [**liquidityAdd**](DefaultApi.md#liquidityAdd) | **POST** /v1/liquidity/pool/{pool_id}/add | Add liquidity to a pool |
| [**liquiditySubtract**](DefaultApi.md#liquiditySubtract) | **POST** /v1/liquidity/pool/{pool_id}/remove | Subtract liquidity from a pool |
| [**listAssets**](DefaultApi.md#listAssets) | **GET** /v1/assets | List assets |
| [**listOrderBooks**](DefaultApi.md#listOrderBooks) | **GET** /v1/orderbooks | List order books |
| [**listOrders**](DefaultApi.md#listOrders) | **GET** /v1/orders | List all orders |
| [**listPositionAccountsSelf**](DefaultApi.md#listPositionAccountsSelf) | **GET** /v1/user/self/position_accounts | List all position accounts for the authenticated user |
| [**payLeverageGetAccruedInterest**](DefaultApi.md#payLeverageGetAccruedInterest) | **POST** /v1/leverage/accrued_interest/pay | Pay current accrued leverage interest for a specific user |
| [**revokeAPIKeyForUser**](DefaultApi.md#revokeAPIKeyForUser) | **PUT** /v1/user/apikey/{key_id}/revoke | Revoke apikey for a user |
| [**revokeAPIKeyForUserID**](DefaultApi.md#revokeAPIKeyForUserID) | **PUT** /v1/user/{user_id}/apikey/{key_id}/revoke | Revoke apikey for a user: admin or integrator only |
| [**settleLeverageAccruedInterest**](DefaultApi.md#settleLeverageAccruedInterest) | **POST** /v1/leverage/accrued_interest/settle | Settle current accrued leverage interest for a specific user |
| [**streamAssetPrices**](DefaultApi.md#streamAssetPrices) | **GET** /v1/prices/stream | Stream real-time asset prices as map objects |
| [**streamCandleData**](DefaultApi.md#streamCandleData) | **GET** /v1/charts/{order_book_id}/candle/stream | Get a snapshot of candlestick data from date provided, and open a stream for real-time updates |
| [**streamOrderBookBalances**](DefaultApi.md#streamOrderBookBalances) | **GET** /v1/orderbooks/{order_book_id}/balances/stream | Get a snapshot of base and quote balances for an order book and open a stream for real-time updates |
| [**streamOrderbookOpenOrders**](DefaultApi.md#streamOrderbookOpenOrders) | **GET** /v1/orderbooks/{order_book_id}/open/stream | Get a snapshot of open orders in an order book and open a stream for real-time updates |
| [**streamTrades**](DefaultApi.md#streamTrades) | **GET** /v1/trades/{order_book_id}/stream | Get a snapshot of trades executed on the given order book from a specific date and open a stream for real-time updates |
| [**transferAvailableBalances**](DefaultApi.md#transferAvailableBalances) | **POST** /v1/positions/transfer_balances | Transfer available balance between a user&#39;s accounts (e.g. global to isolated position) |
| [**updateUserConfig**](DefaultApi.md#updateUserConfig) | **PUT** /v1/user/{user_id}/config | Update user configuration by ID |
| [**updateUserConfigSelf**](DefaultApi.md#updateUserConfigSelf) | **PUT** /v1/user/config/self | Update user configuration for the authenticated user |
| [**validateSubmitOrder**](DefaultApi.md#validateSubmitOrder) | **POST** /v1/orders/validate | Validate submit order request data |
| [**verifyUser**](DefaultApi.md#verifyUser) | **PUT** /v1/user/{user_id}/verify | Verify a user by ID |


<a name="cancelAllOpenOrders"></a>
# **cancelAllOpenOrders**
> ListOrdersResponseEnvelope cancelAllOpenOrders(order\_book\_id, user\_id, order\_kind)

Cancel all open orders, if user passes orderbook on query param it will cancel all orders on specific orderbook, admin can cancel user&#39;s orders on specific orderbook

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **String**|  | [optional] [default to null] |
| **user\_id** | **UUID**|  | [optional] [default to null] |
| **order\_kind** | [**OrderKind**](../Models/.md)|  | [optional] [default to null] [enum: LIMIT, MARKET] |

### Return type

[**ListOrdersResponseEnvelope**](../Models/ListOrdersResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="cancelOrderById"></a>
# **cancelOrderById**
> CancelOrderResponseEnvelope cancelOrderById(order\_id)

Cancel an order by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_id** | **UUID**|  | [default to null] |

### Return type

[**CancelOrderResponseEnvelope**](../Models/CancelOrderResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="checkUserEmailExists"></a>
# **checkUserEmailExists**
> EmailExistsResponseEnvelope checkUserEmailExists(email)

Check whether a user email exists

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **email** | **String**|  | [default to null] |

### Return type

[**EmailExistsResponseEnvelope**](../Models/EmailExistsResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="claimLeverageGetAccruedInterest"></a>
# **claimLeverageGetAccruedInterest**
> ClaimLeverageAccruedInterestResponseEnvelope claimLeverageGetAccruedInterest(ClaimLeverageAccruedInterestRequest)

Claim current accrued leverage interest for a specific user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ClaimLeverageAccruedInterestRequest** | [**ClaimLeverageAccruedInterestRequest**](../Models/ClaimLeverageAccruedInterestRequest.md)|  | |

### Return type

[**ClaimLeverageAccruedInterestResponseEnvelope**](../Models/ClaimLeverageAccruedInterestResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="closeIsolatedPosition"></a>
# **closeIsolatedPosition**
> ClosePositionResponseEnvelope closeIsolatedPosition(ClosePositionRequest)

Close isolated positions, repaying the borrowed

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ClosePositionRequest** | [**ClosePositionRequest**](../Models/ClosePositionRequest.md)|  | |

### Return type

[**ClosePositionResponseEnvelope**](../Models/ClosePositionResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createAPIKeyForUser"></a>
# **createAPIKeyForUser**
> CreateAPIKeyResponseEnvelope createAPIKeyForUser(CreateAPIKeyRequest)

Create apikey for a user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateAPIKeyRequest** | [**CreateAPIKeyRequest**](../Models/CreateAPIKeyRequest.md)|  | |

### Return type

[**CreateAPIKeyResponseEnvelope**](../Models/CreateAPIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createAPIKeyForUserID"></a>
# **createAPIKeyForUserID**
> CreateAPIKeyResponseEnvelope createAPIKeyForUserID(user\_id, CreateAPIKeyRequest)

Create apikey for a user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **String**|  | [default to null] |
| **CreateAPIKeyRequest** | [**CreateAPIKeyRequest**](../Models/CreateAPIKeyRequest.md)|  | |

### Return type

[**CreateAPIKeyResponseEnvelope**](../Models/CreateAPIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createOrder"></a>
# **createOrder**
> CreateOrderResponseEnvelope createOrder(CreateOrderRequest)

Create a new order

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateOrderRequest** | [**CreateOrderRequest**](../Models/CreateOrderRequest.md)|  | |

### Return type

[**CreateOrderResponseEnvelope**](../Models/CreateOrderResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createUser"></a>
# **createUser**
> UserCreatedResponseEnvelope createUser(CreateIntegratorUserRequest)

Create a new user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **CreateIntegratorUserRequest** | [**CreateIntegratorUserRequest**](../Models/CreateIntegratorUserRequest.md)|  | |

### Return type

[**UserCreatedResponseEnvelope**](../Models/UserCreatedResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteUser"></a>
# **deleteUser**
> UserDeletedResponseEnvelope deleteUser(user\_id)

Delete user by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |

### Return type

[**UserDeletedResponseEnvelope**](../Models/UserDeletedResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAPIKeysForUserID"></a>
# **getAPIKeysForUserID**
> APIKeyResponseEnvelope getAPIKeysForUserID(user\_id)

Get user&#39;s api keys: admin or integrator only

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **String**|  | [default to null] |

### Return type

[**APIKeyResponseEnvelope**](../Models/APIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllAssetPrices"></a>
# **getAllAssetPrices**
> ListAssetPriceResponseEnvelope getAllAssetPrices()

Get the current price of all assets

### Parameters
This endpoint does not need any parameter.

### Return type

[**ListAssetPriceResponseEnvelope**](../Models/ListAssetPriceResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAssetById"></a>
# **getAssetById**
> GetAssetByIDResponseEnvelope getAssetById(asset\_id)

Get asset by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset\_id** | **UUID**|  | [default to null] |

### Return type

[**GetAssetByIDResponseEnvelope**](../Models/GetAssetByIDResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAssetPrice"></a>
# **getAssetPrice**
> AssetPriceResponseEnvelope getAssetPrice(asset\_id)

Get the current price of an asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset\_id** | **UUID**|  | [default to null] |

### Return type

[**AssetPriceResponseEnvelope**](../Models/AssetPriceResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAssetYTMById"></a>
# **getAssetYTMById**
> GetAssetYTMByIDResponseEnvelope getAssetYTMById(asset\_id)

Get annualized yield to maturity for a bond asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset\_id** | **UUID**|  | [default to null] |

### Return type

[**GetAssetYTMByIDResponseEnvelope**](../Models/GetAssetYTMByIDResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAssetsStream"></a>
# **getAssetsStream**
> List getAssetsStream(since, until)

Get all inserts or updates for assets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **since** | **Date**|  | [optional] [default to null] |
| **until** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamAssetsEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCandleData"></a>
# **getCandleData**
> ListCandlesResponseEnvelope getCandleData(order\_book\_id, start, end, resolution)

Get candlestick data for an orderbook

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **String**|  | [default to null] |
| **start** | **Date**|  | [optional] [default to null] |
| **end** | **Date**|  | [optional] [default to null] |
| **resolution** | [**CandleResolution**](../Models/.md)|  | [optional] [default to null] [enum: 1m, 5m, 15m, 1h, 4h, 1d] |

### Return type

[**ListCandlesResponseEnvelope**](../Models/ListCandlesResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCouponPaymentsByAssetId"></a>
# **getCouponPaymentsByAssetId**
> ListCouponPaymentsResponseEnvelope getCouponPaymentsByAssetId(asset\_id)

Get coupon payments for a bond asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset\_id** | **UUID**|  | [default to null] |

### Return type

[**ListCouponPaymentsResponseEnvelope**](../Models/ListCouponPaymentsResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getL1Depth"></a>
# **getL1Depth**
> GetTopOfBookResponseEnvelope getL1Depth(order\_book\_id)

Get the top price levels for a specific orderbook (L1 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**GetTopOfBookResponseEnvelope**](../Models/GetTopOfBookResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getL2Depth"></a>
# **getL2Depth**
> ListOrderBookDepthResponseEnvelope getL2Depth(order\_book\_id)

Get the aggregated price levels for a specific orderbook (L2 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**ListOrderBookDepthResponseEnvelope**](../Models/ListOrderBookDepthResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getL3Depth"></a>
# **getL3Depth**
> ListOrdersResponseEnvelope getL3Depth(order\_book\_id)

Get all open orders for a specific orderbook (L3 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**ListOrdersResponseEnvelope**](../Models/ListOrdersResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerBalancesSelf"></a>
# **getLedgerBalancesSelf**
> UserBalanceResponseEnvelope getLedgerBalancesSelf()

Get your own available, locked, and borrowed assets

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserBalanceResponseEnvelope**](../Models/UserBalanceResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerInterestSelf"></a>
# **getLedgerInterestSelf**
> UserInterestResponseEnvelope getLedgerInterestSelf()

Get your own interest

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserInterestResponseEnvelope**](../Models/UserInterestResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerModule"></a>
# **getLedgerModule**
> LedgerModuleResponseEnvelope getLedgerModule()

Get the entire module object, including unborrowed leverage assets and total leverage trackers

### Parameters
This endpoint does not need any parameter.

### Return type

[**LedgerModuleResponseEnvelope**](../Models/LedgerModuleResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerModuleByAsset"></a>
# **getLedgerModuleByAsset**
> LedgerModuleByAssetResponseEnvelope getLedgerModuleByAsset(asset\_id)

Get the module object for a single asset ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **asset\_id** | **UUID**|  | [default to null] |

### Return type

[**LedgerModuleByAssetResponseEnvelope**](../Models/LedgerModuleByAssetResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerPositionsSelf"></a>
# **getLedgerPositionsSelf**
> UserPositionResponseEnvelope getLedgerPositionsSelf()

Get your own positions

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserPositionResponseEnvelope**](../Models/UserPositionResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getLedgerValueSelf"></a>
# **getLedgerValueSelf**
> UserValueResponseEnvelope getLedgerValueSelf()

Get your own available, locked, and borrowed USD value; and realized and unrealized PnL

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserValueResponseEnvelope**](../Models/UserValueResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderById"></a>
# **getOrderById**
> OrderResponseEnvelope getOrderById(order\_id)

Get order by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_id** | **UUID**|  | [default to null] |

### Return type

[**OrderResponseEnvelope**](../Models/OrderResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookById"></a>
# **getOrderbookById**
> OrderBookResponseEnvelope getOrderbookById(order\_book\_id)

Get orderbook by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**OrderBookResponseEnvelope**](../Models/OrderBookResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookDepth"></a>
# **getOrderbookDepth**
> ListOrderBookDepthResponseEnvelope getOrderbookDepth(order\_book\_id)

Get the aggregated price levels for a specific orderbook (L2 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**ListOrderBookDepthResponseEnvelope**](../Models/ListOrderBookDepthResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookOrders"></a>
# **getOrderbookOrders**
> ListOrdersResponseEnvelope getOrderbookOrders(order\_book\_id)

Get all open orders for a specific orderbook (L3 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**ListOrdersResponseEnvelope**](../Models/ListOrdersResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookStats"></a>
# **getOrderbookStats**
> OrderbookStatsResponseEnvelope getOrderbookStats(order\_book\_id)

Get orderbook stats

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**OrderbookStatsResponseEnvelope**](../Models/OrderbookStatsResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookStatsStream"></a>
# **getOrderbookStatsStream**
> OrderbookStats getOrderbookStatsStream(order\_book\_id)

Orderbook stats stream

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**OrderbookStats**](../Models/OrderbookStats.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookSummary"></a>
# **getOrderbookSummary**
> OrderBookSummaryResponseEnvelope getOrderbookSummary(order\_book\_id)

Get summary of an orderbook

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**OrderBookSummaryResponseEnvelope**](../Models/OrderBookSummaryResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getOrderbookTop"></a>
# **getOrderbookTop**
> GetTopOfBookResponseEnvelope getOrderbookTop(order\_book\_id)

Get the top price levels for a specific orderbook (L1 market depth)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |

### Return type

[**GetTopOfBookResponseEnvelope**](../Models/GetTopOfBookResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPLForSelfByAccount"></a>
# **getPLForSelfByAccount**
> PLResponseEnvelope getPLForSelfByAccount()

Get account-by-account PL breakdown for the logged in user

### Parameters
This endpoint does not need any parameter.

### Return type

[**PLResponseEnvelope**](../Models/PLResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPoolPrice"></a>
# **getPoolPrice**
> PoolPriceResponseEnvelope getPoolPrice(pool\_id)

Get the current price of a pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pool\_id** | **UUID**|  | [default to null] |

### Return type

[**PoolPriceResponseEnvelope**](../Models/PoolPriceResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTradeById"></a>
# **getTradeById**
> TradeResponseEnvelope getTradeById(trade\_id)

Get a trade by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **trade\_id** | **UUID**|  | [default to null] |

### Return type

[**TradeResponseEnvelope**](../Models/TradeResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTrades"></a>
# **getTrades**
> ListTradeResponseEnvelope getTrades(order\_book\_ids, user\_ids, start, end, page, limit)

Get a filtered, paginated list of trades

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_ids** | [**List**](../Models/String.md)|  | [optional] [default to null] |
| **user\_ids** | [**List**](../Models/UUID.md)|  | [optional] [default to null] |
| **start** | **Date**|  | [optional] [default to null] |
| **end** | **Date**|  | [optional] [default to null] |
| **page** | **Integer**|  | [optional] [default to 1] |
| **limit** | **Integer**|  | [optional] [default to 100] |

### Return type

[**ListTradeResponseEnvelope**](../Models/ListTradeResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactionById"></a>
# **getTransactionById**
> TransactionResponseEnvelope getTransactionById(transaction\_id)

Get a transaction by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **transaction\_id** | **UUID**|  | [default to null] |

### Return type

[**TransactionResponseEnvelope**](../Models/TransactionResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTransactions"></a>
# **getTransactions**
> ListTransactionsResponseEnvelope getTransactions(pools, user\_ids, tx\_kinds, start, end, page, limit)

Get a filtered, paginated list of transactions

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pools** | [**List**](../Models/String.md)|  | [optional] [default to null] |
| **user\_ids** | [**List**](../Models/UUID.md)|  | [optional] [default to null] |
| **tx\_kinds** | [**List**](../Models/TransactionKind.md)|  | [optional] [default to null] |
| **start** | **Date**|  | [optional] [default to null] |
| **end** | **Date**|  | [optional] [default to null] |
| **page** | **Integer**|  | [optional] [default to 1] |
| **limit** | **Integer**|  | [optional] [default to 100] |

### Return type

[**ListTransactionsResponseEnvelope**](../Models/ListTransactionsResponseEnvelope.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserById"></a>
# **getUserById**
> UserEnvelope getUserById(user\_id)

Get user by ID (admin only)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |

### Return type

[**UserEnvelope**](../Models/UserEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserCouponPaymentsStream"></a>
# **getUserCouponPaymentsStream**
> StreamUserCouponPaymentsResponse getUserCouponPaymentsStream(user\_id)

Stream user&#39;s coupon payment accruals in real time

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |

### Return type

[**StreamUserCouponPaymentsResponse**](../Models/StreamUserCouponPaymentsResponse.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserLedgerStream"></a>
# **getUserLedgerStream**
> List getUserLedgerStream(user\_id)

Get a snapshot of user&#39;s ledger updates since a specific time, and opens a stream for further updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |

### Return type

[**List**](../Models/StreamPositionsEntry.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserOrderUpdatesStream"></a>
# **getUserOrderUpdatesStream**
> List getUserOrderUpdatesStream(user\_id, order\_book\_id, since)

Get a snapshot of user&#39;s order updates for the given order book since a specific time, and opens a stream for further updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **order\_book\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamOrderUpdatesEntry.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserOrdersUpdatesStreamAll"></a>
# **getUserOrdersUpdatesStreamAll**
> List getUserOrdersUpdatesStreamAll(user\_id, since)

Get a snapshot of user&#39;s order updates across all order books since a specific time, and opens a stream for further updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamOrderUpdatesEntry.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserSelf"></a>
# **getUserSelf**
> UserEnvelope getUserSelf()

Get user details for the authenticated user

### Parameters
This endpoint does not need any parameter.

### Return type

[**UserEnvelope**](../Models/UserEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUserTransactionsStream"></a>
# **getUserTransactionsStream**
> List getUserTransactionsStream(user\_id, since)

Get a snapshot of user&#39;s executed transactions since a specific time, and opens a stream for further updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamTransactionsEntry.md)

### Authorization

[apiKeyAuthQuery](../README.md#apiKeyAuthQuery)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUsersAPIKeys"></a>
# **getUsersAPIKeys**
> APIKeyResponseEnvelope getUsersAPIKeys()

Get user&#39;s api keys

### Parameters
This endpoint does not need any parameter.

### Return type

[**APIKeyResponseEnvelope**](../Models/APIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="ledgerDeposit"></a>
# **ledgerDeposit**
> FundUserResponseEnvelope ledgerDeposit(user\_id, FundUserRequest)

Deposit assets into this user&#39;s account from the outside world

    Deposit assets into this user&#39;s account from the outside world. Note that this does not interact with any external systems; it simply adds the amount to the user&#39;s available balance in the ledger. Actual transfer of assets must be handled separately.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **FundUserRequest** | [**FundUserRequest**](../Models/FundUserRequest.md)|  | |

### Return type

[**FundUserResponseEnvelope**](../Models/FundUserResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="ledgerWithdraw"></a>
# **ledgerWithdraw**
> FundUserResponseEnvelope ledgerWithdraw(user\_id, DefundUserRequest)

Withdraw assets from this user to the outside world

    Withdraw assets from this user&#39;s account to the outside world. Note that this does not interact with any external systems; it simply deducts the amount from the user&#39;s available balance in the ledger. Actual transfer of assets must be handled separately.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **DefundUserRequest** | [**DefundUserRequest**](../Models/DefundUserRequest.md)|  | |

### Return type

[**FundUserResponseEnvelope**](../Models/FundUserResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="leverageGetAccruedInterestByUser"></a>
# **leverageGetAccruedInterestByUser**
> CurrentLeverageAccruedInterestResponseEnvelope leverageGetAccruedInterestByUser(position\_id, asset\_id)

Get current accrued leverage interest for the user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **position\_id** | **UUID**|  | [optional] [default to null] |
| **asset\_id** | **UUID**|  | [optional] [default to null] |

### Return type

[**CurrentLeverageAccruedInterestResponseEnvelope**](../Models/CurrentLeverageAccruedInterestResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="leverageIsolateCollateral"></a>
# **leverageIsolateCollateral**
> IsolateCollateralResponse leverageIsolateCollateral(IsolateCollateralRequest)

Create an isolated position by transferring collateral to the position from the user&#39;s global collateral

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **IsolateCollateralRequest** | [**IsolateCollateralRequest**](../Models/IsolateCollateralRequest.md)|  | |

### Return type

[**IsolateCollateralResponse**](../Models/IsolateCollateralResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="leverageSupply"></a>
# **leverageSupply**
> SupplyResponseEnvelope leverageSupply(SupplyRequest)

Supply leverage for a specific asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **SupplyRequest** | [**SupplyRequest**](../Models/SupplyRequest.md)|  | |

### Return type

[**SupplyResponseEnvelope**](../Models/SupplyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="leverageUnite"></a>
# **leverageUnite**
> UnitePositionResponseEnvelope leverageUnite(UnitePositionRequest)

Combines all isolated positions into a single global position

    Combines all isolated positions into a single global position

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UnitePositionRequest** | [**UnitePositionRequest**](../Models/UnitePositionRequest.md)|  | |

### Return type

[**UnitePositionResponseEnvelope**](../Models/UnitePositionResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="leverageWithdraw"></a>
# **leverageWithdraw**
> WithdrawResponseEnvelope leverageWithdraw(WithdrawRequest)

Withdraw leverage for a specific asset

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **WithdrawRequest** | [**WithdrawRequest**](../Models/WithdrawRequest.md)|  | |

### Return type

[**WithdrawResponseEnvelope**](../Models/WithdrawResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="liquidityAdd"></a>
# **liquidityAdd**
> LiquidityResponseEnvelope liquidityAdd(pool\_id, LiquidityRequest)

Add liquidity to a pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pool\_id** | **UUID**|  | [default to null] |
| **LiquidityRequest** | [**LiquidityRequest**](../Models/LiquidityRequest.md)|  | |

### Return type

[**LiquidityResponseEnvelope**](../Models/LiquidityResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="liquiditySubtract"></a>
# **liquiditySubtract**
> LiquidityResponseEnvelope liquiditySubtract(pool\_id, LiquidityRequest)

Subtract liquidity from a pool

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **pool\_id** | **UUID**|  | [default to null] |
| **LiquidityRequest** | [**LiquidityRequest**](../Models/LiquidityRequest.md)|  | |

### Return type

[**LiquidityResponseEnvelope**](../Models/LiquidityResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="listAssets"></a>
# **listAssets**
> ResponseEnvelopeOfListAssets listAssets(created\_after, created\_before, asset\_kind, can\_add\_liquidity, can\_direct\_borrow, can\_onboard, can\_trade, can\_virtual\_borrow, page, limit)

List assets

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **created\_after** | **Date**|  | [optional] [default to null] |
| **created\_before** | **Date**|  | [optional] [default to null] |
| **asset\_kind** | [**AssetKind**](../Models/.md)| Asset kind (BOND, CURRENCY, INTEREST, POOL_SHARE) | [optional] [default to null] [enum: BOND, CURRENCY, INTEREST, POOL_SHARE] |
| **can\_add\_liquidity** | **Boolean**|  | [optional] [default to null] |
| **can\_direct\_borrow** | **Boolean**|  | [optional] [default to null] |
| **can\_onboard** | **Boolean**|  | [optional] [default to null] |
| **can\_trade** | **Boolean**|  | [optional] [default to null] |
| **can\_virtual\_borrow** | **Boolean**|  | [optional] [default to null] |
| **page** | **Integer**|  | [optional] [default to 1] |
| **limit** | **Integer**|  | [optional] [default to 100] |

### Return type

[**ResponseEnvelopeOfListAssets**](../Models/ResponseEnvelopeOfListAssets.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listOrderBooks"></a>
# **listOrderBooks**
> ListOrderbookResponseEnvelope listOrderBooks(status, base\_asset\_id, quote\_asset\_id, page, limit)

List order books

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **status** | [**OrderBookStatus**](../Models/.md)|  | [optional] [default to null] [enum: CLOSED, OPEN, SUSPENDED] |
| **base\_asset\_id** | **UUID**|  | [optional] [default to null] |
| **quote\_asset\_id** | **UUID**|  | [optional] [default to null] |
| **page** | **Integer**|  | [optional] [default to 1] |
| **limit** | **Integer**|  | [optional] [default to 100] |

### Return type

[**ListOrderbookResponseEnvelope**](../Models/ListOrderbookResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listOrders"></a>
# **listOrders**
> ListOrdersResponseEnvelope listOrders(order\_book\_id, kind, status, side, from, to, page, limit)

List all orders

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | [**List**](../Models/UUID.md)|  | [optional] [default to null] |
| **kind** | [**List**](../Models/OrderKind.md)|  | [optional] [default to null] |
| **status** | [**List**](../Models/OrderStatus.md)|  | [optional] [default to null] |
| **side** | [**Side**](../Models/.md)|  | [optional] [default to null] [enum: BUY, SELL] |
| **from** | **Date**|  | [optional] [default to null] |
| **to** | **Date**|  | [optional] [default to null] |
| **page** | **Integer**|  | [optional] [default to 1] |
| **limit** | **Integer**|  | [optional] [default to 100] |

### Return type

[**ListOrdersResponseEnvelope**](../Models/ListOrdersResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listPositionAccountsSelf"></a>
# **listPositionAccountsSelf**
> ListPositionAccountsResponseEnvelope listPositionAccountsSelf()

List all position accounts for the authenticated user

### Parameters
This endpoint does not need any parameter.

### Return type

[**ListPositionAccountsResponseEnvelope**](../Models/ListPositionAccountsResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="payLeverageGetAccruedInterest"></a>
# **payLeverageGetAccruedInterest**
> PayLeverageAccruedInterestResponseEnvelope payLeverageGetAccruedInterest(PayLeverageAccruedInterestRequest)

Pay current accrued leverage interest for a specific user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **PayLeverageAccruedInterestRequest** | [**PayLeverageAccruedInterestRequest**](../Models/PayLeverageAccruedInterestRequest.md)|  | |

### Return type

[**PayLeverageAccruedInterestResponseEnvelope**](../Models/PayLeverageAccruedInterestResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="revokeAPIKeyForUser"></a>
# **revokeAPIKeyForUser**
> RevokeAPIKeyResponseEnvelope revokeAPIKeyForUser(key\_id)

Revoke apikey for a user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **key\_id** | **String**|  | [default to null] |

### Return type

[**RevokeAPIKeyResponseEnvelope**](../Models/RevokeAPIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="revokeAPIKeyForUserID"></a>
# **revokeAPIKeyForUserID**
> RevokeAPIKeyResponseEnvelope revokeAPIKeyForUserID(user\_id, key\_id)

Revoke apikey for a user: admin or integrator only

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **String**|  | [default to null] |
| **key\_id** | **String**|  | [default to null] |

### Return type

[**RevokeAPIKeyResponseEnvelope**](../Models/RevokeAPIKeyResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="settleLeverageAccruedInterest"></a>
# **settleLeverageAccruedInterest**
> SettleLeverageAccruedInterestResponseEnvelope settleLeverageAccruedInterest(SettleLeverageAccruedInterestRequest)

Settle current accrued leverage interest for a specific user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **SettleLeverageAccruedInterestRequest** | [**SettleLeverageAccruedInterestRequest**](../Models/SettleLeverageAccruedInterestRequest.md)|  | |

### Return type

[**SettleLeverageAccruedInterestResponseEnvelope**](../Models/SettleLeverageAccruedInterestResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="streamAssetPrices"></a>
# **streamAssetPrices**
> Map streamAssetPrices(since, asset\_id)

Stream real-time asset prices as map objects

    Opens a WebSocket stream for real-time asset price updates. First message contains all current prices, subsequent messages contain only changed prices. Data is sent as JSON objects keyed by asset ID.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **since** | **Date**|  | [optional] [default to null] |
| **asset\_id** | **UUID**|  | [optional] [default to null] |

### Return type

[**Map**](../Models/AssetPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="streamCandleData"></a>
# **streamCandleData**
> List streamCandleData(order\_book\_id, since, resolution)

Get a snapshot of candlestick data from date provided, and open a stream for real-time updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **String**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |
| **resolution** | [**CandleResolution**](../Models/.md)|  | [optional] [default to null] [enum: 1m, 5m, 15m, 1h, 4h, 1d] |

### Return type

[**List**](../Models/StreamCandlesEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="streamOrderBookBalances"></a>
# **streamOrderBookBalances**
> List streamOrderBookBalances(order\_book\_id, since)

Get a snapshot of base and quote balances for an order book and open a stream for real-time updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamOrderBookBalanceEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="streamOrderbookOpenOrders"></a>
# **streamOrderbookOpenOrders**
> LiveOrderbook streamOrderbookOpenOrders(order\_book\_id, since)

Get a snapshot of open orders in an order book and open a stream for real-time updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**LiveOrderbook**](../Models/LiveOrderbook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="streamTrades"></a>
# **streamTrades**
> List streamTrades(order\_book\_id, since)

Get a snapshot of trades executed on the given order book from a specific date and open a stream for real-time updates

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **order\_book\_id** | **UUID**|  | [default to null] |
| **since** | **Date**|  | [optional] [default to null] |

### Return type

[**List**](../Models/StreamTradesEntry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="transferAvailableBalances"></a>
# **transferAvailableBalances**
> TransferBalancesResponseEnvelope transferAvailableBalances(TransferBalancesRequest)

Transfer available balance between a user&#39;s accounts (e.g. global to isolated position)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **TransferBalancesRequest** | [**TransferBalancesRequest**](../Models/TransferBalancesRequest.md)|  | |

### Return type

[**TransferBalancesResponseEnvelope**](../Models/TransferBalancesResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserConfig"></a>
# **updateUserConfig**
> UserUpdatedResponseEnvelope updateUserConfig(user\_id, UpdateUserConfigRequest)

Update user configuration by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |
| **UpdateUserConfigRequest** | [**UpdateUserConfigRequest**](../Models/UpdateUserConfigRequest.md)|  | |

### Return type

[**UserUpdatedResponseEnvelope**](../Models/UserUpdatedResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUserConfigSelf"></a>
# **updateUserConfigSelf**
> UserUpdatedResponseEnvelope updateUserConfigSelf(UpdateUserConfigRequest)

Update user configuration for the authenticated user

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UpdateUserConfigRequest** | [**UpdateUserConfigRequest**](../Models/UpdateUserConfigRequest.md)|  | |

### Return type

[**UserUpdatedResponseEnvelope**](../Models/UserUpdatedResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="validateSubmitOrder"></a>
# **validateSubmitOrder**
> ValidateSubmitOrderResponse validateSubmitOrder(ValidateSubmitOrderRequest)

Validate submit order request data

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **ValidateSubmitOrderRequest** | [**ValidateSubmitOrderRequest**](../Models/ValidateSubmitOrderRequest.md)|  | |

### Return type

[**ValidateSubmitOrderResponse**](../Models/ValidateSubmitOrderResponse.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="verifyUser"></a>
# **verifyUser**
> UserUpdatedResponseEnvelope verifyUser(user\_id)

Verify a user by ID

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **user\_id** | **UUID**|  | [default to null] |

### Return type

[**UserUpdatedResponseEnvelope**](../Models/UserUpdatedResponseEnvelope.md)

### Authorization

[apiKeyAuthHeader](../README.md#apiKeyAuthHeader), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

