---
page_title: "twilio_twiml_app Data Source - twilio"
subcategory: "TwiML"
description: |-
  
---

# twilio_twiml_app Data Source

Use this data source to access information about an existing TwiML application. See the [API docs](https://www.twilio.com/docs/usage/api/applications) for more information

## Example Usage

```hcl
data "twilio_twiml_app" "app" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "friendly_name" {
  value = data.twilio_twiml_app.app.friendly_name
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this TwiML application
- `sid` (String) The SID of the TwiML application to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the TwiML application was created, in RFC 3339 format
- `date_updated` (String) The date and time the TwiML application was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the TwiML application
- `id` (String) The ID of this resource.
- `messaging` (List of Object) The messaging settings for the TwiML application (see [below for nested schema](#nestedatt--messaging))
- `voice` (List of Object) The voice settings for the TwiML application (see [below for nested schema](#nestedatt--voice))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--messaging"></a>
### Nested Schema for `messaging`

Read-Only:

- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `status_callback_url` (String)
- `url` (String)


<a id="nestedatt--voice"></a>
### Nested Schema for `voice`

Read-Only:

- `caller_id_lookup` (Boolean)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `status_callback_method` (String)
- `status_callback_url` (String)
- `url` (String)
