# DepositResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **network\_chain\_id** | **Integer** | Internal numeric identifier of the chain. | [default to null] |
| **network\_name** | **String** | Human-readable network name. | [default to null] |
| **chain\_id** | **String** | EVM chain ID. | [default to null] |
| **tx\_hash** | **String** | Transaction hash as a 0x-prefixed hex string. | [default to null] |
| **log\_index** | **Integer** | Index of the log within the transaction. | [default to null] |
| **block\_number** | **Long** |  | [default to null] |
| **block\_time** | **Date** |  | [default to null] |
| **contract\_address** | **String** | Vault contract address as a 0x-prefixed hex string. | [default to null] |
| **depositor\_address** | **String** | Address that made the deposit, as a 0x-prefixed hex string. | [default to null] |
| **user\_id** | **UUID** |  | [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC value (base units divided by 10^6). | [default to null] |
| **client\_reference\_id** | **String** | Optional client-supplied reference, as a 0x-prefixed hex string. | [optional] [default to null] |
| **status** | [**Web3EventStatus**](Web3EventStatus.md) |  | [default to null] |
| **transaction\_id** | **UUID** | Internal transaction ID, present once the deposit has been credited. | [optional] [default to null] |
| **observed\_at** | **Date** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

