# CashReserveResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **enforced** | **Boolean** | Whether the minimum cash reserve guard is active in this environment. | [default to null] |
| **available\_usd** | **String** | The user&#39;s currently available USD balance in their Global Account. | [default to null] |
| **committed\_usd** | **String** | USD still counted in available_usd but already claimed by the user&#39;s open market buy orders on the Global Account, which reserve no funds at submission time. The reserve is evaluated against available_usd minus committed_usd. | [default to null] |
| **required\_usd** | **String** | The user&#39;s minimum USD cash reserve requirement. | [default to null] |
| **satisfied** | **Boolean** | Whether available_usd minus committed_usd is at least required_usd. | [default to null] |
| **breakdown** | [**CashReserveBreakdown**](CashReserveBreakdown.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

