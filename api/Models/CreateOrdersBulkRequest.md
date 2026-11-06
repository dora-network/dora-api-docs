# CreateOrdersBulkRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **orders** | [**List**](CreateOrderRequest.md) | Orders to submit, all with the same &#x60;order_book_id&#x60;. The maximum is configurable (matching.bulk_order_max_orders, default 20). | [default to null] |
| **atomic** | **Boolean** | When true, all orders are submitted in one transaction: either every order is created or none is. The first rejected order fails the whole request with its status and an error prefixed by &#x60;order &lt;index&gt;:&#x60;; a rejection of the batch as a whole, e.g. by the combined minimum cash reserve, has no prefix. On a 409 nothing was created, so a retry cannot duplicate orders: a transaction conflict, reported with the database conflict message and no order prefix, can be retried as is. When false, orders are processed independently and partial creation is allowed. | [optional] [default to false] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

