---
page_title: "twilio_twiml_app Resource - twilio"
subcategory: "TwiML"
description: |-
  
---

# twilio_twiml_app Resource

Manages a TwiML application. See the [API docs](https://www.twilio.com/docs/usage/api/applications) for more information

!> During testing it was noticed that removing the `messaging.0.url` or `voice.0.url` from your configuration will cause the corresponding value to be retained after a Terraform apply. This does not affect updating either of the URLs

## Example Usage

### With Account SID

```hcl
resource "twilio_twiml_app" "app" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

### With `twilio_account_details` data source

```hcl
data "twilio_account_details" "account_details" {}

resource "twilio_twiml_app" "app" {
  account_sid = data.twilio_account_details.account_details.sid
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this TwiML application. Changing this forces a new resource

### Optional

- `friendly_name` (String) A human-readable label for the TwiML application
- `messaging` (Block List, Max: 1) A block to configure messaging settings for the TwiML application (see [below for nested schema](#nestedblock--messaging))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `voice` (Block List, Max: 1) A block to configure voice settings for the TwiML application (see [below for nested schema](#nestedblock--voice))

### Read-Only

- `date_created` (String) The date and time the TwiML application was created, in RFC 3339 format
- `date_updated` (String) The date and time the TwiML application was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this TwiML application by Twilio

<a id="nestedblock--messaging"></a>
### Nested Schema for `messaging`

Optional:

- `fallback_method` (String) The HTTP method used to call the messaging fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs retrieving or executing the TwiML for incoming messages
- `method` (String) The HTTP method used to call the messaging URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_url` (String) The URL to call for messaging status callback events
- `url` (String) The URL to call when the application receives an incoming message


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedblock--voice"></a>
### Nested Schema for `voice`

Optional:

- `caller_id_lookup` (Boolean) Whether to perform a caller ID lookup on incoming voice calls. Defaults to `false`
- `fallback_method` (String) The HTTP method used to call the voice fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs retrieving or executing the TwiML for incoming voice calls
- `method` (String) The HTTP method used to call the voice URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_method` (String) The HTTP method used to call the voice status callback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_url` (String) The URL to call for voice status callback events
- `url` (String) The URL to call when the application receives an incoming voice call

## Import

An application can be imported using the `/Accounts/{applicationSid}/Applications/{sid}` format, e.g.

```shell
terraform import twilio_voice_app.app /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
