---
page_title: "Twilio Conversations Users"
subcategory: "Conversations"
---

# twilio_conversations_users Data Source

Use this data source to access information about the users associated with an existing conversations service. See the [API docs](https://www.twilio.com/docs/conversations/api/user-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
data "twilio_conversations_users" "users" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "users" {
  value = data.twilio_conversations_users.users
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the users are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `service_sid`)
- `account_sid` - The SID of the account the users are associated with
- `service_sid` - The SID of the service the users are associated with (Same as the `id`)
- `users` - A list of `user` blocks as documented below

---

A `user` block supports the following:

- `sid` - The SID of the user
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

- `read` - (Defaults to 10 minutes) Used when retrieving users
