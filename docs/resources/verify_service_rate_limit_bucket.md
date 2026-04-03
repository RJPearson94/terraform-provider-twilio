---
page_title: "twilio_verify_service_rate_limit_bucket Resource - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service_rate_limit_bucket Resource

Manages a Verify service rate limit bucket. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limit-buckets) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = twilio_verify_service.service.sid
  unique_name = "Test Service Rate Limit"
}

resource "twilio_verify_service_rate_limit_bucket" "rate_limit_bucket" {
  service_sid    = twilio_verify_service_rate_limit.rate_limit.service_sid
  rate_limit_sid = twilio_verify_service_rate_limit.rate_limit.sid
  max            = 10
  interval       = 2
}
```

## Schema

### Required

- `interval` (Number) The time interval in seconds for the rate limit bucket
- `max` (Number) The maximum number of requests permitted in the given time interval
- `rate_limit_sid` (String) The SID of the rate limit that this bucket belongs to. Changing this forces a new resource
- `service_sid` (String) The SID of the Verify service. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this rate limit bucket
- `date_created` (String) The date and time the rate limit bucket was created, in RFC 3339 format
- `date_updated` (String) The date and time the rate limit bucket was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this rate limit bucket by Twilio
- `url` (String) The absolute URL of the rate limit bucket resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A rate limit can be imported using the `/Services/{serviceSid}/RateLimits/{RateLimitsSid}/Buckets/{sid}` format, e.g.

```shell
terraform import twilio_verify_service_rate_limit_bucket.rate_limit_bucket /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
