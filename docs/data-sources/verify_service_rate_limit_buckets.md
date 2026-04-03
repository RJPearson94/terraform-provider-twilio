---
page_title: "twilio_verify_service_rate_limit_buckets Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service_rate_limit_buckets Data Source

Use this data source to access information about existing Verify service rate limit buckets. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limit-buckets) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_service_rate_limit_buckets" "rate_limit_buckets" {
  service_sid    = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  rate_limit_sid = "RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "rate_limit_buckets" {
  value = data.twilio_verify_service_rate_limit_buckets.rate_limit_buckets
}
```

## Schema

### Required

- `rate_limit_sid` (String) The SID of the rate limit that the buckets belong to
- `service_sid` (String) The SID of the Verify service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these rate limit buckets
- `id` (String) The ID of this resource.
- `rate_limit_buckets` (List of Object) A list of rate limit buckets for the service rate limit (see [below for nested schema](#nestedatt--rate_limit_buckets))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--rate_limit_buckets"></a>
### Nested Schema for `rate_limit_buckets`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `interval` (Number)
- `max` (Number)
- `sid` (String)
- `url` (String)
