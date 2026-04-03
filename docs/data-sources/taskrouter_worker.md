---
page_title: "twilio_taskrouter_worker Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
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

## Schema

### Required

- `sid` (String) The SID of the worker to retrieve
- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this worker
- `activity_name` (String) The friendly name of the worker's current activity
- `activity_sid` (String) The SID of the worker's current activity
- `attributes` (String) A JSON string of attributes for the worker
- `available` (Boolean) Whether the worker is available to receive tasks
- `date_created` (String) The date and time the worker was created, in RFC 3339 format
- `date_status_changed` (String) The date and time the worker's activity status last changed, in RFC 3339 format
- `date_updated` (String) The date and time the worker was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the worker
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the worker resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
