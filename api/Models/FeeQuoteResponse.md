# FeeQuoteResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **withdrawal\_id** | **UUID** | The withdrawal this quote was issued for. The quote token is bound to it and cannot be redeemed against any other withdrawal. | [default to null] |
| **to** | **String** | The withdrawal destination address, read from the withdrawal row. | [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC withdrawal quantity, read from the withdrawal row. | [default to null] |
| **fee** | **BigDecimal** | The estimated network fee, in human USDC. | [default to null] |
| **fee\_base\_units** | **String** | The estimated network fee, in micro-USDC base units. | [default to null] |
| **chain\_id** | **String** | EVM chain ID the quote was computed for. | [default to null] |
| **quote\_token** | **String** | Signed, TTL-bound quote token to submit to PUT /v1/web3/withdrawals/{withdrawal_id} so the server can validate the fee it quoted. It names the withdrawal it was issued for. | [default to null] |
| **expires\_at** | **Date** | When the quote token expires. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

