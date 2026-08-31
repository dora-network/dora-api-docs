# UserDeactivation
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **deactivation\_id** | **UUID** |  | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **requested\_by** | **UUID** | Admin that requested the deactivation. | [default to null] |
| **reason** | **String** |  | [default to null] |
| **status** | **String** | PENDING: wind-down in progress. FAILED: wind-down gave up; admin can re-trigger. COMPLETED: account deactivated. REACTIVATED: blocks lifted. | [default to null] |
| **attempts** | **Integer** | Wind-down attempts performed so far. | [default to null] |
| **result** | **String** | Latest wind-down outcome or error summary. | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **updated\_at** | **Date** |  | [default to null] |
| **completed\_at** | **Date** |  | [optional] [default to null] |
| **reactivated\_by** | **UUID** |  | [optional] [default to null] |
| **reactivated\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

