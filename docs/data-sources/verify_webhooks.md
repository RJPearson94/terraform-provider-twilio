---
page_title: "twilio_verify_webhooks Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_webhooks Data Source

Use this data source to access information about existing Verify webhooks. See the [API docs](https://www.twilio.com/docs/verify/api/webhooks) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_webhooks" "webhooks" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhooks" {
  value = data.twilio_verify_webhooks.webhooks
}
```


## Schema

### Required

- `service_sid` (String) The SID of the Verify service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these webhooks
- `id` (String) The ID of this resource.
- `webhooks` (List of Object) A list of webhooks for the Verify service (see [below for nested schema](#nestedatt--webhooks))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--webhooks"></a>
### Nested Schema for `webhooks`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `event_types` (List of String)
- `friendly_name` (String)
- `sid` (String)
- `status` (String)
- `url` (String)
- `version` (String)
- `webhook_method` (String)
- `webhook_url` (String)
