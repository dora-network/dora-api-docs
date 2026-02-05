# PLAsset
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **symbol** | **String** | The symbol of the asset | [default to null] |
| **side** | **String** | The side of the position (LONG or SHORT) | [default to null] |
| **avg\_entry\_price** | **String** | The average entry price of the position | [default to null] |
| **mark\_price** | **String** | The current mark price for the asset to calculate daily PL. This is usually the close price of the previous day | [default to null] |
| **liquidation\_price** | **String** | The liquidation price of the position | [default to null] |
| **available** | **String** | The available quantity in units of the asset | [default to null] |
| **borrowed** | **String** | The borrowed quantity in units of the asset | [default to null] |
| **margin** | [**Margin**](Margin.md) |  | [default to null] |
| **unrealized\_pl** | **String** | The unrealized profit or loss of the position | [default to null] |
| **leverage\_limit** | **String** | The leverage limit for the position | [default to null] |
| **tp** | **String** | The take profit price set for the position, if any | [optional] [default to null] |
| **sl** | **String** | The stop loss price set for the position, if any | [optional] [default to null] |
| **initial\_capital** | **String** | The initial capital of the position | [default to null] |
| **impending\_borrows** | **String** | The impending borrows of the position | [optional] [default to null] |
| **locked** | **String** | The locked amount of the position | [default to null] |
| **unused\_collateral** | **String** | The unused collateral of the position | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

