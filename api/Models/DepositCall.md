# DepositCall
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **to** | **String** | The vault contract address, as a 0x-prefixed hex string. | [default to null] |
| **chain\_id** | **Long** | EVM chain ID, as a JSON number. | [default to null] |
| **value** | **String** | Native currency value to send with the transaction. Always &#39;0&#39;. | [default to null] |
| **function\_signature** | **String** | Canonical Solidity signature of the vault&#39;s permit-aware deposit function: deposit(uint256,bytes16,bytes32,uint256,uint8,bytes32,bytes32). | [default to null] |
| **selector** | **String** | The 4-byte ABI call selector of function_signature, as a 0x-prefixed hex string. | [default to null] |
| **args** | [**DepositArgs**](DepositArgs.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

