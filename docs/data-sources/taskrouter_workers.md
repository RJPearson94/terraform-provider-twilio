---
page_title: "Twilio TaskRouter Workers"
subcategory: "TaskRouter"
---

# twilio_taskrouter_workers Data Source

Use this data source to access information about the workers associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/worker) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_workers" "workers" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "workers" {
  value = data.twilio_taskrouter_workers.workers
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the workers are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the workspace SID)
- `workspace_sid` - The SID of the workspace the workers are associated with
- `account_sid` - The account SID associated with the workers
- `workers` - A list of `worker` blocks as documented below

---

A `worker` block supports the following:

- `sid` - The SID of the worker
- `friendly_name` - The name of the worker
- `attributes` - JSON string of worker attributes
- `activity_sid` - Activity SID to be assigned to the worker
- `activity_name` - Friendly name of activity
- `available` - Is the worker available to receive tasks
- `date_created` - The date in RFC3339 format that the worker was created
- `date_updated` - The date in RFC3339 format that the worker was updated
- `date_status_changed` - The date in RFC3339 format that the worker status was changed
- `url` - The URL of the worker

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving workers
