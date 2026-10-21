# IssuePromoLinkBatchRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **source\_type** | [**PromoSourceType**](PromoSourceType.md) |  | [default to null] |
| **source\_id** | **String** |  | [default to null] |
| **source\_name** | **String** |  | [optional] [default to null] |
| **allocation** | **Integer** | Required for SINGLE_USE; omit for REUSABLE. Bounded by max_links_per_batch. Minting reserves no campaign place. | [optional] [default to null] |
| **note** | **String** |  | [optional] [default to null] |
| **expires\_at** | **Date** | Defaults to challenge end and must fall between challenge start and end in the future. | [optional] [default to null] |
| **link\_type** | **String** | Defaults to SINGLE_USE for existing clients. REUSABLE creates one tracked link without a quantity. | [optional] [default to null] |
| **affiliate\_program\_id** | **UUID** |  | [optional] [default to null] |
| **affiliate\_member\_id** | **UUID** | Existing affiliate referrer/member ID, not user ID. | [optional] [default to null] |
| **referral\_code** | **String** | An existing code belonging to the selected member. Required with both IDs for reusable links; bound on the server and revalidated at claim time. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

