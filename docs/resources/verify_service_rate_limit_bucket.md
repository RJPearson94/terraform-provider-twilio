---
page_title: "Twilio Verify Service Rate Limit Bucket"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the rate limit bucket with. Changing this forces a new resource to be created
- `rate_limit_sid` - (Mandatory) The rate limit SID to associate the rate limit bucket with. Changing this forces a new resource to be created
- `max` - (Mandatory) The maximum number of requests that can occur during the interval
- `interval` - (Mandatory) The duration (in seconds) which the rate limit will be monitored/ enforced

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limit bucket (Same as the `sid`)
- `sid` - The SID of the rate limit bucket (Same as the `id`)
- `account_sid` - The account SID the rate limit bucket is associated with
- `service_sid` - The service SID the rate limit bucket is associated with
- `rate_limit_sid` - The rate limit SID the rate limit bucket is associated with
- `max` - The maximum number of requests that can occur during the interval
- `interval` - The duration (in seconds) which the rate limit will be monitored/ enforced
- `date_created` - The date in RFC3339 format that the rate limit bucket was created
- `date_updated` - The date in RFC3339 format that the rate limit bucket was updated
- `url` - The URL of the rate limit bucket

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the rate limit bucket
- `update` - (Defaults to 10 minutes) Used when updating the rate limit bucket
- `read` - (Defaults to 5 minutes) Used when retrieving the rate limit bucket
- `delete` - (Defaults to 10 minutes) Used when deleting the rate limit bucket

## Import

A rate limit can be imported using the `/Services/{serviceSid}/RateLimits/{RateLimitsSid}/Buckets/{sid}` format, e.g.

```shell
terraform import twilio_verify_service_rate_limit_bucket.rate_limit_bucket /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Buckets/BLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
