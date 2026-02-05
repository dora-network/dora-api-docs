# Integrator users and login via API key

This document details the DORA workflow for integrators who want to manage their own user creation
and how they can use API keys generated for their own users to login directly to DORA without having
to use the DORA login screen.

## Pre-requisites

Before you can create a user on DORA, you will need to use the `GET /v1/assets` endpoint to retrieve
a list of available assets from DORA to obtain the asset id to use for the `native_asset_id` in the create
user request. The `native_asset_id` should be a currency asset. Currently only USD is supported. After you
have obtained the assets from DORA, you can obtain the asset ID for USD and use that as the `native_asset_id`
in the request body for the `POST /v1/integrators/user` endpoint.

> NOTE: You can apply a filter to the `GET /v1/assets` endpoint to filter by only currencies. e.g.
> `GET /v1/assets?asset_kind=CURRENCY`.

## Creating users via DORA's API

Integrators can create users on DORA directly using our the endpoint [`POST /v1/integrators/user`](./api/Apis/DefaultApi.md#createUser).

The endpoint takes a body containing a minimum:

```json
{
    "email": "someone@somewhere.com",
    "name": "S. Omeone",
    "native_asset_id": "asset_id_for_usd"
}
```

This will create the user and return a user object containing the details of the user including the user's
unique DORA id. Save this user id as you will need it to create the API key for the user in the next step.

Please note there is a unique email constraint, so the same email cannot be used to create multiple users on
DORA.

## Generating user's API keys

To generate a unique API key for each user, integrators can use the `POST /v1/user/{user_id}/apikey` endpoint
using the unique DORA id that was provided when the user was created.

The request body for this endpoint requires at least:

```json
{
    "label": "SomeLabelForTheApiKey"
}
```

if you wish to set an automatic expiry date for the API key, you may also provide that in the request body:

```json
{
    "label": "SomeLabelForTheApiKey",
    "expires": "2027-01-01T00:00:00Z"
}
```

You will receive a response that includes a DORA generated identifier for the key, the label you chose for
the key, and the actual key itself. Please make sure you record this key as it cannot be provided again by
DORA. You will need to generate a new key if the key is lost.

```json
{
    "key_id": "unique_key_id",
    "label": "your_label",
    "api_key": "secret_key_generated"
}
```

You should save this API key for the user you created as you can use it to log in directly to the DORA UI
without the user having to re-enter any credentials using it.

### Key rotation

For security purposes, it is recommended to rotate the API keys on a regular basis. To do this, you can:

- create an API with the expiry date set as described above, and create a new API key with a new
  expiry date shortly before the old one expires to replace it. The old key will automatically expire on the
  given expiry date.
- use the `PUT /v1/user/{user_id}/apikey/{key_id}/revoke` endpoint to revoke the existing key, after generating
  a new API key. The user_id is the unique DORA `user_id` generated for the user when it was created. The `key_id`
  is the unique key id that DORA assigned the key when it was created, NOT the secret api key that secret api key
  used to authenticate the user.

More details about key generation and management can be found in [the getting started guide](./getting-started.md#api-key)

## Login via API key

To login to the DORA UI automatically, users can use an API key and log in directly by going to:
`https://<your-dora-base-url>/login?apikey=<users-api-key>`

This will initiate user authentication via api key and if the api key is valid, will log the user into
DORA directly without needing to pass further authentication via the DORA login screen.

## Suggested workflow for integrating users with DORA white labelled UI

Below is the recommended workflow for integrators onboarding users to DORA's white labelled UI.

1. User creates an account or already has an account with the integrator
2. Integrator uses DORA's API to create a new user on DORA using the credentials the user already provided to the Integrator
3. Integrator saves the user's DORA user ID along with the user's data on their platform
4. Integrator uses DORA's API to create a DORA API key for the user
5. Integrator saves the API key for the user along with the user's data on their platform
6. Integrator creates a link on their platform to redirect the user to DORA's white labelled solution for the integrator
   1. The link should be unique for each user as it should contain the user's API key as described [above](#login-via-api-key)
7. The user logs onto the Integrator's platform using their credentials
8. The user is presented with their landing page and clicks on their custom link to DORA
9. The custom link containing the user's API key opens up the DORA white labelled UI for the integrator on the API authentication page.
10. DORA authenticates the API key that was provided by the Integrator's custom login link for the user and once verified will launch the DORA UI ready for trading.
