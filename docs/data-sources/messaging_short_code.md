---
page_title: "twilio_messaging_short_code Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_short_code Data Source

Use this data source to access information about an existing Programmable Messaging short code resource. See the [API docs](https://www.twilio.com/docs/sms/services/api/shortcode-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_short_code" "short_code" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_code" {
  value = data.twilio_messaging_short_code.short_code
}
```

## Schema

### Required

- `service_sid` (String) The SID of the messaging service the short code is associated with
- `sid` (String) The SID of the short code

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this short code
- `capabilities` (List of String) The list of capabilities for the short code
- `country_code` (String) The two-character ISO country code of the short code
- `date_created` (String) The date and time the short code was created, in RFC 3339 format
- `date_updated` (String) The date and time the short code was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `short_code` (String) The short code value
- `url` (String) The absolute URL of the short code resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
