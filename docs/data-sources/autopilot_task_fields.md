---
page_title: "Twilio Autopilot Task Fields"
subcategory: "Autopilot"
---

# twilio_autopilot_task_fields Data Source

Use this data source to access information about the task field associated with an existing Autopilot assistant and task. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-field) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_task_fields" "task_fields" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  task_sid      = "UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_fields" {
  value = data.twilio_autopilot_task_fields.task_fields
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the fields are associated with
- `task_sid` - (Mandatory) The SID of the task the fields are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `assistant_sid/task_sid`
- `account_sid` - The SID of the account the fields are associated with
- `assistant_sid` - The SID of the assistant the fields are associated with
- `task_sid` - The SID of the task the fields are associated with
- `fields` - A list of `field` blocks as documented below

---

A `field` block supports the following:

- `sid` - The SID of the field
- `unique_name` - The unique name of the field
- `field_type` - The type of field
- `date_created` - The date in RFC3339 format that the field was created
- `date_updated` - The date in RFC3339 format that the field was updated
- `url` - The url of the field

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving fields
