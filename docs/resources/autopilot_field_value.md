---
page_title: "Twilio Autopilot Field Value"
subcategory: "Autopilot"
---

# twilio_autopilot_field_value Resource

Manages a Autopilot Field Value

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = twilio_autopilot_field_type.field_type.assistant_sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "en-US"
  value          = "test"
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the field value to. Changing this forces a new resource to be created
- `field_type_sid` - (Mandatory) The SID of the field type to attach the field value to. Changing this forces a new resource to be created
- `language` - (Mandatory) The field value language. Changing this forces a new resource to be created
- `value` - (Mandatory) The field value. Changing this forces a new resource to be created
- `synonym_of` - (Optional) The word which this field value is a synonym of. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field value (Same as the SID)
- `sid` - The SID of the field value (Same as the ID)
- `account_sid` - The Account SID associated with the field value
- `assistant_sid` - The SID of the assistant to attach the field value to
- `field_type_sid` - The SID of the field type to attach the field value to
- `language` - The field value language
- `value` - The field value
- `synonym_of` - The word which this field value is a synonym of
- `date_created` - The date in RFC3339 format that the field value was created
- `date_updated` - The date in RFC3339 format that the field value was updated
- `url` - The url of the field value resource
