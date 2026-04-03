---
page_title: "twilio_conversations_service Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service Resource

Manages a conversation service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the conversations service. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the conversations service was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversations service was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this conversations service by Twilio
- `url` (String) The absolute URL of the conversations service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_conversations_service.service /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
