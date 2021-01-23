---
page_title: "Twilio Conversations Conversation"
subcategory: "Conversations"
---

# twilio_conversations_conversation Resource

Manages a Twilio Conversations conversation resource. See the [API docs](https://www.twilio.com/docs/conversations/api/conversation-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_conversation" "conversation" {
  service_sid = twilio_conversations_service.service.sid
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the conversation with. Changing this forces a new resource to be created
- `unique_name` - (Optional) The unique name of the conversation
- `friendly_name` - (Optional) The friendly name of the conversation
- `attributes` - (Optional) JSON string of attributes
- `messaging_service_sid` - (Optional) The messaging service SID to associate the conversation with
- `state` - (Optional) The state of the conversation. Valid values are `active`, `inactive` or `closed`.
- `timers` - (Optional) A `timer` block as documented below.

---

A `timer` block supports the following:

- `closed` - (Optional) ISO8601 duration before a conversation will be marked as closed
- `inactive` - (Optional) ISO8601 duration before a conversation will be marked as inactive

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the conversation (Same as the SID)
- `sid` - The SID of the conversation (Same as the ID)
- `account_sid` - The account SID associated with the conversation
- `service_sid` - The service SID associated with the conversation
- `unique_name` - The unique name of the conversation
- `friendly_name` - The friendly name of the conversation
- `attributes` - JSON string of attributes
- `messaging_service_sid` - The messaging service SID associated with the conversation
- `state` - The state of the conversation
- `timers` - A `timer` block as documented below
- `date_created` - The date in RFC3339 format that the conversation was created
- `date_updated` - The date in RFC3339 format that the conversation was updated
- `url` - The URL of the conversation

---

A `timer` block supports the following:

- `date_closed` - The date in RFC3339 format that the conversation will close
- `date_inactive` - The date in RFC3339 format that the conversation will be marked as inactive

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the conversation
- `update` - (Defaults to 10 minutes) Used when updating the conversation
- `read` - (Defaults to 5 minutes) Used when retrieving the conversation
- `delete` - (Defaults to 10 minutes) Used when deleting the conversation

## Import

A conversation can be imported using the `/Services/{serviceSid}/Conversations/{sid}` format, e.g.

```shell
terraform import twilio_conversations_conversation.conversation /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Conversations/CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments `timers.closed` and `timers.inactive` cannot be imported
