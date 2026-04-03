---
page_title: "twilio_chat_channels Data Source - twilio"
subcategory: "Programmable Chat"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the Programmable Chat service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the channels
- `channels` (List of Object) The list of channels (see [below for nested schema](#nestedatt--channels))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--channels"></a>
### Nested Schema for `channels`

Read-Only:

- `attributes` (String)
- `created_by` (String)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `members_count` (Number)
- `messages_count` (Number)
- `sid` (String)
- `type` (String)
- `unique_name` (String)
- `url` (String)
