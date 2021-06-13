---
page_title: "Twilio TwiML App"
subcategory: "TwiML"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the application with. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the application
- `messaging` - (Optional) A `messaging` block as documented below.
- `voice` - (Optional) A `voice` block as documented below.

---

A `messaging` block supports the following:

- `url` - (Optional) The URL which should be called on each incoming message
- `method` - (Optional) The HTTP method that should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method that should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`
- `status_callback_url` (Optional) The URL to POST message status information to

---

A `voice` block supports the following:

- `url` - (Optional) The URL which should be called on each incoming call
- `method` - (Optional) The HTTP method that should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method that should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`
- `caller_id_lookup` - (Optional) Whether caller ID lookup is enabled for the phone number. The default value is `false`
- `status_callback_url` (Optional) The URL to send status information to
- `status_callback_method` (Optional) The HTTP method that should be used to call the status callback URL. Valid values are `GET` or `POST`. The default value is `POST`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the application (Same as the `sid`)
- `sid` - The SID of the application (Same as the `id`)
- `account_sid` - The account SID the application is associated with
- `friendly_name` - The friendly name of the application
- `messaging` - A `messaging` block as documented below.
- `voice` - A `voice` block as documented below.
- `date_created` - The date in RFC3339 format that the application was created
- `date_updated` - The date in RFC3339 format that the application was updated

---

A `messaging` block supports the following:

- `url` - The URL which should be called on each incoming message
- `method` - The HTTP method that should be used to call the URL
- `fallback_url` - The URL which should be called when the URL request fails
- `fallback_method` - The HTTP method that should be used to call the fallback URL
- `status_callback_url` The URL to POST message status information to

---

A `voice` block supports the following:

- `url` - The URL which should be called on each incoming call
- `method` - The HTTP method that should be used to call the URL
- `fallback_url` - The URL which should be called when the URL request fails
- `fallback_method` - The HTTP method that should be used to call the fallback URL
- `caller_id_lookup` - Whether caller ID lookup is enabled for the phone number
- `status_callback_url` - The URL to send status information to
- `status_callback_method` The HTTP method that should be used to call the status callback URL

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the application
- `update` - (Defaults to 10 minutes) Used when updating the application
- `read` - (Defaults to 5 minutes) Used when retrieving the application
- `delete` - (Defaults to 10 minutes) Used when deleting the application

## Import

An application can be imported using the `/Accounts/{applicationSid}/Applications/{sid}` format, e.g.

```shell
terraform import twilio_voice_app.app /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
