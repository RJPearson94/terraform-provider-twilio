---
page_title: "Twilio Conversations Role"
subcategory: "Conversations"
---

# twilio_conversations_role Resource

Manages a conversation role. See the [API docs](https://www.twilio.com/docs/conversations/api/role-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_conversations_role" "role" {
  service_sid   = twilio_conversations_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "conversation"
  permissions = [
    "sendMediaMessage",
    "sendMessage"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the role with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the role. Changing this forces a new resource to be created
- `type` - (Mandatory) The type of role. Valid values are `service` or `conversation`. Changing this forces a new resource to be created
- `permissions` - (Mandatory) The list of permissions the role has

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the role (Same as the `sid`)
- `sid` - The SID of the role (Same as the `id`)
- `account_sid` - The account SID associated with the role
- `service_sid` - The service SID associated with the role
- `friendly_name` - The friendly name of the role
- `type` - The type of role
- `permissions` - The list of permissions the role has
- `date_created` - The date in RFC3339 format that the role was created
- `date_updated` - The date in RFC3339 format that the role was updated
- `url` - The URL of the role

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the role
- `update` - (Defaults to 10 minutes) Used when updating the role
- `read` - (Defaults to 5 minutes) Used when retrieving the role
- `delete` - (Defaults to 10 minutes) Used when deleting the role

## Import

A role can be imported using the `/Services/{serviceSid}/Roles/{sid}` format, e.g.

```shell
terraform import twilio_conversations_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
