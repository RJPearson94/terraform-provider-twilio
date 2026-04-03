---
page_title: "twilio_verify_service_rate_limits Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service_rate_limits Data Source

Use this data source to access information about existing Verify service rate limits. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limits) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_service_rate_limits" "rate_limits" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "rate_limits" {
  value = data.twilio_verify_service_rate_limits.rate_limits
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Verify service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these service rate limits
- `id` (String) The ID of this resource.
- `rate_limits` (List of Object) A list of rate limits for the Verify service (see [below for nested schema](#nestedatt--rate_limits))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--rate_limits"></a>
### Nested Schema for `rate_limits`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `description` (String)
- `sid` (String)
- `unique_name` (String)
- `url` (String)
