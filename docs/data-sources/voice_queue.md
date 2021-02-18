---
page_title: "Twilio Voice Queue"
subcategory: "Voice"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the queue is associated with
- `sid` - (Mandatory) The SID of the queue

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

- `read` - (Defaults to 5 minutes) Used when retrieving the queue details
