---
page_title: "Twilio Programmable Chat Role"
subcategory: "Programmable Chat"
---

# twilio_chat_role Resource

Manages a Chat Role

## Example Usage

```hcl
resource "twilio_chat_service" "service" {
  unique_name   = "twilio-test"
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

- `service_sid` - (Mandatory) The Service SID associated with the role. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the role. Changing this forces a new resource to be created
- `type` - (Mandatory) The type of role. Valid values are `channel` or `deployment`. Changing this forces a new resource to be created
- `permissions` - (Mandatory) The list of permissions the role has

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the role (Same as the SID)
- `sid` - The SID of the role (Same as the ID)
- `account_sid` - The Account SID associated with the role
- `service_sid` - The Service SID associated with the role
- `friendly_name` - The friendly name of the role
- `type` - The type of role
- `permissions` - The list of permissions the role has
- `date_created` - The date that the role was created
- `date_updated` - The date that the role was updated
- `url` - The url of the role