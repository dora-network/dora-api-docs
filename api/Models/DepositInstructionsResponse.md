# DepositInstructionsResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **user\_id** | **UUID** | The authenticated user the instructions are issued for. | [default to null] |
| **owner\_address** | **String** | The wallet address the instructions were issued for, echoed from the request. | [default to null] |
| **quantity** | **BigDecimal** | Human-decimal USDC deposit quantity, echoed from the request. | [default to null] |
| **chains** | [**List**](DepositInstructionForChain.md) |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

