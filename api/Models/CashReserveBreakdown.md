# CashReserveBreakdown
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **completed\_pac** | **String** | Completed PAC (partially accrued coupon) obligations the user owes, in USD. | [default to null] |
| **outstanding\_lai** | **String** | Outstanding LAI (leverage accrued interest) the user owes, in USD. | [default to null] |
| **estimated\_fees** | **String** | Estimated trading fees for the current settlement period, capped at a configured fraction (1% by default) of the user&#39;s traded USD volume since 00:00:00 UTC. | [default to null] |
| **borrowed\_portion** | **String** | Configured fraction (10% by default) of the user&#39;s total outstanding borrowed value, in USD. | [default to null] |
| **floor** | **String** | Configured absolute minimum requirement, in USD. | [default to null] |
| **total** | **String** | The amount of USD the user must keep available in their Global Account. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

