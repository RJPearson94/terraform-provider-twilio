---
page_title: "Twilio Autopilot Task Field"
subcategory: "Autopilot"
---

# twilio_autopilot_task_field Data Source

Use this data source to access information about an existing Autopilot task field. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-field) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_task_field" "task_field" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  task_sid      = "UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_field" {
  value = data.twilio_autopilot_task_field.task_field
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the field is associated with
- `task_sid` - (Mandatory) The SID of the task the field is associated with
- `sid` - (Mandatory) The SID of the field

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the field (Same as the SID)
- `sid` - The SID of the field (Same as the ID)
- `account_sid` - The account SID associated with the field
- `assistant_sid` - The SID of the assistant to attach the task to
- `task_sid` - The SID of the task to attach the field to
- `unique_name` - The unique name of the field
- `field_type` - The type of field
- `date_created` - The date in RFC3339 format that the field was created
- `date_updated` - The date in RFC3339 format that the field was updated
- `url` - The URL of the field

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the field
