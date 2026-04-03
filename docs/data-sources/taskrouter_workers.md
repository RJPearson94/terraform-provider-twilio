---
page_title: "twilio_taskrouter_workers Data Source - twilio"
subcategory: "TaskRouter"
description: |-
  
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

## Schema

### Required

- `workspace_sid` (String) The SID of the TaskRouter workspace

### Optional

- `activity_name` (String) Filter workers by activity name
- `activity_sid` (String) Filter workers by activity SID
- `available` (Boolean) Filter workers by availability
- `friendly_name` (String) Filter workers by friendly name
- `target_workers_expression` (String) Filter workers by a target workers expression
- `task_queue_name` (String) Filter workers by task queue name
- `task_queue_sid` (String) Filter workers by task queue SID
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the workers
- `id` (String) The ID of this resource.
- `workers` (List of Object) A list of workers in the workspace (see [below for nested schema](#nestedatt--workers))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--workers"></a>
### Nested Schema for `workers`

Read-Only:

- `activity_name` (String)
- `activity_sid` (String)
- `attributes` (String)
- `available` (Boolean)
- `date_created` (String)
- `date_status_changed` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
- `url` (String)
