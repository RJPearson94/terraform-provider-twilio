---
page_title: "Twilio Programmable Chat Channel Member"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_member Data Source

Use this data source to access information about an existing Programmable Chat channel member. See the [API docs](https://www.twilio.com/docs/chat/rest/member-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_member" "member" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "MBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "member" {
  value = data.twilio_chat_channel_member.member
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the channel member is associated with
- `channel_sid` - (Mandatory) The SID of the channel the member is associated with
- `sid` - (Mandatory) The SID of the channel member

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel member (Same as the SID)
- `sid` - The SID of the channel member (Same as the ID)
- `account_sid` - The account SID associated with the channel member
- `service_sid` - The service SID associated with the channel member
- `channel_sid` - The channel SID associated with the channel member
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

- `read` - (Defaults to 5 minutes) Used when retrieving the channel member
