# Bond
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** |  | [default to null] |
| **kind** | [**BondKind**](BondKind.md) |  | [default to null] |
| **coupon\_start\_at** | **Date** |  | [optional] [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **final\_coupon\_at** | **Date** |  | [optional] [default to null] |
| **isin** | **String** |  | [default to null] |
| **issued\_at** | **Date** |  | [default to null] |
| **issuer** | **String** |  | [default to null] |
| **maturity\_at** | **Date** |  | [default to null] |
| **principal\_value** | **BigDecimal** |  | [default to null] |
| **payments\_per\_year** | **Integer** |  | [default to null] |
| **payments\_every** | **Integer** | Coupon payment frequency in nanoseconds (minimum 1000 i.e. 1 microsecond) | [default to null] |
| **next\_coupon\_payment** | **Date** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

