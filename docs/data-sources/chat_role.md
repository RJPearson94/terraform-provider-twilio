---
page_title: "Twilio Programmable Chat Role"
subcategory: "Programmable Chat"
---

# twilio_chat_role Data Source

Use this data source to access information about an existing Programmable Chat role. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_role" "role" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "role" {
  value = data.twilio_chat_role.role
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the role is associated with
- `sid` - (Mandatory) The SID of the role

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the role (Same as the SID)
- `sid` - The SID of the role (Same as the ID)
- `account_sid` - The account SID associated with the role
- `service_sid` - The service SID associated with the role
- `friendly_name` - The friendly name of the role
- `type` - The type of role
- `permissions` - The list of permissions the role has
- `date_created` - The date in RFC3339 format that the role was created
- `date_updated` - The date in RFC3339 format that the role was updated
- `url` - The url of the role

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the role
