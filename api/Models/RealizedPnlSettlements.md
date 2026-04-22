# RealizedPnlSettlements
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **settlements** | [**List**](RealizedPnlSettlement.md) | A list of realized PnL settlements matching the query parameters of the request | [optional] [default to null] |
| **user\_totals** | **Map** | A map of user IDs to their total realized PnL in USD across all settlements included in the response | [optional] [default to null] |
| **tenant\_totals** | **Map** | A map of tenant IDs to their total realized PnL in USD across all settlements included in the response | [optional] [default to null] |
| **user\_totals\_unsettled** | **Map** | A map of user IDs to their total realized PnL in USD across unsettled settlements (where settled_at is null) included in the response | [optional] [default to null] |
| **tenant\_totals\_unsettled** | **Map** | A map of tenant IDs to their total realized PnL in USD across unsettled settlements (where settled_at is null) included in the response | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

