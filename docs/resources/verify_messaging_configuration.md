---
page_title: "twilio_verify_messaging_configuration Resource - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_messaging_configuration Resource

Manages a Verify messaging configuration

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "Test Messaging Service"
}

resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = twilio_verify_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
  country_code          = "GB"
}
```

## Schema

### Required

- `country_code` (String) The ISO-3166-1 country code of the messaging configuration. Changing this forces a new resource
- `messaging_service_sid` (String) The SID of the messaging service to associate with this configuration
- `service_sid` (String) The SID of the Verify service. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this messaging configuration
- `date_created` (String) The date and time the messaging configuration was created, in RFC 3339 format
- `date_updated` (String) The date and time the messaging configuration was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the messaging configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A messaging configuration can be imported using the `/Services/{serviceSid}/MessagingConfigurations/{countryCode}` format, e.g.

```shell
terraform import twilio_verify_rate_limit.rate_limit /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/GB
```
