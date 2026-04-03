---
page_title: "twilio_messaging_short_codes Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_short_codes Data Source

Use this data source to access information about the short codes associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/messaging/services/api/shortcode-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_short_codes" "short_codes" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_codes" {
  value = data.twilio_messaging_short_codes.short_codes
}
```

## Schema

### Required

- `service_sid` (String) The SID of the messaging service to retrieve short codes for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these short codes
- `id` (String) The ID of this resource.
- `short_codes` (List of Object) The list of short codes associated with the messaging service (see [below for nested schema](#nestedatt--short_codes))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--short_codes"></a>
### Nested Schema for `short_codes`

Read-Only:

- `capabilities` (List of String)
- `country_code` (String)
- `date_created` (String)
- `date_updated` (String)
- `short_code` (String)
- `sid` (String)
- `url` (String)
