# PromoLinkAdmin
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **token\_prefix** | **String** |  | [default to null] |
| **url** | **URI** | Only present when reveal&#x3D;true. Private; do not log or cache. | [optional] [default to null] |
| **status** | [**PromoLinkStatus**](PromoLinkStatus.md) |  | [default to null] |
| **expires\_at** | **Date** |  | [default to null] |
| **claimed\_at** | **Date** |  | [optional] [default to null] |
| **claimed\_email** | **String** | Masked as the first character, three asterisks, and domain. | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

