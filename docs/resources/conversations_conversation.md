---
page_title: "twilio_conversations_conversation Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_conversation Resource

Manages a Twilio Conversations conversation resource. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

!> Removing the `messaging_service_sid` from your configuration will cause the `messaging_service_sid` value to be retained after a Terraform apply. If you want to change any of the value you will need to either create a new `twilio_messaging_service` resource and set the argument to the generated `sid` or you can set the argument to the `default_messaging_service_sid` which is on the `twilio_conversations_configuration` resource and data source

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}
```

## Schema

### Required

- `service_sid` (String) The SID of the conversations service. Changing this forces a new resource

### Optional

- `attributes` (String) A JSON string of attributes associated with the conversation. Defaults to `{}`
- `friendly_name` (String) A human-readable label for the conversation
- `messaging_service_sid` (String) The SID of the messaging service associated with the conversation
- `state` (String) The state of the conversation. Valid values are `active`, `inactive`, or `closed`. Defaults to `active`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `timers` (Block List, Max: 1) Timer settings for the conversation (see [below for nested schema](#nestedblock--timers))
- `unique_name` (String) A unique, developer-assigned name for the conversation

### Read-Only

- `account_sid` (String) The SID of the account that owns this conversation
- `date_created` (String) The date and time the conversation was created, in RFC 3339 format
- `date_updated` (String) The date and time the conversation was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this conversation by Twilio
- `url` (String) The absolute URL of the conversation resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedblock--timers"></a>
### Nested Schema for `timers`

Optional:

- `closed` (String) ISO 8601 duration after which the conversation will be automatically closed
- `inactive` (String) ISO 8601 duration after which the conversation will be marked as inactive

Read-Only:

- `date_closed` (String) The date and time the conversation will be closed, in RFC 3339 format
- `date_inactive` (String) The date and time the conversation will be marked as inactive, in RFC 3339 format

## Import

A conversation can be imported using the `/Services/{serviceSid}/Conversations/{sid}` format, e.g.

```shell
terraform import twilio_conversations_conversation.conversation /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments `timers.0.closed` and `timers.0.inactive` cannot be imported
