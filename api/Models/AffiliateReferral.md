# AffiliateReferral
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **user\_id** | **UUID** |  | [default to null] |
| **program\_id** | **UUID** |  | [default to null] |
| **referrer\_id** | **UUID** |  | [default to null] |
| **referrer\_user\_id** | **UUID** |  | [default to null] |
| **signup\_source** | **String** | Client-reported signup hostname. Empty means unknown. | [default to null] |
| **first\_name** | **String** |  | [default to null] |
| **last\_name** | **String** |  | [default to null] |
| **email** | **String** |  | [default to null] |
| **signed\_up\_at** | **Date** |  | [default to null] |
| **kyc\_completed** | **Boolean** |  | [default to null] |
| **kyc\_completed\_at** | **Date** |  | [default to null] |
| **discord\_status** | **String** | No Discord membership integration is currently available. Unknown must not be interpreted as not joined. | [default to null] |
| **deposit\_count** | **Long** |  | [default to null] |
| **withdrawal\_count** | **Long** |  | [default to null] |
| **has\_traded** | **Boolean** |  | [default to null] |
| **first\_deposit\_at** | **Date** |  | [default to null] |
| **last\_deposit\_at** | **Date** |  | [default to null] |
| **first\_withdrawal\_at** | **Date** |  | [default to null] |
| **last\_withdrawal\_at** | **Date** |  | [default to null] |
| **daily\_volume\_usd** | **String** | Sum of absolute executed FILL quantity1 on USD-quoted trades during the selected UTC day. Both buy and sell executions count, once per user-side fill. | [default to null] |
| **monthly\_volume\_usd** | **String** | Same executed USD quote-notional definition for the calendar month containing date. | [default to null] |
| **daily\_realized\_pnl\_usd** | **String** | Sum of realized_pnl_settlements.realized_usd created during the selected UTC day, matching the existing PnL ranking convention. Excludes unrealized PnL; this is not total account equity change. | [default to null] |
| **attributed\_at** | **Date** | Immutable referral assignment time. Earlier activity is excluded from affiliate metrics and cash flows. | [default to null] |
| **monthly\_realized\_pnl\_usd** | **String** | Realized PnL for the UTC calendar month containing date, including only settlements at or after attributed_at. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

