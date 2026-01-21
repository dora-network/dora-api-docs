# PLAsset
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **symbol** | **String** | The symbol of the asset | [optional] [default to null] |
| **side** | **String** | The side of the position (LONG or SHORT) | [optional] [default to null] |
| **avg\_entry\_price** | **BigDecimal** | The average entry price of the position | [optional] [default to null] |
| **mark\_price** | **BigDecimal** | The current mark price for the asset to calculate daily PL. This is usually the close price of the previous day | [optional] [default to null] |
| **liquidation\_price** | **BigDecimal** | The liquidation price of the position | [optional] [default to null] |
| **available** | **BigDecimal** | The available quantity in units of the asset | [optional] [default to null] |
| **borrowed** | **BigDecimal** | The borrowed quantity in units of the asset | [optional] [default to null] |
| **margin** | [**Margin**](Margin.md) |  | [optional] [default to null] |
| **unrealized\_pl** | **BigDecimal** | The unrealized profit or loss of the position | [optional] [default to null] |
| **leverage\_limit** | **BigDecimal** | The leverage limit for the position | [optional] [default to null] |
| **tp** | **BigDecimal** | The take profit price set for the position, if any | [optional] [default to null] |
| **sl** | **BigDecimal** | The stop loss price set for the position, if any | [optional] [default to null] |
| **initial\_capital** | **BigDecimal** | The initial capital of the position | [optional] [default to null] |
| **impending\_borrows** | **BigDecimal** | The impending borrows of the position | [optional] [default to null] |
| **locked** | **BigDecimal** | The locked amount of the position | [optional] [default to null] |
| **unused\_collateral** | **BigDecimal** | The unused collateral of the position | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

