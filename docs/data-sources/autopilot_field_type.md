---
page_title: "Twilio Autopilot Field Type"
subcategory: "Autopilot"
---

# twilio_autopilot_field_type Data Source

Use this data source to access information about an existing Autopilot field type. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-type) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

### SID

```hcl
data "twilio_autopilot_field_type" "field_type" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "UBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "field_type" {
  value = data.twilio_autopilot_field_type.field_type
}
```

### Unique Name

```hcl
data "twilio_autopilot_field_type" "field_type" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  unique_name   = "UniqueName"
}

output "field_type" {
  value = data.twilio_autopilot_field_type.field_type
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the field type is associated with
- `sid` - (Optional) The SID of the field type
- `unique_name` - (Optional) The unique name of the field type

~> Either `sid` or `unique_name` must be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field type (Same as the `sid`)
- `sid` - The SID of the field type (Same as the `id`)
- `account_sid` - The account SID associated with the field type
- `assistant_sid` - The SID of the assistant to attach the field type to
- `unique_name` - The unique name of the field type
- `friendly_name` - The friendly name of the field type
- `date_created` - The date in RFC3339 format that the field type was created
- `date_updated` - The date in RFC3339 format that the field type was updated
- `url` - The URL of the field type resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the field type
