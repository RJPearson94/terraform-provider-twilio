---
page_title: "twilio_twiml_apps Data Source - twilio"
subcategory: "TwiML"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account to retrieve TwiML applications for

### Optional

- `friendly_name` (String) A friendly name to filter TwiML applications by
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `apps` (List of Object) A list of TwiML applications associated with the account (see [below for nested schema](#nestedatt--apps))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--apps"></a>
### Nested Schema for `apps`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `messaging` (List of Object) (see [below for nested schema](#nestedobjatt--apps--messaging))
- `sid` (String)
- `voice` (List of Object) (see [below for nested schema](#nestedobjatt--apps--voice))

<a id="nestedobjatt--apps--messaging"></a>
### Nested Schema for `apps.messaging`

Read-Only:

- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `status_callback_url` (String)
- `url` (String)


<a id="nestedobjatt--apps--voice"></a>
### Nested Schema for `apps.voice`

Read-Only:

- `caller_id_lookup` (Boolean)
- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `status_callback_method` (String)
- `status_callback_url` (String)
- `url` (String)
