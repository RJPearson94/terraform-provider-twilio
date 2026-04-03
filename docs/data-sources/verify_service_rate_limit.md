---
page_title: "twilio_verify_service_rate_limit Data Source - twilio"
subcategory: "Verify"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the Verify service
- `sid` (String) The SID of the service rate limit to fetch

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this service rate limit
- `date_created` (String) The date and time the service rate limit was created, in RFC 3339 format
- `date_updated` (String) The date and time the service rate limit was last updated, in RFC 3339 format
- `description` (String) A description of the rate limit
- `id` (String) The ID of this resource.
- `unique_name` (String) The unique name of the rate limit
- `url` (String) The absolute URL of the service rate limit resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
