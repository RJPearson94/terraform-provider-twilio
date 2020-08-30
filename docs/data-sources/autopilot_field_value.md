---
page_title: "Twilio Autopilot Field Value"
subcategory: "Autopilot"
---

# twilio_autopilot_field_value Data Source

Use this data source to access information about an existing Autopilot field value. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-value) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_field_value" "field_value" {
  assistant_sid  = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  field_type_sid = "UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid            = "UCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "field_value" {
  value = data.twilio_autopilot_field_value.field_value
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the field value is associated with
- `field_type_sid` - (Mandatory) The SID of the field type the field value is associated with
- `sid` - (Mandatory) The SID of the field value

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field value (Same as the SID)
- `sid` - The SID of the field value (Same as the ID)
- `account_sid` - The account SID associated with the field value
- `assistant_sid` - The SID of the assistant to attach the field value to
- `field_type_sid` - The SID of the field type to attach the field value to
- `language` - The field value language
- `value` - The field value
- `synonym_of` - The word which this field value is a synonym of
- `date_created` - The date in RFC3339 format that the field value was created
- `date_updated` - The date in RFC3339 format that the field value was updated
- `url` - The url of the field value resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the field value
