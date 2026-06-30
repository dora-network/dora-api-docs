# AssetConfig
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **asset\_id** | **UUID** |  | [default to null] |
| **price** | **String** | if an asset is a CURRENCY, set 1 USD price,If an asset is a BOND and the price isn&#39;t found, set to 0 USD   You can find price details on /price/asset/{asset_id} route | [default to null] |
| **module\_available** | **String** | Optional leverage module available balance for this asset, from /v1/ledger/module/{asset_id}. If provided, validation rejects orders that need to borrow more than the module can supply. | [optional] [default to null] |
| **module\_supplied** | **String** | Optional leverage module total supplied balance for this asset, from /v1/ledger/module/{asset_id}. Required with module_available when the asset has max_utilization. | [optional] [default to null] |
| **module\_borrowed** | **String** | Optional leverage module borrowed balance for this asset, from /v1/ledger/module/{asset_id}. Required with module_available when the asset has max_utilization. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

