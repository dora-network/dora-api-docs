# CreateIntegratorUserRequest
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **referral\_code** | **String** | Optional affiliate code, normalized to uppercase. Accepted only when creating a new account in the program owning tenant. One immutable attribution per user account; a later signup/linking call cannot add or replace it. Invalid or inactive codes fail signup atomically. Independent of QR acquisition attribution. Existing unassigned users can instead use POST /v1/affiliate_referrals/self; earlier activity is excluded. | [optional] [default to null] |
| **signup\_source** | **URI** | Optional client-reported HTTP(S) signup site URL, used only with referral_code. When omitted, a valid HTTP(S) Origin header is used; other origins are ignored. Only the hostname is stored, without path, query, credentials or fragment. Unknown if neither supplies a usable hostname. Ignored when referral_code is empty. It does not select or authenticate the tenant. | [optional] [default to null] |
| **email** | **String** |  | [optional] [default to null] |
| **first\_name** | **String** |  | [optional] [default to null] |
| **last\_name** | **String** |  | [optional] [default to null] |
| **user\_name** | **String** |  | [optional] [default to null] |
| **country\_of\_domicile** | [**CountryCode**](CountryCode.md) |  | [optional] [default to null] |
| **native\_asset\_id** | **UUID** | Optional: the user&#39;s native asset ID. Must be a CURRENCY asset; defaults to USD. The USDC asset is never allowed for integrator-created users. | [optional] [default to null] |
| **photo\_url** | **String** |  | [optional] [default to null] |
| **provider** | **String** |  | [optional] [default to null] |
| **provider\_id** | **UUID** |  | [optional] [default to null] |
| **timezone** | **String** |  | [optional] [default to null] |
| **challenge\_id** | **UUID** | Optional: sign the new user up for this trading challenge. This creates a PENDING registration request that an admin, the tenant&#39;s integrator or one of the challenge&#39;s managers must approve before the user is actually enrolled. The challenge must belong to the new user&#39;s tenant and still be open for entries, otherwise the whole sign-up fails. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

