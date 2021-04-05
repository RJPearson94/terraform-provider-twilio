---
page_title: "Twilio Conversations User"
subcategory: "Conversations"
---

# twilio_conversations_user Resource

Manages a conversation user. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_user" "user" {
  service_sid = twilio_conversations_service.service.sid
  identity    = "test-user"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service to associate the user with. Changing this forces a new resource to be created
- `identity` - (Mandatory) The identity of the user. Changing this forces a new resource to be created. The value cannot be an empty string
- `friendly_name` - (Optional) The friendly name of the user. The length of the string must be between `0` and `256` characters (inclusive). The default value is an empty string/ no configuration specified
- `role_sid` - (Optional) The SID of the role to associate with the user
- `attributes` - (Optional) JSON string of user attributes. The default value is `{}`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the user (Same as the `sid`)
- `sid` - The SID of the user (Same as the `id`)
- `account_sid` - The account SID associated with the user
- `service_sid` - The service SID associated with the user
- `identity` - The identity of the user
- `friendly_name` - The friendly name of the user
- `role_sid` - The SID of the role associated with the user
- `attributes` - JSON string of user attributes
- `is_notifiable` - Whether the user can be reached by push notification
- `is_online` - Whether the user has an active connection to the service
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
terraform import twilio_conversations_user.user /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Users/USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
