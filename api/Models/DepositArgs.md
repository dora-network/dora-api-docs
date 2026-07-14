# DepositArgs
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **quantity** | **String** | Deposit quantity in USDC base units (10^-6 USDC), as a decimal string. Matches the permit&#39;s value. | [default to null] |
| **user\_id** | **String** | The caller&#39;s user ID as 16 bytes, 0x-prefixed hex. Stored on-chain with the deposit and mapped back to the user by the indexer. | [default to null] |
| **client\_reference\_id** | **String** | The bytes32 client reference as a 0x-prefixed hex string. All zero bytes when not supplied. | [default to null] |
| **deadline** | **String** | Unix timestamp (seconds), as a decimal string. Matches the permit&#39;s deadline exactly. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

