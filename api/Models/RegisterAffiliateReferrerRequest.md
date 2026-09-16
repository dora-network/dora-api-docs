# RegisterAffiliateReferrerRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **user\_id** | **UUID** | Existing, nonzero DORA user ID in the program owning tenant. | [default to null] |
| **referral\_code** | **String** | Optional custom code: 3 to 64 letters, digits, hyphens or underscores after trimming, starting with a letter or digit. Stored uppercase and unique across all programs and tenants. Omitted, null or empty generates a random code. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

