---
page_title: "Twilio Verify Service Rate Limit Bucket"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service rate limit bucket
- `service_sid` - (Mandatory) The service SID the rate limit bucket is associated with
- `rate_limit_sid` - (Mandatory) The rate limit SID the rate limit bucket is associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limit bucket (Same as the `sid`)
- `sid` - The SID of the rate limit bucket (Same as the `id`)
- `account_sid` - The account SID of the rate limit bucket is associated with
- `service_sid` - The service SID the rate limit bucket is associated with
- `rate_limit_sid` - The rate limit SID the rate limit bucket is associated with
- `max` - The maximum number of requests that can occur during the interval
- `interval` - The duration (in seconds) which the rate limit will be monitored/ enforced
- `date_created` - The date in RFC3339 format that the rate limit bucket was created
- `date_updated` - The date in RFC3339 format that the rate limit bucket was updated
- `url` - The URL of the rate limit bucket

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the service rate limit bucket
