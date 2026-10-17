# AffiliateAttribution
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **referred\_user\_id** | **UUID** |  | [default to null] |
| **program\_id** | **UUID** |  | [default to null] |
| **referrer\_id** | **UUID** |  | [default to null] |
| **referrer\_user\_id** | **UUID** |  | [default to null] |
| **tenant\_id** | **String** | Tenant that owns the referrer and the program. A referrer in a global tenant may refer users of any tenant. | [default to null] |
| **referred\_tenant\_id** | **String** | Tenant that owns the referred user. Differs from tenant_id only when a global tenant&#39;s referrer referred the user. | [default to null] |
| **signup\_source** | **String** | Signup hostname when assigned at signup; empty for an assignment made later. | [default to null] |
| **created\_at** | **Date** | Immutable assignment time; activity begins counting from this timestamp. | [default to null] |
| **referral\_code** | **String** | Exact code the user entered, retained when the referrer changes their primary code. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

