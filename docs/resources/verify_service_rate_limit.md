---
page_title: "Twilio Verify Service Rate Limit"
subcategory: "Verify"
---

# twilio_verify_service_rate_limit Resource

Manages a Verify service rate limit. See the [API docs](https://www.twilio.com/docs/verify/api/service-rate-limits) for more information

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
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID to associate the rate limit bucket with. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) Unique name of the rate limit. Changing this forces a new resource to be created
- `description` - (Optional) The description of the rate limit

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the rate limit (Same as the `sid`)
- `sid` - The SID of the rate limit (Same as the `id`)
- `account_sid` - The account SID the rate limit is associated with
- `service_sid` - The service SID the rate limit bucket is associated with
- `unique_name` - Unique name of the rate limit
- `description` - The description of the rate limit
- `date_created` - The date in RFC3339 format that the rate limit was created
- `date_updated` - The date in RFC3339 format that the rate limit was updated
- `url` - The URL of the rate limit

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the rate limit
- `update` - (Defaults to 10 minutes) Used when updating the rate limit
- `read` - (Defaults to 5 minutes) Used when retrieving the rate limit
- `delete` - (Defaults to 10 minutes) Used when deleting the rate limit

## Import

A rate limit can be imported using the `/Services/{serviceSid}/RateLimits/{sid}` format, e.g.

```shell
terraform import twilio_verify_rate_limit.rate_limit /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
