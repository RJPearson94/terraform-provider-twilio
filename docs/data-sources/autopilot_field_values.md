---
page_title: "Twilio Autopilot Field Values"
subcategory: "Autopilot"
---

# twilio_autopilot_field_values Data Source

Use this data source to access information about the field values associated with an existing Autopilot assistant and field type. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-value) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_field_values" "field_values" {
  assistant_sid  = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  field_type_sid = "UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "field_values" {
  value = data.twilio_autopilot_field_values.field_values
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the field values are associated with
- `field_type_sid` - (Mandatory) The SID of the field type the field values are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `assistant_sid/field_type_sid`
- `account_sid` - The SID of the account the field values are associated with
- `assistant_sid` - The SID of the assistant the field values are associated with
- `field_type_sid` - The SID of the field type the field values are associated with
- `field_values` - A list of `field_value` blocks as documented below

---

A `field_value` block supports the following:

- `sid` - The SID of the field value
- `language` - The field value language
- `value` - The field value
- `synonym_of` - The word which this field value is a synonym of
- `date_created` - The date in RFC3339 format that the field value was created
- `date_updated` - The date in RFC3339 format that the field value was updated
- `url` - The URL of the field value resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving field values
