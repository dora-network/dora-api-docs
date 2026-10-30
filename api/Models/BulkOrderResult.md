# BulkOrderResult
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **index** | **Integer** | Position of the order in the request | [default to null] |
| **client\_order\_id** | **String** |  | [optional] [default to null] |
| **order\_id** | **UUID** | Present when the order was created | [optional] [default to null] |
| **status** | **Integer** | HTTP status the order would have received from POST /v1/orders: 201 when created | [default to null] |
| **error** | **String** | Present when the order was rejected | [optional] [default to null] |
| **funding** | [**OrderFundingResult**](OrderFundingResult.md) | Present only when this committed order received automatic USD funding. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

