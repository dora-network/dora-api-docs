# CreateIntegratorUserRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **email** | **String** |  | [optional] [default to null] |
| **first\_name** | **String** |  | [optional] [default to null] |
| **last\_name** | **String** |  | [optional] [default to null] |
| **user\_name** | **String** |  | [optional] [default to null] |
| **country\_of\_domicile** | [**CountryCode**](CountryCode.md) |  | [optional] [default to null] |
| **native\_asset\_id** | **UUID** | Optional: the user&#39;s native asset ID. Must be a CURRENCY asset; defaults to USD. The USDC asset is never allowed for integrator-created users. | [optional] [default to null] |
| **photo\_url** | **String** |  | [optional] [default to null] |
| **provider** | **String** |  | [optional] [default to null] |
| **provider\_id** | **UUID** |  | [optional] [default to null] |
| **timezone** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

