---
page_title: "Twilio Autopilot Field Type"
subcategory: "Autopilot"
---

# twilio_autopilot_field_type Resource

Manages a Autopilot field type. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-type) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

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
- `account_sid` - The account SID associated with the field type
- `assistant_sid` - The SID of the assistant to attach the field type to
- `unique_name` - The unique name of the field type
- `friendly_name` - The friendly name of the field type
- `date_created` - The date in RFC3339 format that the field type was created
- `date_updated` - The date in RFC3339 format that the field type was updated
- `url` - The url of the field type resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the field type
- `update` - (Defaults to 10 minutes) Used when updating the field type
- `read` - (Defaults to 5 minutes) Used when retrieving the field type
- `delete` - (Defaults to 10 minutes) Used when deleting the field type

## Import

A field type can be imported using the `/Assistants/{assistantSid}/FieldTypes/{sid}` format, e.g.

```shell
terraform import twilio_autopilot_field_type.field_type /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/FieldTypes/UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
