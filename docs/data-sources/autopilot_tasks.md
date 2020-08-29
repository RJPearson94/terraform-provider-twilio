---
page_title: "Twilio Autopilot Tasks"
subcategory: "Autopilot"
---

# twilio_autopilot_tasks Data Source

Use this data source to access information about the tasks associated with an existing Autopilot assistant. See the [API docs](https://www.twilio.com/docs/autopilot/api/task) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_tasks" "tasks" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "tasks" {
  value = data.twilio_autopilot_tasks.tasks
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the tasks are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the assistant SID)
- `account_sid` - The SID of the account the tasks are associated with
- `assistant_sid` - The SID of the assistant the tasks are associated with
- `tasks` - A list of `task` blocks as documented below

---

A `task` block supports the following:

- `sid` - The SID of the task
- `unique_name` - The unique name of the task
- `friendly_name` - The friendly name of the task
- `actions_url` - The url to retrieve the actions json
- `actions` - JSON string of an Autopilot task
- `date_created` - The date in RFC3339 format that the task was created
- `date_updated` - The date in RFC3339 format that the task was updated
- `url` - The url of the task resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving tasks
