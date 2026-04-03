---
page_title: "twilio_verify_messaging_configurations Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_messaging_configurations Data Source

Use this data source to access information about existing Verify messaging configurations

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_messaging_configurations" "messaging_configurations" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "messaging_configurations" {
  value = data.twilio_verify_messaging_configurations.messaging_configurations
}
```


## Schema

### Required

- `service_sid` (String) The SID of the Verify service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these messaging configurations
- `id` (String) The ID of this resource.
- `messaging_configurations` (List of Object) A list of messaging configurations for the Verify service (see [below for nested schema](#nestedatt--messaging_configurations))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--messaging_configurations"></a>
### Nested Schema for `messaging_configurations`

Read-Only:

- `country_code` (String)
- `date_created` (String)
- `date_updated` (String)
- `messaging_service_sid` (String)
- `url` (String)
