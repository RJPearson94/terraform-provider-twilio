---
page_title: "Twilio TaskRouter Task Queues"
subcategory: "TaskRouter"
---

# twilio_taskrouter_task_queues Data Source

Use this data source to access information about the task queues associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/task-queue) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_task_queues" "task_queues" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_queues" {
  value = data.twilio_taskrouter_task_queues.task_queues
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the task queues are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `workspace_sid`)
- `workspace_sid` - The SID of the workspace the task queues are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the task queues
- `task_queues` - A list of `task_queue` blocks as documented below

---

A `task_queue` block supports the following:

- `sid` - The SID of the task queue
- `friendly_name` - The name of the task queue
- `task_order` - How TaskRouter will assign workers tasks on the queue
- `assignment_activity_name` - The assignment activity name for the task queue
- `assignment_activity_sid` - The assignment activity SID for the task queue
- `reservation_activity_name` - The reservation activity name for the task queue
- `reservation_activity_sid` - The reservation activity SID for the task queue
- `target_workers` - Worker selection criteria for any tasks that enter the task queue
- `max_reserved_workers` - The max number of workers to create a reservation for
- `date_created` - The date in RFC3339 format that the task queue was created
- `date_updated` - The date in RFC3339 format that the task queue was updated
- `url` - The URL of the task queue

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving task queues
