---
page_title: "Twilio Autopilot Field Types"
subcategory: "Autopilot"
---

# twilio_autopilot_field_types Data Source

Use this data source to access information about the field types associated with an existing Autopilot assistant. See the [API docs](https://www.twilio.com/docs/autopilot/api/field-type) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_field_types" "field_types" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "field_types" {
  value = data.twilio_autopilot_field_types.field_types
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the field types are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the assistant SID)
- `account_sid` - The SID of the account the field types are associated with
- `assistant_sid` - The SID of the assistant the field types are associated with
- `field_types` - A list of `field_type` blocks as documented below

---

A `field_type` block supports the following:

- `sid` - The SID of the field type
- `unique_name` - The unique name of the field type
- `friendly_name` - The friendly name of the field type
- `date_created` - The date in RFC3339 format that the field type was created
- `date_updated` - The date in RFC3339 format that the field type was updated
- `url` - The url of the field type resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving field types
