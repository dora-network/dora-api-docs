# Documentation for DORA

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *https://localhost:8084*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *DefaultApi* | [**cancelAllOpenOrders**](Apis/DefaultApi.md#cancelAllOpenOrders) | **DELETE** /v1/orders | Cancel all open orders, if user passes orderbook on query param it will cancel all orders on specific orderbook, admin can cancel user's orders on specific orderbook |
*DefaultApi* | [**cancelOrderById**](Apis/DefaultApi.md#cancelOrderById) | **DELETE** /v1/orders/{order_id} | Cancel an order by ID |
*DefaultApi* | [**checkUserEmailExists**](Apis/DefaultApi.md#checkUserEmailExists) | **GET** /v1/user/exists | Check whether a user email exists |
*DefaultApi* | [**claimLeverageGetAccruedInterest**](Apis/DefaultApi.md#claimLeverageGetAccruedInterest) | **POST** /v1/leverage/accrued_interest/claim | Claim current accrued leverage interest for a specific user |
*DefaultApi* | [**closeIsolatedPosition**](Apis/DefaultApi.md#closeIsolatedPosition) | **POST** /v1/positions/close | Close isolated positions, repaying the borrowed |
*DefaultApi* | [**createAPIKeyForUser**](Apis/DefaultApi.md#createAPIKeyForUser) | **POST** /v1/user/apikey | Create apikey for a user |
*DefaultApi* | [**createNewIsolatedPosition**](Apis/DefaultApi.md#createNewIsolatedPosition) | **POST** /v1/positions/new_isolated | Create a new isolated position for a user transferring available assets into the position |
*DefaultApi* | [**createOrder**](Apis/DefaultApi.md#createOrder) | **POST** /v1/orders | Create a new order |
*DefaultApi* | [**deleteUser**](Apis/DefaultApi.md#deleteUser) | **DELETE** /v1/user/{user_id} | Delete user by ID |
*DefaultApi* | [**getAllAssetPrices**](Apis/DefaultApi.md#getAllAssetPrices) | **GET** /v1/price | Get the current price of all assets |
*DefaultApi* | [**getAssetById**](Apis/DefaultApi.md#getAssetById) | **GET** /v1/assets/{asset_id} | Get asset by ID |
*DefaultApi* | [**getAssetPrice**](Apis/DefaultApi.md#getAssetPrice) | **GET** /v1/price/asset/{asset_id} | Get the current price of an asset |
*DefaultApi* | [**getAssetsStream**](Apis/DefaultApi.md#getAssetsStream) | **GET** /v1/assets/stream | Get all inserts or updates for assets |
*DefaultApi* | [**getCandleData**](Apis/DefaultApi.md#getCandleData) | **GET** /v1/charts/{order_book_id}/candle | Get candlestick data for an orderbook |
*DefaultApi* | [**getCouponPaymentsByAssetId**](Apis/DefaultApi.md#getCouponPaymentsByAssetId) | **GET** /v1/assets/{asset_id}/coupon_payments | Get coupon payments for a bond asset |
*DefaultApi* | [**getL1Depth**](Apis/DefaultApi.md#getL1Depth) | **GET** /v1/orderbooks/{order_book_id}/L1 | Get the top price levels for a specific orderbook (L1 market depth) |
*DefaultApi* | [**getL2Depth**](Apis/DefaultApi.md#getL2Depth) | **GET** /v1/orderbooks/{order_book_id}/L2 | Get the aggregated price levels for a specific orderbook (L2 market depth) |
*DefaultApi* | [**getL3Depth**](Apis/DefaultApi.md#getL3Depth) | **GET** /v1/orderbooks/{order_book_id}/L3 | Get all open orders for a specific orderbook (L3 market depth) |
*DefaultApi* | [**getLedgerBalancesSelf**](Apis/DefaultApi.md#getLedgerBalancesSelf) | **GET** /v1/ledger/balances/self | Get your own available, locked, and borrowed assets |
*DefaultApi* | [**getLedgerInterestSelf**](Apis/DefaultApi.md#getLedgerInterestSelf) | **GET** /v1/ledger/interest/self | Get your own interest |
*DefaultApi* | [**getLedgerModule**](Apis/DefaultApi.md#getLedgerModule) | **GET** /v1/ledger/module | Get the entire module object, including unborrowed leverage assets and total leverage trackers |
*DefaultApi* | [**getLedgerModuleByAsset**](Apis/DefaultApi.md#getLedgerModuleByAsset) | **GET** /v1/ledger/module/{asset_id} | Get the module object for a single asset ID |
*DefaultApi* | [**getLedgerPositionsSelf**](Apis/DefaultApi.md#getLedgerPositionsSelf) | **GET** /v1/ledger/positions/self | Get your own positions |
*DefaultApi* | [**getLedgerValueSelf**](Apis/DefaultApi.md#getLedgerValueSelf) | **GET** /v1/ledger/value/self | Get your own available, locked, and borrowed USD value; and realized and unrealized PnL |
*DefaultApi* | [**getOrderById**](Apis/DefaultApi.md#getOrderById) | **GET** /v1/orders/{order_id} | Get order by ID |
*DefaultApi* | [**getOrderbookById**](Apis/DefaultApi.md#getOrderbookById) | **GET** /v1/orderbooks/{order_book_id} | Get orderbook by ID |
*DefaultApi* | [**getOrderbookDepth**](Apis/DefaultApi.md#getOrderbookDepth) | **GET** /v1/orderbooks/{order_book_id}/depth | Get the aggregated price levels for a specific orderbook (L2 market depth) |
*DefaultApi* | [**getOrderbookOrders**](Apis/DefaultApi.md#getOrderbookOrders) | **GET** /v1/orderbooks/{order_book_id}/orders | Get all open orders for a specific orderbook (L3 market depth) |
*DefaultApi* | [**getOrderbookStats**](Apis/DefaultApi.md#getOrderbookStats) | **GET** /v1/orderbooks/{order_book_id}/stats | Get orderbook stats |
*DefaultApi* | [**getOrderbookStatsStream**](Apis/DefaultApi.md#getOrderbookStatsStream) | **GET** /v1/orderbooks/{order_book_id}/stats/stream | Orderbook stats stream |
*DefaultApi* | [**getOrderbookSummary**](Apis/DefaultApi.md#getOrderbookSummary) | **GET** /v1/orderbooks/{order_book_id}/summary | Get summary of an orderbook |
*DefaultApi* | [**getOrderbookTop**](Apis/DefaultApi.md#getOrderbookTop) | **GET** /v1/orderbooks/{order_book_id}/top | Get the top price levels for a specific orderbook (L1 market depth) |
*DefaultApi* | [**getPoolPrice**](Apis/DefaultApi.md#getPoolPrice) | **GET** /v1/price/pool/{pool_id} | Get the current price of a pool |
*DefaultApi* | [**getTradeById**](Apis/DefaultApi.md#getTradeById) | **GET** /v1/trades/{trade_id} | Get a trade by ID |
*DefaultApi* | [**getTrades**](Apis/DefaultApi.md#getTrades) | **GET** /v1/trades | Get a filtered, paginated list of trades |
*DefaultApi* | [**getTransactionById**](Apis/DefaultApi.md#getTransactionById) | **GET** /v1/transactions/{transaction_id} | Get a transaction by ID |
*DefaultApi* | [**getTransactions**](Apis/DefaultApi.md#getTransactions) | **GET** /v1/transactions | Get a filtered, paginated list of transactions |
*DefaultApi* | [**getUserById**](Apis/DefaultApi.md#getUserById) | **GET** /v1/user/{user_id} | Get user by ID (admin only) |
*DefaultApi* | [**getUserLedgerStream**](Apis/DefaultApi.md#getUserLedgerStream) | **GET** /v1/user/{user_id}/ledger/stream | Get a snapshot of user's ledger updates since a specific time, and opens a stream for further updates |
*DefaultApi* | [**getUserOrderUpdatesStream**](Apis/DefaultApi.md#getUserOrderUpdatesStream) | **GET** /v1/user/{user_id}/orders/{order_book_id}/updates/stream | Get a snapshot of user's order updates for the given order book since a specific time, and opens a stream for further updates |
*DefaultApi* | [**getUserOrdersUpdatesStreamAll**](Apis/DefaultApi.md#getUserOrdersUpdatesStreamAll) | **GET** /v1/user/{user_id}/orders/all/updates/stream | Get a snapshot of user's order updates across all order books since a specific time, and opens a stream for further updates |
*DefaultApi* | [**getUserSelf**](Apis/DefaultApi.md#getUserSelf) | **GET** /v1/user/self | Get user details for the authenticated user |
*DefaultApi* | [**getUserTransactionsStream**](Apis/DefaultApi.md#getUserTransactionsStream) | **GET** /v1/user/{user_id}/transactions/stream | Get a snapshot of user's executed transactions since a specific time, and opens a stream for further updates |
*DefaultApi* | [**getUsersAPIKeys**](Apis/DefaultApi.md#getUsersAPIKeys) | **GET** /v1/user/apikey | Get user's api keys |
*DefaultApi* | [**leverageGetAccruedInterestByUser**](Apis/DefaultApi.md#leverageGetAccruedInterestByUser) | **GET** /v1/leverage/accrued_interest/self | Get current accrued leverage interest for the user |
*DefaultApi* | [**leverageIsolateCollateral**](Apis/DefaultApi.md#leverageIsolateCollateral) | **POST** /v1/leverage/isolate_collateral | Create an isolated position by transferring collateral to the position from the user's global collateral |
*DefaultApi* | [**leverageSupply**](Apis/DefaultApi.md#leverageSupply) | **POST** /v1/leverage/supply | Supply leverage for a specific asset |
*DefaultApi* | [**leverageUnite**](Apis/DefaultApi.md#leverageUnite) | **POST** /v1/leverage/unite | Combines all isolated positions into a single global position |
*DefaultApi* | [**leverageWithdraw**](Apis/DefaultApi.md#leverageWithdraw) | **POST** /v1/leverage/withdraw | Withdraw leverage for a specific asset |
*DefaultApi* | [**liquidityAdd**](Apis/DefaultApi.md#liquidityAdd) | **POST** /v1/liquidity/pool/{pool_id}/add | Add liquidity to a pool |
*DefaultApi* | [**liquiditySubtract**](Apis/DefaultApi.md#liquiditySubtract) | **POST** /v1/liquidity/pool/{pool_id}/remove | Subtract liquidity from a pool |
*DefaultApi* | [**listAssets**](Apis/DefaultApi.md#listAssets) | **GET** /v1/assets | List assets |
*DefaultApi* | [**listOrderBooks**](Apis/DefaultApi.md#listOrderBooks) | **GET** /v1/orderbooks | List order books |
*DefaultApi* | [**listOrders**](Apis/DefaultApi.md#listOrders) | **GET** /v1/orders | List all orders |
*DefaultApi* | [**listPositionAccountsSelf**](Apis/DefaultApi.md#listPositionAccountsSelf) | **GET** /v1/user/self/position_accounts | List all position accounts for the authenticated user |
*DefaultApi* | [**payLeverageGetAccruedInterest**](Apis/DefaultApi.md#payLeverageGetAccruedInterest) | **POST** /v1/leverage/accrued_interest/pay | Pay current accrued leverage interest for a specific user |
*DefaultApi* | [**revokeAPIKeyForUser**](Apis/DefaultApi.md#revokeAPIKeyForUser) | **PUT** /v1/user/apikey/{key_id}/revoke | Revoke apikey for a user |
*DefaultApi* | [**streamAssetPrices**](Apis/DefaultApi.md#streamAssetPrices) | **GET** /v1/prices/stream | Stream real-time asset prices as map objects |
*DefaultApi* | [**streamCandleData**](Apis/DefaultApi.md#streamCandleData) | **GET** /v1/charts/{order_book_id}/candle/stream | Get a snapshot of candlestick data from date provided, and open a stream for real-time updates |
*DefaultApi* | [**streamOrderBookBalances**](Apis/DefaultApi.md#streamOrderBookBalances) | **GET** /v1/orderbooks/{order_book_id}/balances/stream | Get a snapshot of base and quote balances for an order book and open a stream for real-time updates |
*DefaultApi* | [**streamOrderbookOpenOrders**](Apis/DefaultApi.md#streamOrderbookOpenOrders) | **GET** /v1/orderbooks/{order_book_id}/open/stream | Get a snapshot of open orders in an order book and open a stream for real-time updates |
*DefaultApi* | [**streamTrades**](Apis/DefaultApi.md#streamTrades) | **GET** /v1/trades/{order_book_id}/stream | Get a snapshot of trades executed on the given order book from a specific date and open a stream for real-time updates |
*DefaultApi* | [**transferAvailableBalances**](Apis/DefaultApi.md#transferAvailableBalances) | **POST** /v1/positions/transfer_balances | Transfer available balance between a user's accounts (e.g. global to isolated position) |
*DefaultApi* | [**updateUserConfig**](Apis/DefaultApi.md#updateUserConfig) | **PUT** /v1/user/{user_id}/config | Update user configuration by ID |
*DefaultApi* | [**updateUserConfigSelf**](Apis/DefaultApi.md#updateUserConfigSelf) | **PUT** /v1/user/config/self | Update user configuration for the authenticated user |
*DefaultApi* | [**validateSubmitOrder**](Apis/DefaultApi.md#validateSubmitOrder) | **POST** /v1/orders/validate | Validate submit order request data |
*DefaultApi* | [**verifyUser**](Apis/DefaultApi.md#verifyUser) | **PUT** /v1/user/{user_id}/verify | Verify a user by ID |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [APIKeyResponse](./Models/APIKeyResponse.md)
 - [APIKeyResponseEnvelope](./Models/APIKeyResponseEnvelope.md)
 - [APIKeys](./Models/APIKeys.md)
 - [Asset](./Models/Asset.md)
 - [AssetConfig](./Models/AssetConfig.md)
 - [AssetKind](./Models/AssetKind.md)
 - [AssetPrice](./Models/AssetPrice.md)
 - [AssetPriceResponseEnvelope](./Models/AssetPriceResponseEnvelope.md)
 - [AssetRequestError](./Models/AssetRequestError.md)
 - [BalanceTransfer](./Models/BalanceTransfer.md)
 - [BalancesResponse](./Models/BalancesResponse.md)
 - [Bond](./Models/Bond.md)
 - [BondKind](./Models/BondKind.md)
 - [CancelOrderResponseEnvelope](./Models/CancelOrderResponseEnvelope.md)
 - [Candle](./Models/Candle.md)
 - [CandleResolution](./Models/CandleResolution.md)
 - [ClaimLeverageAccruedInterest](./Models/ClaimLeverageAccruedInterest.md)
 - [ClaimLeverageAccruedInterestReq](./Models/ClaimLeverageAccruedInterestReq.md)
 - [ClaimLeverageAccruedInterestRequest](./Models/ClaimLeverageAccruedInterestRequest.md)
 - [ClaimLeverageAccruedInterestResponseEnvelope](./Models/ClaimLeverageAccruedInterestResponseEnvelope.md)
 - [ClosePositionRequest](./Models/ClosePositionRequest.md)
 - [ClosePositionResp](./Models/ClosePositionResp.md)
 - [ClosePositionResponseEnvelope](./Models/ClosePositionResponseEnvelope.md)
 - [Collateral](./Models/Collateral.md)
 - [CouponPayment](./Models/CouponPayment.md)
 - [CreateAPIKeyData](./Models/CreateAPIKeyData.md)
 - [CreateAPIKeyRequest](./Models/CreateAPIKeyRequest.md)
 - [CreateAPIKeyResponseEnvelope](./Models/CreateAPIKeyResponseEnvelope.md)
 - [CreateOrUpdateUserResponse](./Models/CreateOrUpdateUserResponse.md)
 - [CreateOrderRequest](./Models/CreateOrderRequest.md)
 - [CreateOrderResponseEnvelope](./Models/CreateOrderResponseEnvelope.md)
 - [CurrentLeverageAccruedInterest](./Models/CurrentLeverageAccruedInterest.md)
 - [CurrentLeverageAccruedInterestResponseEnvelope](./Models/CurrentLeverageAccruedInterestResponseEnvelope.md)
 - [EmailExistsResponseEnvelope](./Models/EmailExistsResponseEnvelope.md)
 - [GetAssetByIDResponseEnvelope](./Models/GetAssetByIDResponseEnvelope.md)
 - [GetTopOfBookResponseEnvelope](./Models/GetTopOfBookResponseEnvelope.md)
 - [IsolateCollateralRequest](./Models/IsolateCollateralRequest.md)
 - [IsolateCollateralResponse](./Models/IsolateCollateralResponse.md)
 - [IsolatedCollateral](./Models/IsolatedCollateral.md)
 - [IsolatedPosition](./Models/IsolatedPosition.md)
 - [LedgerModuleByAssetResponseEnvelope](./Models/LedgerModuleByAssetResponseEnvelope.md)
 - [LedgerModuleResponseEnvelope](./Models/LedgerModuleResponseEnvelope.md)
 - [LeverageBalanceResponse](./Models/LeverageBalanceResponse.md)
 - [LeverageModuleResponse](./Models/LeverageModuleResponse.md)
 - [LeverageRequestError](./Models/LeverageRequestError.md)
 - [LeverageType](./Models/LeverageType.md)
 - [LiquidationTargetsResponseEnvelope](./Models/LiquidationTargetsResponseEnvelope.md)
 - [Liquidity](./Models/Liquidity.md)
 - [LiquidityRequest](./Models/LiquidityRequest.md)
 - [LiquidityResponseEnvelope](./Models/LiquidityResponseEnvelope.md)
 - [ListAssetPriceResponseEnvelope](./Models/ListAssetPriceResponseEnvelope.md)
 - [ListCandlesResponseEnvelope](./Models/ListCandlesResponseEnvelope.md)
 - [ListCouponPaymentsResponseEnvelope](./Models/ListCouponPaymentsResponseEnvelope.md)
 - [ListOrderBookDepthResponseEnvelope](./Models/ListOrderBookDepthResponseEnvelope.md)
 - [ListOrderbookResponseEnvelope](./Models/ListOrderbookResponseEnvelope.md)
 - [ListOrdersResponseEnvelope](./Models/ListOrdersResponseEnvelope.md)
 - [ListPositionAccountsResponseEnvelope](./Models/ListPositionAccountsResponseEnvelope.md)
 - [ListPositionsResponse](./Models/ListPositionsResponse.md)
 - [ListTradeResponseEnvelope](./Models/ListTradeResponseEnvelope.md)
 - [ListTransactionsResponseEnvelope](./Models/ListTransactionsResponseEnvelope.md)
 - [LiveOrderbook](./Models/LiveOrderbook.md)
 - [Metadata](./Models/Metadata.md)
 - [ModuleBalance](./Models/ModuleBalance.md)
 - [NewIsolatedPositionRequest](./Models/NewIsolatedPositionRequest.md)
 - [NewIsolatedPositionResponseEnvelope](./Models/NewIsolatedPositionResponseEnvelope.md)
 - [Order](./Models/Order.md)
 - [OrderBook](./Models/OrderBook.md)
 - [OrderBookBalance](./Models/OrderBookBalance.md)
 - [OrderBookDepth](./Models/OrderBookDepth.md)
 - [OrderBookHaltResponseEnvelope](./Models/OrderBookHaltResponseEnvelope.md)
 - [OrderBookResponseEnvelope](./Models/OrderBookResponseEnvelope.md)
 - [OrderBookResumeResponseEnvelope](./Models/OrderBookResumeResponseEnvelope.md)
 - [OrderBookStatus](./Models/OrderBookStatus.md)
 - [OrderBookSummary](./Models/OrderBookSummary.md)
 - [OrderBookSummaryResponseEnvelope](./Models/OrderBookSummaryResponseEnvelope.md)
 - [OrderBookTerminateResponseEnvelope](./Models/OrderBookTerminateResponseEnvelope.md)
 - [OrderBookTop](./Models/OrderBookTop.md)
 - [OrderId](./Models/OrderId.md)
 - [OrderKind](./Models/OrderKind.md)
 - [OrderModifierKind](./Models/OrderModifierKind.md)
 - [OrderResponseEnvelope](./Models/OrderResponseEnvelope.md)
 - [OrderStatus](./Models/OrderStatus.md)
 - [OrderbookStats](./Models/OrderbookStats.md)
 - [OrderbookStatsResponseEnvelope](./Models/OrderbookStatsResponseEnvelope.md)
 - [PayLeverageAccruedInterest](./Models/PayLeverageAccruedInterest.md)
 - [PayLeverageAccruedInterestReq](./Models/PayLeverageAccruedInterestReq.md)
 - [PayLeverageAccruedInterestRequest](./Models/PayLeverageAccruedInterestRequest.md)
 - [PayLeverageAccruedInterestResponseEnvelope](./Models/PayLeverageAccruedInterestResponseEnvelope.md)
 - [PoolPrice](./Models/PoolPrice.md)
 - [PoolPriceResponseEnvelope](./Models/PoolPriceResponseEnvelope.md)
 - [PoolRequestError](./Models/PoolRequestError.md)
 - [Portfolio](./Models/Portfolio.md)
 - [Position](./Models/Position.md)
 - [PositionAccount](./Models/PositionAccount.md)
 - [PositionAsset](./Models/PositionAsset.md)
 - [PositionResponse](./Models/PositionResponse.md)
 - [PositionType](./Models/PositionType.md)
 - [PriceLevel](./Models/PriceLevel.md)
 - [ResponseEnvelope](./Models/ResponseEnvelope.md)
 - [ResponseEnvelopeOfListAssets](./Models/ResponseEnvelopeOfListAssets.md)
 - [RevokeAPIKeyData](./Models/RevokeAPIKeyData.md)
 - [RevokeAPIKeyResponseEnvelope](./Models/RevokeAPIKeyResponseEnvelope.md)
 - [Side](./Models/Side.md)
 - [StreamAssetsEntry](./Models/StreamAssetsEntry.md)
 - [StreamCandlesEntry](./Models/StreamCandlesEntry.md)
 - [StreamEntry](./Models/StreamEntry.md)
 - [StreamOrderBookBalanceEntry](./Models/StreamOrderBookBalanceEntry.md)
 - [StreamOrderUpdatesEntry](./Models/StreamOrderUpdatesEntry.md)
 - [StreamOrdersEntry](./Models/StreamOrdersEntry.md)
 - [StreamPositionsEntry](./Models/StreamPositionsEntry.md)
 - [StreamTradesEntry](./Models/StreamTradesEntry.md)
 - [StreamTransactionsEntry](./Models/StreamTransactionsEntry.md)
 - [StreamedAssetPrice](./Models/StreamedAssetPrice.md)
 - [Supply](./Models/Supply.md)
 - [SupplyRequest](./Models/SupplyRequest.md)
 - [SupplyResponseEnvelope](./Models/SupplyResponseEnvelope.md)
 - [Trade](./Models/Trade.md)
 - [TradeRequestError](./Models/TradeRequestError.md)
 - [TradeResponseEnvelope](./Models/TradeResponseEnvelope.md)
 - [Transaction](./Models/Transaction.md)
 - [TransactionKind](./Models/TransactionKind.md)
 - [TransactionRequestError](./Models/TransactionRequestError.md)
 - [TransactionResponseEnvelope](./Models/TransactionResponseEnvelope.md)
 - [TransferBalancesRequest](./Models/TransferBalancesRequest.md)
 - [TransferBalancesResponseEnvelope](./Models/TransferBalancesResponseEnvelope.md)
 - [TransformedAssets](./Models/TransformedAssets.md)
 - [TriggerType](./Models/TriggerType.md)
 - [UnitePositionRequest](./Models/UnitePositionRequest.md)
 - [UnitePositionResponseEnvelope](./Models/UnitePositionResponseEnvelope.md)
 - [UnitedPosition](./Models/UnitedPosition.md)
 - [UpdateFieldString](./Models/UpdateFieldString.md)
 - [UpdateRolesString](./Models/UpdateRolesString.md)
 - [UpdateUserConfigRequest](./Models/UpdateUserConfigRequest.md)
 - [User](./Models/User.md)
 - [UserBalanceResponseEnvelope](./Models/UserBalanceResponseEnvelope.md)
 - [UserConfig](./Models/UserConfig.md)
 - [UserConfigResponseEnvelope](./Models/UserConfigResponseEnvelope.md)
 - [UserDeletedResponseEnvelope](./Models/UserDeletedResponseEnvelope.md)
 - [UserEnvelope](./Models/UserEnvelope.md)
 - [UserInterest](./Models/UserInterest.md)
 - [UserInterestResponseEnvelope](./Models/UserInterestResponseEnvelope.md)
 - [UserPositionResponseEnvelope](./Models/UserPositionResponseEnvelope.md)
 - [UserRole](./Models/UserRole.md)
 - [UserUpdatedResponseEnvelope](./Models/UserUpdatedResponseEnvelope.md)
 - [UserValue](./Models/UserValue.md)
 - [UserValueResponseEnvelope](./Models/UserValueResponseEnvelope.md)
 - [ValidateSubmitOrderRequest](./Models/ValidateSubmitOrderRequest.md)
 - [ValidateSubmitOrderResponse](./Models/ValidateSubmitOrderResponse.md)
 - [Withdraw](./Models/Withdraw.md)
 - [WithdrawRequest](./Models/WithdrawRequest.md)
 - [WithdrawResponseEnvelope](./Models/WithdrawResponseEnvelope.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="bearerAuth"></a>
### bearerAuth

- **Type**: HTTP Bearer Token authentication (JWT)

<a name="apiKeyAuthHeader"></a>
### apiKeyAuthHeader

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

<a name="apiKeyAuthQuery"></a>
### apiKeyAuthQuery

- **Type**: API key
- **API key parameter name**: x-api-key
- **Location**: URL query string

