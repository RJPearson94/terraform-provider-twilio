---
page_title: "twilio_verify_messaging_configuration Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_messaging_configuration Data Source

Use this data source to access information about an existing Verify messaging configuration

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid  = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  country_code = "GB"
}

output "messaging_configuration" {
  value = data.twilio_verify_messaging_configuration.messaging_configuration
}
```


## Schema

### Required

- `country_code` (String) The ISO-3166-1 country code of the messaging configuration to fetch
- `service_sid` (String) The SID of the Verify service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this messaging configuration
- `date_created` (String) The date and time the messaging configuration was created, in RFC 3339 format
- `date_updated` (String) The date and time the messaging configuration was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `messaging_service_sid` (String) The SID of the messaging service associated with this configuration
- `url` (String) The absolute URL of the messaging configuration resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
