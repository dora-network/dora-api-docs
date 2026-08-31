# TradingChallengeRegistrationRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **trading\_challenge\_id** | **UUID** |  | [default to null] |
| **trading\_challenge\_name** | **String** |  | [optional] [default to null] |
| **trading\_challenge\_type** | **String** |  | [optional] [default to null] |
| **trading\_challenge\_status** | **String** |  | [optional] [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **user\_email** | **String** |  | [optional] [default to null] |
| **user\_name** | **String** |  | [optional] [default to null] |
| **tenant\_id** | **String** |  | [default to null] |
| **status** | **String** |  | [default to null] |
| **reviewed\_by** | **UUID** | Who settled the request. Absent while it is PENDING. | [optional] [default to null] |
| **reviewed\_at** | **Date** | When it was settled. Absent while it is PENDING. | [optional] [default to null] |
| **review\_reason** | **String** | Free-text note kept for the audit trail. Optional on both decisions. | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **updated\_at** | **Date** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

