---
page_title: "Twilio Programmable Chat Role"
subcategory: "Programmable Chat"
---

# twilio_chat_role Resource

Manages a Programmable Chat role. See the [API docs](https://www.twilio.com/docs/chat/rest/role-resource) for more information

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name = "twilio-test"
}

resource "twilio_chat_role" "role" {
  service_sid   = twilio_chat_service.service.sid
  friendly_name = "twilio-test-role"
  type          = "channel"
  permissions = [
    "sendMessage",
    "leaveChannel"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID associated with the role. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the role. Changing this forces a new resource to be created
- `type` - (Mandatory) The type of role. Valid values are `channel` or `deployment`. Changing this forces a new resource to be created
- `permissions` - (Mandatory) The list of permissions the role has

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

- `create` - (Defaults to 10 minutes) Used when creating the role
- `update` - (Defaults to 10 minutes) Used when updating the role
- `read` - (Defaults to 5 minutes) Used when retrieving the role
- `delete` - (Defaults to 10 minutes) Used when deleting the role

## Import

A role can be imported using the `/Services/{serviceSid}/Roles/{sid}` format, e.g.

```shell
terraform import twilio_chat_role.role /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Roles/RLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
