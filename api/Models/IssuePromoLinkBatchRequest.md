# IssuePromoLinkBatchRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **source\_type** | [**PromoSourceType**](PromoSourceType.md) |  | [default to null] |
| **source\_id** | **String** |  | [default to null] |
| **source\_name** | **String** |  | [optional] [default to null] |
| **allocation** | **Integer** | Cannot exceed the configured max_links_per_batch or remaining campaign capacity. | [default to null] |
| **note** | **String** |  | [optional] [default to null] |
| **expires\_at** | **Date** | Defaults to challenge end and must fall between challenge start and end in the future. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

