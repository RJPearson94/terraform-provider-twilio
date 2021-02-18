---
page_title: "Twilio Programmable Chat Channels"
subcategory: "Programmable Chat"
---

# twilio_chat_channels Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the channels associated with an existing Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channels" "channels" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "channels" {
  value = data.twilio_chat_channels.channels
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the channels are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `service_sid`)
- `account_sid` - The SID of the account the channels are associated with
- `service_sid` - The SID of the service the channels are associated with (Same as the `id`)
- `channels` - A list of `channel` blocks as documented below

---

A `channel` block supports the following:

- `sid` - The SID of the channel
- `friendly_name` - The friendly name of the channel
- `unique_name` - The unique name of the channel
- `attributes` - JSON string of channel attributes
- `type` - The type of channel
- `created_by` - Who created the chat channel
- `members_count` - The number of members that are associated with the channel
- `messages_count` - The number of messages that are associated with the channel
- `date_created` - The date in RFC3339 format that the channel was created
- `date_updated` - The date in RFC3339 format that the channel was updated
- `url` - The URL of the channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving channels
