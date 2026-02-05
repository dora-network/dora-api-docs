# Order
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **order\_id** | **UUID** |  | [default to null] |
| **order\_book\_id** | **UUID** |  | [default to null] |
| **kind** | [**OrderKind**](OrderKind.md) |  | [default to null] |
| **original\_price** | **String** | If Kind is LIMIT, this is the original limit price. If Kind is MARKET, this may be 0 or omitted. | [default to null] |
| **avg\_fill\_price** | **String** |  | [default to null] |
| **cancelled\_quantity** | **String** | Quantity that was cancelled, if any. | [default to null] |
| **open\_quantity** | **String** | Quantity that is still open, i.e., not filled or cancelled. | [default to null] |
| **original\_quantity** | **String** | The original quantity of the order when it was created. | [default to null] |
| **filled\_quantity** | **String** | Quantity that has been filled so far. | [default to null] |
| **filled\_notional** | **String** | Quote quantity that has been filled so far. | [default to null] |
| **last\_update\_at** | **Date** |  | [optional] [default to null] |
| **opened\_at** | **Date** |  | [default to null] |
| **inverse\_leverage** | **String** |  | [default to null] |
| **side** | [**Side**](Side.md) |  | [default to null] |
| **status** | [**OrderStatus**](OrderStatus.md) |  | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **order\_modifiers** | [**List**](OrderModifierKind.md) |  | [optional] [default to null] |
| **position\_id** | **UUID** |  | [default to null] |
| **order\_info** | **String** |  | [optional] [default to null] |
| **good\_till\_date** | **Date** |  | [optional] [default to null] |
| **trigger\_price** | **String** |  | [optional] [default to null] |
| **trigger\_type** | [**TriggerType**](TriggerType.md) |  | [optional] [default to null] |
| **client\_order\_id** | **String** | An optional client-provided identifier for the order. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

