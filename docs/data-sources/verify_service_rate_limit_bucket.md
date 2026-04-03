---
page_title: "twilio_verify_service_rate_limit_bucket Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service_rate_limit_bucket Data Source

Use this data source to access information about an existing Verify service rate limit bucket. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limit-buckets) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  rate_limit_sid = "RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid            = "YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "rate_limit_bucket" {
  value = data.twilio_verify_service_rate_limit_bucket.rate_limit_bucket
}
```

## Schema

### Required

- `rate_limit_sid` (String) The SID of the rate limit that this bucket belongs to
- `service_sid` (String) The SID of the Verify service
- `sid` (String) The SID of the rate limit bucket to fetch

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this rate limit bucket
- `date_created` (String) The date and time the rate limit bucket was created, in RFC 3339 format
- `date_updated` (String) The date and time the rate limit bucket was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `interval` (Number) The time interval in seconds for the rate limit bucket
- `max` (Number) The maximum number of requests permitted in the given time interval
- `url` (String) The absolute URL of the rate limit bucket resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
