# PromoLinkBatchSummary
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **source\_type** | [**PromoSourceType**](PromoSourceType.md) |  | [default to null] |
| **source\_id** | **String** |  | [default to null] |
| **source\_name** | **String** |  | [default to null] |
| **allocation** | **Integer** |  | [optional] [default to null] |
| **issued** | **Long** |  | [optional] [default to null] |
| **claimed** | **Long** |  | [optional] [default to null] |
| **revoked** | **Long** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **link\_type** | **String** | Defaults to SINGLE_USE for existing clients. REUSABLE creates one tracked link without a quantity. | [optional] [default to null] |
| **affiliate\_program\_id** | **UUID** |  | [optional] [default to null] |
| **affiliate\_member\_id** | **UUID** | Existing affiliate referrer/member ID, not user ID. | [optional] [default to null] |
| **referral\_code** | **String** | An existing code belonging to the selected member. Required with both IDs for reusable links; bound on the server and revalidated at claim time. | [optional] [default to null] |
| **campaign\_id** | **UUID** |  | [optional] [default to null] |
| **signups** | **Long** | Successful signups; reusable links omit allocation/issued/claimed/revoked counts. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

