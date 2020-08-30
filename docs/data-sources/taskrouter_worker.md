---
page_title: "Twilio TaskRouter Worker"
subcategory: "TaskRouter"
---

# twilio_taskrouter_worker Data Source

Use this data source to access information about an existing TaskRouter worker. See the [API docs](https://www.twilio.com/docs/taskrouter/api/worker) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_worker" "worker" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "WKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "worker" {
  value = data.twilio_taskrouter_worker.worker
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the worker is associated with
- `sid` - (Mandatory) The SID of the worker

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the worker (Same as the SID)
- `sid` - The SID of the worker (Same as the ID)
- `account_sid` - The account SID of the worker is deployed into
- `workspace_sid` - The workspace SID to create the worker under
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

- `read` - (Defaults to 5 minutes) Used when retrieving the worker
