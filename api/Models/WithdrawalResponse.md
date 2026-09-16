# WithdrawalResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **withdrawal\_id** | **UUID** |  | [optional] [default to null] |
| **network\_chain\_id** | **Integer** | Internal numeric identifier of the chain. | [optional] [default to null] |
| **network\_name** | **String** | Human-readable network name. | [optional] [default to null] |
| **chain\_id** | **String** | EVM chain ID. | [optional] [default to null] |
| **user\_id** | **UUID** |  | [optional] [default to null] |
| **account\_id** | **UUID** |  | [optional] [default to null] |
| **to\_address** | **String** | Destination wallet address as a 0x-prefixed hex string. | [optional] [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC quantity to withdraw (base units divided by 10^6). | [optional] [default to null] |
| **fee** | **BigDecimal** | Human-decimal USDC network fee (base units divided by 10^6). 0 until the requester locks a quoted fee as part of approval. | [optional] [default to null] |
| **status** | [**Web3WithdrawalStatus**](Web3WithdrawalStatus.md) |  | [optional] [default to null] |
| **tx\_hash** | **String** | Broadcast withdraw() transaction hash as a 0x-prefixed hex string. Present from &#x60;BROADCAST&#x60; onward. | [optional] [default to null] |
| **failure\_reason** | **String** | Reason the withdrawal was rejected or failed. Present for REJECTED/FAILED. | [optional] [default to null] |
| **approved\_by** | **UUID** | Admin who approved the withdrawal. Present once approved. | [optional] [default to null] |
| **approved\_at** | **Date** | When the withdrawal was approved. Present once approved. | [optional] [default to null] |
| **settlement\_transaction\_id** | **UUID** | Ledger settlement transaction. Present once confirmed. | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **updated\_at** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

