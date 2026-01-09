# Order
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **order\_id** | **UUID** |  | [optional] [default to null] |
| **order\_book\_id** | **UUID** |  | [optional] [default to null] |
| **kind** | [**OrderKind**](OrderKind.md) |  | [optional] [default to null] |
| **original\_price** | **String** | If Kind is LIMIT, this is the original limit price. If Kind is MARKET, this may be 0 or omitted. | [optional] [default to null] |
| **avg\_fill\_price** | **String** |  | [optional] [default to null] |
| **cancelled\_quantity** | **String** | Quantity that was cancelled, if any. | [optional] [default to null] |
| **open\_quantity** | **String** | Quantity that is still open, i.e., not filled or cancelled. | [optional] [default to null] |
| **original\_quantity** | **String** | The original quantity of the order when it was created. | [optional] [default to null] |
| **filled\_quantity** | **String** | Quantity that has been filled so far. | [optional] [default to null] |
| **filled\_notional** | **String** | Quote quantity that has been filled so far. | [optional] [default to null] |
| **last\_update\_at** | **Date** |  | [optional] [default to null] |
| **opened\_at** | **Date** |  | [optional] [default to null] |
| **inverse\_leverage** | **String** |  | [optional] [default to null] |
| **side** | [**Side**](Side.md) |  | [optional] [default to null] |
| **status** | [**OrderStatus**](OrderStatus.md) |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |
| **order\_modifiers** | [**List**](OrderModifierKind.md) |  | [optional] [default to null] |
| **position\_id** | **UUID** |  | [optional] [default to null] |
| **order\_info** | **String** |  | [optional] [default to null] |
| **good\_till\_date** | **Date** |  | [optional] [default to null] |
| **trigger\_price** | **String** |  | [optional] [default to null] |
| **trigger\_type** | [**TriggerType**](TriggerType.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

