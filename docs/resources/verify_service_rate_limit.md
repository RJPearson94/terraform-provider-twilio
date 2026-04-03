---
page_title: "twilio_verify_service_rate_limit Resource - twilio"
subcategory: "Verify"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the Verify service. Changing this forces a new resource
- `unique_name` (String) A unique name for the rate limit. Changing this forces a new resource

### Optional

- `description` (String) A description of the rate limit
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this service rate limit
- `date_created` (String) The date and time the service rate limit was created, in RFC 3339 format
- `date_updated` (String) The date and time the service rate limit was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this service rate limit by Twilio
- `url` (String) The absolute URL of the service rate limit resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A rate limit can be imported using the `/Services/{serviceSid}/RateLimits/{sid}` format, e.g.

```shell
terraform import twilio_verify_rate_limit.rate_limit /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/RateLimits/RKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
