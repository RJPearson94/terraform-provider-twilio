---
page_title: "Twilio Conversations User"
subcategory: "Conversations"
---

# twilio_conversations_user Data Source

Use this data source to access information about an existing conversations user. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_user" "user" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "USXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "user" {
  value = data.twilio_conversations_user.user
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the user is associated with
- `sid` - (Mandatory) The SID of the user

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the user (Same as the `sid`)
- `sid` - The SID of the user (Same as the `id`)
- `account_sid` - The account SID associated with the user
- `service_sid` - The service SID associated with the user with
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

- `read` - (Defaults to 5 minutes) Used when retrieving the user
