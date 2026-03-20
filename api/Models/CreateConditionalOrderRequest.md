# CreateConditionalOrderRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **price** | **BigDecimal** |  | [default to null] |
| **order\_book\_id** | **UUID** | Required: the order book to submit the order to | [default to null] |
| **position\_id** | **UUID** | Required: the position to submit the order to | [default to null] |
| **asset\_id** | **UUID** | Required: the asset to submit the order to | [default to null] |
| **stop\_loss\_price** | **BigDecimal** | Stop loss price | [optional] [default to null] |
| **take\_profit\_price** | **BigDecimal** | Take profit price | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

