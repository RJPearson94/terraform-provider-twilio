---
page_title: "twilio_messaging_alpha_sender Resource - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_alpha_sender Resource

Manages a Programmable Messaging alphanumeric sender resource. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid  = twilio_messaging_service.service.sid
  alpha_sender = "test"
}
```

## Schema

### Required

- `alpha_sender` (String) The alphanumeric sender ID string. Changing this forces a new resource
- `service_sid` (String) The SID of the messaging service to associate the alpha sender with. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this alpha sender
- `capabilities` (List of String) The list of capabilities for the alpha sender
- `date_created` (String) The date and time the alpha sender was created, in RFC 3339 format
- `date_updated` (String) The date and time the alpha sender was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this alpha sender by Twilio
- `url` (String) The absolute URL of the alpha sender resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A alpha sender can be imported using the `/Services/{serviceSid}/AlphaSenders/{sid}` format, e.g.

```shell
terraform import twilio_messaging_alpha_sender.alpha_sender /Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AlphaSenders/AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
