---
page_title: "Twilio Voice Queues"
subcategory: "Voice"
---

# twilio_voice_queues Data Source

Use this data source to access information about the queues associated with an existing account. See the [API docs](https://www.twilio.com/docs/voice/api/queue-resource) for more information

## Example Usage

```hcl
data "twilio_account_queues" "queues" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "queues" {
  value = data.twilio_account_queues.queues
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the queues are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `account_sid`)
- `account_sid` - The SID of the account the queues are associated with (Same as the `id`)
- `queues` - A list of `queue` blocks as documented below

---

An `queue` block supports the following:

- `sid` - The SID of the queue
- `friendly_name` - The friendly name of the queue
- `max_size` - The maximum number of calls which can be on the queue
- `average_wait_time` - The average wait time of calls on the queue
- `current_size` - The current size of the queue
- `date_created` - The date in RFC3339 format that the queue was created
- `date_updated` - The date in RFC3339 format that the queue was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving queues
