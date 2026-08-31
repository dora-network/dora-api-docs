# UpdateTradingChallengeRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **name** | [**UpdateFieldString**](UpdateFieldString.md) | Trading challenge name. | [optional] [default to null] |
| **type** | [**UpdateFieldString**](UpdateFieldString.md) | CASH or TOURNAMENT. Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **max\_users** | [**UpdateFieldInteger**](UpdateFieldInteger.md) | Must be &gt; 0 and cannot be lowered below the number of users already registered. | [optional] [default to null] |
| **start** | [**UpdateFieldDateTime**](UpdateFieldDateTime.md) | Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **end** | [**UpdateFieldDateTime**](UpdateFieldDateTime.md) | Must be after start and in the future. | [optional] [default to null] |
| **initial\_user\_balance** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt; 0. Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **gold\_prize\_quantity** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt; 0 for a TOURNAMENT challenge. | [optional] [default to null] |
| **silver\_prize\_quantity** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt;&#x3D; 0. | [optional] [default to null] |
| **bronze\_prize\_quantity** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt;&#x3D; 0. | [optional] [default to null] |
| **pnl\_condition** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt;&#x3D; 0. Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **total\_volume\_condition** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt;&#x3D; 0. Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **avg\_daily\_volume\_condition** | [**UpdateFieldDecimal**](UpdateFieldDecimal.md) | Must be &gt;&#x3D; 0. Only updatable while the challenge is PENDING. | [optional] [default to null] |
| **minimum\_equity\_percentage\_condition** | [**UpdateFieldInteger**](UpdateFieldInteger.md) | In the range [0,100). Only updatable while the challenge is PENDING. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

