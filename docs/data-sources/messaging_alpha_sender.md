---
page_title: "twilio_messaging_alpha_sender Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_alpha_sender Data Source

Use this data source to access information about an existing Programmable Messaging alphanumeric sender. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "alpha_sender" {
  value = data.twilio_messaging_alpha_sender.alpha_sender
}
```

## Schema

### Required

- `service_sid` (String) The SID of the messaging service the alpha sender is associated with
- `sid` (String) The SID of the alpha sender

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this alpha sender
- `alpha_sender` (String) The alphanumeric sender ID string
- `capabilities` (List of String) The list of capabilities for the alpha sender
- `date_created` (String) The date and time the alpha sender was created, in RFC 3339 format
- `date_updated` (String) The date and time the alpha sender was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the alpha sender resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
