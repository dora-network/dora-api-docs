# AccountV2
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **UUID** | The unique identifier for the account. Used, for example, when creating an order from an account, or deciding collateral should be transferred from account A to account B. | [default to null] |
| **asset\_id** | **UUID** |  | [default to null] |
| **seq** | **Integer** |  | [default to null] |
| **is\_global** | **Boolean** |  | [optional] [default to null] |
| **available** | **String** | The available balance in the account for this asset that are not locked, supplied, or used as collateral | [default to null] |
| **locked** | **String** | The balance that has been reserved for a current order. If spent by the order, they are removed. If the order is cancelled, they are returned to the account&#39;s available balance. | [default to null] |
| **supplied** | **String** | The balance that user has supplied to the leverage module. The user remains entitled to these assets and can withdraw them into their available balance. | [default to null] |
| **borrowed** | **String** | The total amount of debt outstanding for this account. This account cannot be closed until all debt is fully repaid, i.e. borrowed &#x3D; 0. | [default to null] |
| **impending\_borrows** | **String** | The equivalent of locked balances, but for leveraged orders. If a user has an active order that would borrow assets as part of its input, then their borrow limit must be reduced until the order is executed or cancelled. | [default to null] |
| **avg\_entry\_price** | **String** | average cost per unit quantity that was paid (long accounts) or received (short accounts) for this asset. | [default to null] |
| **borrow\_limit** | **String** | The borrow limit | [default to null] |
| **liquidation\_threshold** | **String** | The borrow limit | [default to null] |
| **created\_at** | **Date** |  | [default to null] |
| **pending\_withdrawal** | **String** | The amount of asset that is pending withdrawal from the account. | [default to null] |
| **account\_name** | **String** |  | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

