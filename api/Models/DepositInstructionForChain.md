# DepositInstructionForChain
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **network\_name** | **String** | Human-readable network name, e.g. &#39;Ethereum Mainnet&#39;. | [default to null] |
| **chain\_id** | **String** | EVM chain ID. | [default to null] |
| **contract\_address** | **String** | The DoraUSDCVault contract address for this chain, as a 0x-prefixed hex string. | [default to null] |
| **usdc\_address** | **String** | The ERC-20 USDC token contract on this chain, as a 0x-prefixed hex string. It is the verifying contract of the permit; the spender granted by the permit is contract_address (the vault). | [default to null] |
| **approval** | [**PermitTypedData**](PermitTypedData.md) |  | [default to null] |
| **deposit** | [**DepositCall**](DepositCall.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

