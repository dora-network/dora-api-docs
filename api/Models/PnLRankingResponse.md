# PnLRankingResponse
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **user\_id** | **UUID** |  | [default to null] |
| **first\_name** | **String** |  | [default to null] |
| **total\_pnl** | **BigDecimal** |  | [default to null] |
| **realized\_pnl** | **BigDecimal** | Cumulative realized PnL across the user&#39;s full trading lifetime. | [default to null] |
| **total\_trades** | **Integer** |  | [default to null] |
| **winning\_trades** | **Integer** |  | [default to null] |
| **losing\_trades** | **Integer** |  | [default to null] |
| **win\_rate** | **BigDecimal** |  | [default to null] |
| **daily\_trading\_volume** | **BigDecimal** | Executed trading volume for the current UTC day. | [default to null] |
| **total\_trading\_volume** | **BigDecimal** | Cumulative executed trading volume across all UTC trading days. | [default to null] |
| **avg\_daily\_volume** | **BigDecimal** | total_trading_volume divided by active_trading_days. 0 if the user has no active trading days. | [default to null] |
| **avg\_daily\_volume\_on\_n\_days** | **BigDecimal** | Executed trading volume within [start, end) divided by the number of requested days elapsed so far (end clamped to now, partial days rounded up). 0 if start is in the future. | [default to null] |
| **active\_trading\_days** | **Integer** | Number of distinct UTC days on which the user has at least one executed fill. | [default to null] |
| **activated** | **Boolean** | True once the user has traded on at least 5 distinct UTC days. | [default to null] |
| **kyc\_approved** | **Boolean** | Whether the user is currently considered KYC/verification approved. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

