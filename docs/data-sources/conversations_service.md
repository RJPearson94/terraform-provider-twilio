---
page_title: "twilio_conversations_service Data Source - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_service Data Source

Use this data source to access information about an existing conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/service-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_service" "service" {
  sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_conversations_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the conversations service to retrieve

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the conversations service was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversations service was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the conversations service
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the conversations service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
