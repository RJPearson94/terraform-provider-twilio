---
page_title: "twilio_account_address Data Source - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_address Data Source

Use this data source to access information about an existing address. See the [API docs](https://www.twilio.com/docs/usage/api/address) for more information

## Example Usage

```hcl
data "twilio_account_address" "address" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "customer_name" {
  value = data.twilio_account_address.address.customer_name
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this address
- `sid` (String) The SID of the address to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `city` (String) The city of the address
- `customer_name` (String) The name of the customer associated with the address
- `date_created` (String) The date and time the address was created, in RFC 3339 format
- `date_updated` (String) The date and time the address was last updated, in RFC 3339 format
- `emergency_enabled` (Boolean) Whether emergency calling is enabled for the address
- `friendly_name` (String) A human-readable label for the address
- `id` (String) The ID of this resource.
- `iso_country` (String) The ISO 3166-1 alpha-2 country code of the address
- `postal_code` (String) The postal code of the address
- `region` (String) The state or region of the address
- `street` (String) The street address
- `street_secondary` (String) The secondary street address information (e.g., suite or apartment number)
- `validated` (Boolean) Whether the address has been validated by Twilio
- `verified` (Boolean) Whether the address has been verified by the customer

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
