---
page_title: "Twilio TwiML Apps"
subcategory: "TwiML"
---

# twilio_twiml_apps Data Source

Use this data source to access information about the TwiML applications associated with an existing account. See the [API docs](https://www.twilio.com/docs/usage/api/applications) for more information

## Example Usage

```hcl
data "twilio_twiml_apps" "apps" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "apps" {
  value = data.twilio_twiml_apps.apps
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the applications are associated with
- `friendly_name` - (Optional) Search for all applications which have the friendly name specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `account_sid`)
- `account_sid` - The account SID associated with the applications (Same as the `id`)
- `apps` - A list of `app` blocks as documented below

---

A `app` block supports the following:

- `sid` - The SID of the application
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

- `read` - (Defaults to 10 minutes) Used when retrieving applications
