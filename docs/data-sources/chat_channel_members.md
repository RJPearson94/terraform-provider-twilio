---
page_title: "Twilio Programmable Chat Channel Members"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_member Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the members associated with an existing Programmable Chat service and channel. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_members" "members" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "members" {
  value = data.twilio_chat_channel_members.members
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the channel members are associated with
- `channel_sid` - (Mandatory) The SID of the channel the members are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `service_sid/channel_sid`
- `account_sid` - The SID of the account the channel members are associated with
- `service_sid` - The SID of the service the channel members are associated with
- `channel_sid` - The SID of the channel the members are associated with
- `members` - A list of `member` blocks as documented below

---

A `member` block supports the following:

- `sid` - The SID of the channel member
- `identity` - The identity of the chat member
- `attributes` - JSON string of member attributes
- `role_sid` - The role SID assignment to the member
- `last_consumed_message_index` - The index of the last message read by the member
- `last_consumption_timestamp` - The timestamp of the last message read by the member
- `date_created` - The date in RFC3339 format that the channel member was created
- `date_updated` - The date in RFC3339 format that the channel member was updated
- `url` - The URL of the channel member

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving channel members
