# AffiliateReferrer
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **program\_id** | **UUID** |  | [default to null] |
| **tenant\_id** | **String** | Owning tenant for administration. This does not define a referral destination. | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **referral\_code** | **String** | Current code to share; previously issued codes remain valid. Stored uppercase and unique across all programs and tenants. | [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **customized\_at** | **Date** | When the owner used their one code change; null until then. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

