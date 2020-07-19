---
page_title: "Twilio Autopilot Field Type"
subcategory: "Autopilot"
---

# twilio_autopilot_field_type Resource

Manages a Autopilot Field Type

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
  friendly_name = "test"
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the field type to. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the field type. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the field type

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field type (Same as the SID)
- `sid` - The SID of the field type (Same as the ID)
- `account_sid` - The Account SID associated with the field type
- `assistant_sid` - The SID of the assistant to attach the field type to
- `unique_name` - The unique name of the field type
- `friendly_name` - The friendly name of the field type
- `date_created` - The date in RFC3339 format that the field type was created
- `date_updated` - The date in RFC3339 format that the field type was updated
- `url` - The url of the field type resource
