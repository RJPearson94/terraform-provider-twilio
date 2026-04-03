---
page_title: "twilio_voice_queue Resource - twilio"
subcategory: "Voice"
description: |-
  
---

# twilio_voice_queue Resource

Manages a Twilio queue. See the [API docs](https://www.twilio.com/docs/voice/api/queue-resource) for more information

## Example Usage

```hcl
resource "twilio_voice_queue" "queue" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test Queue"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account to create the queue in. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the queue (1–64 characters)

### Optional

- `max_size` (Number) The maximum number of calls that can be in the queue at one time (1–5000). Defaults to `100`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `average_wait_time` (Number) The average wait time of calls currently in the queue, in seconds
- `current_size` (Number) The current number of calls in the queue
- `date_created` (String) The date and time the queue was created, in RFC 3339 format
- `date_updated` (String) The date and time the queue was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this queue by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An queue can be imported using the `/Accounts/{queueSid}/Queues/{sid}` format, e.g.

```shell
terraform import twilio_voice_queue.queue /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
