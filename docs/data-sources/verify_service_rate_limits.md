---
page_title: "Twilio Verify Service Rate Limits"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID the service rate limits are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limits (Same as the `service_sid`)
- `account_sid` - The account SID the rate limits are associated with
- `service_sid` - The service SID the rate limits are associated with
- `rate_limits` - A list of `rate_limit` blocks as documented below

---

A `rate_limit` block supports the following:

- `sid` - The SID of the rate limit
- `unique_name` - Unique name of the rate limit
- `description` - The description of the rate limit
- `date_created` - The date in RFC3339 format that the rate limit was created
- `date_updated` - The date in RFC3339 format that the rate limit was updated
- `url` - The URL of the rate limit

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the service rate limits
