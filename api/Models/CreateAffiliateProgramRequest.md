# CreateAffiliateProgramRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **tenant\_id** | **String** | Owning tenant for administration. This does not define a referral destination. | [default to null] |
| **name** | **String** | No surrounding whitespace or control characters. | [default to null] |
| **description** | **String** | Omitted or null defaults to an empty description. | [optional] [default to ] |
| **is\_active** | **Boolean** | Inactive programs retain registrations, but cannot issue or resolve codes. Set true to create an active program. | [optional] [default to false] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

