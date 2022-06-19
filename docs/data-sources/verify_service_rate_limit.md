---
page_title: "Twilio Verify Service Rate Limit"
subcategory: "Verify"
---

# twilio_verify_service_rate_limit Data Source

Use this data source to access information about an existing Verify service rate limit. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limits) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_service_rate_limit" "rate_limit" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "rate_limit" {
  value = data.twilio_verify_service_rate_limit.rate_limit
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service rate limit
- `service_sid` - (Mandatory) The service SID the rate limit is associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limit (Same as the `sid`)
- `sid` - The SID of the rate limit (Same as the `id`)
- `account_sid` - The account SID the rate limit is associated with
- `service_sid` - The service SID the rate limit is associated with
- `unique_name` - Unique name of the rate limit
- `description` - The description of the rate limit
- `date_created` - The date in RFC3339 format that the rate limit was created
- `date_updated` - The date in RFC3339 format that the rate limit was updated
- `url` - The URL of the rate limit

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the service rate limit
