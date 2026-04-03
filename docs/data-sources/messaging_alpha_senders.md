---
page_title: "twilio_messaging_alpha_senders Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_alpha_senders Data Source

Use this data source to access information about the alphanumeric senders associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_alpha_senders" "alpha_senders" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "alpha_senders" {
  value = data.twilio_messaging_alpha_senders.alpha_senders
}
```

## Schema

### Required

- `service_sid` (String) The SID of the messaging service to retrieve alpha senders for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these alpha senders
- `alpha_senders` (List of Object) The list of alpha senders associated with the messaging service (see [below for nested schema](#nestedatt--alpha_senders))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--alpha_senders"></a>
### Nested Schema for `alpha_senders`

Read-Only:

- `alpha_sender` (String)
- `capabilities` (List of String)
- `date_created` (String)
- `date_updated` (String)
- `sid` (String)
- `url` (String)
