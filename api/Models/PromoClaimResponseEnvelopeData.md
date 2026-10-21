# PromoClaimResponseEnvelopeData
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **state** | [**PromoClaimState**](PromoClaimState.md) |  | [default to null] |
| **email** | **String** |  | [optional] [default to null] |
| **campaign\_name** | **String** |  | [optional] [default to null] |
| **promo\_credit** | **String** |  | [optional] [default to null] |
| **min\_reward\_amount** | **String** |  | [optional] [default to null] |
| **max\_reward\_amount** | **String** |  | [optional] [default to null] |
| **target\_equity** | **String** |  | [optional] [default to null] |
| **expires\_at** | **Date** |  | [optional] [default to null] |
| **source\_name** | **String** |  | [optional] [default to null] |
| **link\_type** | **String** | Defaults to SINGLE_USE for existing clients. REUSABLE creates one tracked link without a quantity. | [optional] [default to null] |
| **affiliate\_program\_id** | **UUID** |  | [optional] [default to null] |
| **affiliate\_member\_id** | **UUID** | Existing affiliate referrer/member ID, not user ID. | [optional] [default to null] |
| **referral\_code** | **String** | An existing code belonging to the selected member. Required with both IDs for reusable links; bound on the server and revalidated at claim time. | [optional] [default to null] |
| **campaign\_id** | **UUID** |  | [optional] [default to null] |
| **link\_id** | **UUID** |  | [optional] [default to null] |
| **source\_id** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

