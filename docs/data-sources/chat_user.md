---
page_title: "Twilio Programmable Chat User"
subcategory: "Programmable Chat"
---

# twilio_chat_user Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about an existing Programmable Chat user. See the [API docs](https://www.twilio.com/docs/chat/rest/user-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_user" "role" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "user" {
  value = data.twilio_chat_user.user
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the user is associated with
- `sid` - (Mandatory) The SID of the user

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the user (Same as the SID)
- `sid` - The SID of the user (Same as the ID)
- `account_sid` - The account SID associated with the user
- `service_sid` - The service SID associated with the user
- `friendly_name` - The friendly name of the user
- `attributes` - JSON string of user attributes
- `identity` - The identity of the user
- `is_notifiable` - Whether the user can be reached by push notification
- `is_online` - Whether the user has an active connection to the service
- `joined_channels_count` - The number of channels the user has joined
- `role_sid` - The SID of the role associated with the user
- `date_created` - The date in RFC3339 format that the user was created
- `date_updated` - The date in RFC3339 format that the user was updated
- `url` - The URL of the user

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the user
