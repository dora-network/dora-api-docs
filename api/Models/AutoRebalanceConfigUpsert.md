# AutoRebalanceConfigUpsert
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **account\_id** | **UUID** |  | [default to null] |
| **enabled** | **Boolean** |  | [default to null] |
| **min\_utilization** | **BigDecimal** | Lower utilization threshold, greater than 0 and less than max_utilization. | [default to null] |
| **max\_utilization** | **BigDecimal** | Upper utilization threshold, greater than min_utilization and less than or equal to 1. | [default to null] |
| **order\_rebalance\_enabled** | **Boolean** | Shared opt-in for CFD order funding and automatic USD return. Defaults to false on insert; omission preserves the stored value on update. Both capabilities must be deployed before enabling. | [optional] [default to null] |
| **lowest\_threshold** | **BigDecimal** | Return trigger threshold, greater than 0 and below min_utilization. Required in the effective enabled configuration when order_rebalance_enabled is true; omission preserves a stored threshold. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

