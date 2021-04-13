---
page_title: "Twilio Autopilot Field Value"
subcategory: "Autopilot"
---

# twilio_autopilot_field_value Resource

Manages an Autopilot field value. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-value) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

### Basic

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

### With Synonym

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "test"
}

resource "twilio_autopilot_field_type" "field_type" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "test"
}

resource "twilio_autopilot_field_value" "field_value_hello" {
  assistant_sid  = twilio_autopilot_field_type.field_type.assistant_sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "en-US"
  value          = "hello"
}

resource "twilio_autopilot_field_value" "field_value_hi" {
  assistant_sid  = twilio_autopilot_field_type.field_type.assistant_sid
  field_type_sid = twilio_autopilot_field_type.field_type.sid
  language       = "en-US"
  value          = "hi"
  synonym_of     = twilio_autopilot_field_value.field_value_hello.value
}
```

!> The `synonym_of` argument can be set to the `sid` or the `value` of the field value. The API currently returns the synonym value so if the `sid` is supplied Terraform will show a difference

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to associate the field value with. Changing this forces a new resource to be created
- `field_type_sid` - (Mandatory) The SID of the field type to associate the field value with. Changing this forces a new resource to be created
- `language` - (Mandatory) The field value language. Changing this forces a new resource to be created
- `value` - (Mandatory) The field value. Changing this forces a new resource to be created
- `synonym_of` - (Optional) The word which this field value is a synonym of. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field value (Same as the `sid`)
- `sid` - The SID of the field value (Same as the `id`)
- `account_sid` - The account SID associated with the field value
- `assistant_sid` - The SID of the assistant to attach the field value to
- `field_type_sid` - The SID of the field type to attach the field value to
- `language` - The field value language
- `value` - The field value
- `synonym_of` - The word which this field value is a synonym of
- `date_created` - The date in RFC3339 format that the field value was created
- `date_updated` - The date in RFC3339 format that the field value was updated
- `url` - The URL of the field value resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the field value
- `read` - (Defaults to 5 minutes) Used when retrieving the field value
- `delete` - (Defaults to 10 minutes) Used when deleting the field value

## Import

A field value can be imported using the `/Assistants/{assistantSid}/FieldTypes/{fieldTypeSid}/FieldValues/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_field_value.field_value /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldValues/UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
