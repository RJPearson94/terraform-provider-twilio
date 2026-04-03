---
page_title: "twilio_voice_queues Data Source - twilio"
subcategory: "Voice"
description: |-
  
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

## Schema

### Required

- `account_sid` (String) The SID of the account to list queues for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `id` (String) The ID of this resource.
- `queues` (List of Object) The list of voice queues in the account (see [below for nested schema](#nestedatt--queues))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--queues"></a>
### Nested Schema for `queues`

Read-Only:

- `average_wait_time` (Number)
- `current_size` (Number)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `max_size` (Number)
- `sid` (String)
