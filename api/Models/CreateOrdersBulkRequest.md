# CreateOrdersBulkRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **orders** | [**List**](CreateOrderRequest.md) | Orders to submit, all with the same &#x60;order_book_id&#x60;. The maximum is configurable (matching.bulk_order_max_orders, default 20). | [default to null] |
| **atomic** | **Boolean** | Not currently supported: must be omitted or false, and &#x60;true&#x60; is rejected with 400. Orders are processed independently and partial creation is allowed. | [optional] [default to false] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

