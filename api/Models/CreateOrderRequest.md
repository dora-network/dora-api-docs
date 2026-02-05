# CreateOrderRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **quantity** | **String** |  | [default to null] |
| **inverse\_leverage** | **String** |  | [default to null] |
| **price** | **String** |  | [optional] [default to null] |
| **kind** | [**OrderKind**](OrderKind.md) |  | [default to null] |
| **side** | [**Side**](Side.md) | Required: Must be either &#39;BUY&#39; or &#39;SELL&#39; | [default to null] |
| **from\_global\_position** | **Boolean** | use global position for the order or isolated. required. | [default to null] |
| **order\_book\_id** | **UUID** | Required: the order book to submit the order to | [default to null] |
| **order\_modifiers** | [**List**](OrderModifierKind.md) |  | [optional] [default to null] |
| **good\_till\_date** | **Date** |  | [optional] [default to null] |
| **trigger\_price** | **String** |  | [optional] [default to null] |
| **trigger\_type** | [**TriggerType**](TriggerType.md) |  | [optional] [default to null] |
| **client\_order\_id** | **String** | An optional client-provided identifier for the order. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

