---
page_title: "Twilio Programmable Chat Channel"
subcategory: "Programmable Chat"
---

# twilio_chat_channel Data Source

Use this data source to access information about an existing Programmable Chat channel. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel" "channel" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "channel" {
  value = data.twilio_chat_channel.channel
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the channel is associated with
- `sid` - (Mandatory) The SID of the channel

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel (Same as the SID)
- `sid` - The SID of the channel (Same as the ID)
- `account_sid` - The account SID associated with the channel
- `service_sid` - The service SID associated with the channel
- `friendly_name` - The friendly name of the channel
- `unique_name` - The unique name of the channel
- `attributes` - JSON string of channel attributes
- `type` - The type of channel
- `created_by` - Who created the chat channel
- `members_count` - The number of members currently associated with the channel
- `messages_count` - The number of messages that are currently associated with the channel
- `date_created` - The date in RFC3339 format that the channel was created
- `date_updated` - The date in RFC3339 format that the channel was updated
- `url` - The URL of the channel

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the channel
