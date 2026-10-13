# AutoRebalanceConfigStatus
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **account\_id** | **UUID** |  | [default to null] |
| **enabled** | **Boolean** |  | [default to null] |
| **min\_utilization** | **BigDecimal** |  | [default to null] |
| **max\_utilization** | **BigDecimal** |  | [default to null] |
| **available\_usd** | **BigDecimal** |  | [optional] [default to null] |
| **borrowed\_value** | **BigDecimal** |  | [optional] [default to null] |
| **borrow\_limit** | **BigDecimal** |  | [optional] [default to null] |
| **current\_utilization** | **BigDecimal** |  | [optional] [default to null] |
| **above\_max** | **Boolean** | True when current_utilization is greater than max_utilization. | [optional] [default to null] |
| **in\_active\_liquidation** | **Boolean** |  | [optional] [default to null] |
| **last\_rebalance\_outcome** | **String** | Most recent persisted auto-rebalance attempt status for this account, when available. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

