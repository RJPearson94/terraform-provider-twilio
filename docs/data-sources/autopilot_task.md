---
page_title: "Twilio Autopilot Task"
subcategory: "Autopilot"
---

# twilio_autopilot_task Data Source

Use this data source to access information about an existing Autopilot task. See the [API docs](https://www.twilio.com/docs/autopilot/api/task) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

### SID

```hcl
data "twilio_autopilot_task" "task" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task" {
  value = data.twilio_autopilot_task.task
}
```

### Unique Name

```hcl
data "twilio_autopilot_task" "task" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  unique_name   = "UniqueName"
}

output "task" {
  value = data.twilio_autopilot_task.task
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the task is associated with
- `sid` - (Optional) The SID of the task
- `unique_name` - (Optional) The unique name of the task

~> Either `sid` or `unique_name` must be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the task (Same as the `sid`)
- `sid` - The SID of the task (Same as the `id`)
- `account_sid` - The account SID associated with the task
- `assistant_sid` - The SID of the assistant to attach the task to
- `unique_name` - The unique name of the task
- `friendly_name` - The friendly name of the task
- `actions_url` - The URL to retrieve the actions JSON
- `actions` - JSON string of an Autopilot task
- `date_created` - The date in RFC3339 format that the task was created
- `date_updated` - The date in RFC3339 format that the task was updated
- `url` - The URL of the task resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the task
