# Position
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** | The unique identifier for the position. Used, for example, when creating an order from a position, or deciding collateral should be transferred from position A to position B. | [optional] [default to null] |
| **asset\_id** | **UUID** |  | [optional] [default to null] |
| **seq** | **Integer** |  | [optional] [default to null] |
| **is\_global** | **Boolean** |  | [optional] [default to null] |
| **available** | **String** | The available balance in the position for this asset that are not locked, supplied, or used as collateral | [optional] [default to null] |
| **locked** | **String** | The balance that has been reserved for a current order. If spent by the order, they are removed. If the order is cancelled, they are returned to the position&#39;s available balance. | [optional] [default to null] |
| **supplied** | **String** | The balance that user has supplied to the leverage module. The user remains entitled to these assets and can withdraw them into their available balance. | [optional] [default to null] |
| **borrowed** | **String** | The total amount of debt outstanding for this position. This position cannot be closed until all debt is fully repaid, i.e. borrowed &#x3D; 0. | [optional] [default to null] |
| **impending\_borrows** | **String** | The equivalent of locked balances, but for leveraged orders. If a user has an active order that would borrow assets as part of its input, then their borrow limit must be reduced until the order is executed or cancelled. | [optional] [default to null] |
| **avg\_entry\_price** | **String** | average cost per unit quantity that was paid (long positions) or received (short positions) for this asset. | [optional] [default to null] |
| **borrow\_limit** | **String** | The borrow limit | [optional] [default to null] |
| **liquidation\_threshold** | **String** | The borrow limit | [optional] [default to null] |
| **created\_at** | **Date** |  | [optional] [default to null] |
| **position\_name** | **String** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

