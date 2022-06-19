---
page_title: "Twilio Verify Service Rate Limit Buckets"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID the service rate limit buckets are associated with
- `rate_limit_sid` - (Mandatory) The rate limit SID the service rate limit buckets are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limit buckets
- `account_sid` - The account SID of the rate limit buckets are associated with
- `service_sid` - The service SID the rate limit buckets are associated with
- `rate_limit_sid` - The service rate limit SID the rate limit buckets are associated with
- `rate_limit_buckets` - A list of `rate_limit_bucket` blocks as documented below

---

A `rate_limit_bucket` block supports the following:

- `sid` - The SID of the rate limit bucket
- `max` - The maximum number of requests that can occur during the interval
- `interval` - The duration (in seconds) which the rate limit will be monitored/ enforced
- `date_created` - The date in RFC3339 format that the rate limit bucket was created
- `date_updated` - The date in RFC3339 format that the rate limit bucket was updated
- `url` - The URL of the rate limit bucket

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the service rate limit buckets
