# FeeQuoteResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **to** | **String** | The withdrawal destination address, echoed from the request. | [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC withdrawal quantity, echoed from the request. | [default to null] |
| **fee** | **BigDecimal** | The estimated network fee, in human USDC. | [default to null] |
| **fee\_base\_units** | **String** | The estimated network fee, in micro-USDC base units. | [default to null] |
| **chain\_id** | **String** | EVM chain ID the quote was computed for. | [default to null] |
| **quote\_token** | **String** | Signed, TTL-bound quote token to submit with a later withdrawal so the server can validate the fee it was quoted. | [default to null] |
| **expires\_at** | **Date** | When the quote token expires. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

