---
page_title: "Twilio Programmable Chat Channel"
subcategory: "Programmable Chat"
---

# twilio_chat_channel Resource

Manages a Chat Channel

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name   = "twilio-test"
}

resource "twilio_chat_channel" "channel" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-channel"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID associated with the channel. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the channel
- `unique_name` - (Optional) The unique name of the channel
- `attributes` - (Optional) JSON string of channel attributes
- `type` - (Optional) The type of channel. Valid values are `public` or `private`. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel (Same as the SID)
- `sid` - The SID of the channel (Same as the ID)
- `account_sid` - The Account SID associated with the channel
- `service_sid` - The Service SID associated with the channel
- `friendly_name` - The friendly name of the channel
- `unique_name` - The unique name of the channel
- `attributes` - JSON string of channel attributes
- `type` - The type of channel
- `created_by` - Who created the chat channel
- `members_count` - The number of members currently associated with the channel
- `messages_count` - The number of message currently associated with the channel
- `date_created` - The date in RFC3339 format that the channel was created
- `date_updated` - The date in RFC3339 format that the channel was updated
- `url` - The url of the channel
