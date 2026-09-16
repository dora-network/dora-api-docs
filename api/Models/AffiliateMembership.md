# AffiliateMembership
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **program\_id** | **UUID** |  | [default to null] |
| **tenant\_id** | **String** | Owning tenant for administration. This does not define a referral destination. | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **referral\_code** | **String** | Reusable referral code, stored uppercase. Letters, digits, hyphens and underscores are accepted; the first character must be a letter or digit. Matching is case-insensitive. Separate from QR claim tokens. | [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **program\_name** | **String** |  | [default to null] |
| **is\_active** | **Boolean** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

