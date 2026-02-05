# ValidateSubmitOrderRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **quantity** | **String** |  | [default to null] |
| **tick** | **String** | Minimum tradable increment for the selected order book | [default to null] |
| **kind** | [**OrderKind**](OrderKind.md) | Must be LIMIT or MARKET | [default to null] |
| **side** | [**Side**](Side.md) | Must be BUY or SELL | [optional] [default to null] |
| **price** | **String** | If kind is LIMIT, must be &gt; 0; if MARKET it must be 0 or omitted | [default to null] |
| **good\_till\_date** | **Date** |  | [optional] [default to null] |
| **inverse\_leverage** | **String** |  | [default to null] |
| **user\_balance** | **String** | User balance used to ensure they can afford the requested quantity | [default to null] |
| **base\_asset\_id** | **UUID** | base asset of orderbook | [optional] [default to null] |
| **quote\_asset\_id** | **UUID** | quote asset of orderbook | [optional] [default to null] |
| **client\_order\_id** | **String** | An optional client-provided identifier for the order. | [optional] [default to null] |
| **position\_assets** | [**List**](PositionAsset.md) | Full list of assets in the position with their price and collateral weight, required when inverse_leverage &lt; 1 for leverage health checks | [optional] [default to null] |
| **assets\_config** | [**List**](AssetConfig.md) | Configuration for the assets in the order | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

