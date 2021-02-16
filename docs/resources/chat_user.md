---
page_title: "Twilio Programmable Chat User"
subcategory: "Programmable Chat"
---

# twilio_chat_user Resource

!> This resource is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Manages a Programmable Chat user. See the [API docs](https://www.twilio.com/docs/chat/rest/user-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}

resource "twilio_chat_user" "user" {
  service_sid = twilio_chat_service.service.sid
  identity    = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the user with. Changing this forces a new resource to be created
- `identity` - (Mandatory) The identity of the user. Changing this forces a new resource to be created
- `attributes` - (Optional) JSON string of user attributes
- `friendly_name` - (Optional) The friendly name of the user
- `role_sid` - (Optional) The SID of the role to associate with the user

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

- `create` - (Defaults to 10 minutes) Used when creating the user
- `update` - (Defaults to 10 minutes) Used when updating the user
- `read` - (Defaults to 5 minutes) Used when retrieving the user
- `delete` - (Defaults to 10 minutes) Used when deleting the user

## Import

A user can be imported using the `/Services/{serviceSid}/Users/{sid}` format, e.g.

```shell
terraform import twilio_chat_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
