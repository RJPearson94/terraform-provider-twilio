---
page_title: "Twilio Voice Queue"
subcategory: "Voice"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the queue with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The friendly name of the queue. The length of the string must be between `1` and `64` characters (inclusive)
- `max_size` - (Optional) The maximum number of calls which can be on the queue. The value must be between `1` and `5000` (inclusive). The default value is `100`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the queue (Same as the `sid`)
- `sid` - The SID of the queue (Same as the `id`)
- `account_sid` - The account SID the queue is associated with
- `friendly_name` - The friendly name of the queue
- `max_size` - The maximum number of calls which can be on the queue
- `average_wait_time` - The average wait time of calls on the queue
- `current_size` - The current size of the queue
- `date_created` - The date in RFC3339 format that the queue was created
- `date_updated` - The date in RFC3339 format that the queue was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the queue
- `update` - (Defaults to 10 minutes) Used when updating the queue
- `read` - (Defaults to 5 minutes) Used when retrieving the queue
- `delete` - (Defaults to 10 minutes) Used when deleting the queue

## Import

An queue can be imported using the `/Accounts/{queueSid}/Queues/{sid}` format, e.g.

```shell
terraform import twilio_voice_queue.queue /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Queues/QUXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
