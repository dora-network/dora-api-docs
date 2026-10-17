# PnLRankingSummary
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **all\_users\_daily\_trading\_volume** | **BigDecimal** | Sum of daily_trading_volume across all users matching ranking filters before pagination. | [default to null] |
| **all\_users\_total\_trading\_volume** | **BigDecimal** | Sum of total_trading_volume across all users matching ranking filters before pagination. | [default to null] |
| **all\_users\_avg\_daily\_trading\_volume** | **BigDecimal** | all_users_total_trading_volume divided by the number of distinct UTC days on which any matching user had an executed fill. 0 if there are none. | [default to null] |
| **all\_users\_avg\_daily\_trading\_volume\_on\_n\_days** | **BigDecimal** | Executed trading volume within [start, end) across all users matching ranking filters, divided by the number of requested days elapsed so far (end clamped to now, partial days rounded up). 0 if start is in the future. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

