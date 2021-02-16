---
page_title: "Twilio Programmable Chat Roles"
subcategory: "Programmable Chat"
---

# twilio_chat_roles Data Source

!> This data source is deprecated. Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see <https://www.twilio.com/changelog/programmable-chat-end-of-life> for more information

Use this data source to access information about the roles associated with an existing Programmable Chat service. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_roles" "roles" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "roles" {
  value = data.twilio_chat_roles.roles
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the roles are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `account_sid` - The SID of the account the channels are associated with
- `service_sid` - The SID of the service the channels are associated with
- `roles` - A list of `role` blocks as documented below

---

A `role` block supports the following:

- `sid` - The SID of the role
- `friendly_name` - The friendly name of the role
- `type` - The type of role
- `permissions` - The list of permissions the role has
- `date_created` - The date in RFC3339 format that the role was created
- `date_updated` - The date in RFC3339 format that the role was updated
- `url` - The URL of the role

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving roles
