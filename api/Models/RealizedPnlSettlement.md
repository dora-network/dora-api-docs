# RealizedPnlSettlement
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** | The ID of the realized PnL settlement | [default to null] |
| **user\_id** | **UUID** | The ID of the user associated with the realized PnL settlement | [default to null] |
| **tenant\_id** | **String** | The ID of the tenant associated with the realized PnL settlement | [default to null] |
| **position\_id** | **UUID** | The ID of the position associated with the realized PnL settlement | [default to null] |
| **order\_id** | **UUID** | The ID of the position-closing order associated with the realized PnL settlement | [default to null] |
| **realized\_usd** | **BigDecimal** | The amount of realized PnL in USD | [default to null] |
| **settled\_at** | **Date** | The timestamp when the realized PnL settlement was settled | [optional] [default to null] |
| **created\_at** | **Date** | The timestamp when the realized PnL settlement was created | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

