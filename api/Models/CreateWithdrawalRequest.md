# CreateWithdrawalRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **withdrawal\_id** | **UUID** | Client-supplied idempotency key (also the on-chain correlation key). Repeating a request with the same withdrawal_id has no additional effect. | [default to null] |
| **to\_address** | **String** | Destination wallet address as a 0x-prefixed hex string. Must not be the zero address. | [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC quantity to withdraw. Must be positive and no finer than USDC&#39;s 6 on-chain decimals. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

