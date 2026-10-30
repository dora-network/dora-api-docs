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
| **last\_rebalance\_outcome** | **String** | Most recent risk-rebalance outcome; order funding and return audit entries are excluded. | [optional] [default to null] |
| **order\_rebalance\_enabled** | **Boolean** | Shared order-funding and auto-return opt-in. | [default to null] |
| **order\_funding\_capacity\_usd** | **BigDecimal** | Read-only upper estimate of USD currently fundable into this opted-in CFD account, including conditional/global reserve caps. Omitted while disabled or ineligible. Does not reserve funds or include a specific new order reserve increase; submission replans atomically. | [optional] [default to null] |
| **lowest\_threshold** | **BigDecimal** | Return trigger threshold, when configured. Zero effective debt returns all free USD rather than targeting a positive utilization. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

