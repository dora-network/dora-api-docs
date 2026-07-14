# PermitMessage
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **owner** | **String** | The user&#39;s wallet address (permit owner), as a 0x-prefixed hex string. | [default to null] |
| **spender** | **String** | The vault contract address granted the allowance, as a 0x-prefixed hex string. | [default to null] |
| **value** | **String** | Approved quantity in USDC base units (10^-6 USDC), as a decimal string. | [default to null] |
| **nonce** | **String** | The owner&#39;s current USDC permit nonce on this chain, as a decimal string. | [default to null] |
| **deadline** | **String** | Unix timestamp (seconds) at which the permit expires, as a decimal string. One hour from issuance. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

