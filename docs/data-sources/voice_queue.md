---
page_title: "twilio_voice_queue Data Source - twilio"
subcategory: "Voice"
description: |-
  
---

# twilio_voice_queue Data Source

Use this data source to access information about an existing queue. See the [API docs](https://www.twilio.com/docs/voice/api/queue-resource) for more information

## Example Usage

```hcl
data "twilio_voice_queue" "queue" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "current_size" {
  value = data.twilio_voice_queue.queue.current_size
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns the queue
- `sid` (String) The SID of the voice queue to read

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `average_wait_time` (Number) The average wait time of calls currently in the queue, in seconds
- `current_size` (Number) The current number of calls in the queue
- `date_created` (String) The date and time the queue was created, in RFC 3339 format
- `date_updated` (String) The date and time the queue was last updated, in RFC 3339 format
- `friendly_name` (String) The human-readable label for the queue
- `id` (String) The ID of this resource.
- `max_size` (Number) The maximum number of calls that can be in the queue at one time

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
