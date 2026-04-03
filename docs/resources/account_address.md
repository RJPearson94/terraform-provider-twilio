---
page_title: "twilio_account_address Resource - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_address Resource

Manages a Twilio address. See the [API docs](https://www.twilio.com/docs/usage/api/address) for more information

## Example Usage

```hcl
resource "twilio_account_address" "address" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  customer_name = "Test User"
  street        = "123 Fake Street"
  city          = "Fake City"
  region        = "Fake Region"
  postal_code   = "AB12DC"
  iso_country   = "GB"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this address. Changing this forces a new resource
- `city` (String) The city of the address
- `customer_name` (String) The name of the customer associated with the address
- `iso_country` (String) The ISO 3166-1 alpha-2 country code of the address. Changing this forces a new resource
- `postal_code` (String) The postal code of the address
- `region` (String) The state or region of the address
- `street` (String) The street address

### Optional

- `emergency_enabled` (Boolean) Whether emergency calling is enabled for the address. Defaults to `false`
- `friendly_name` (String) A human-readable label for the address
- `street_secondary` (String) The secondary street address information (e.g., suite or apartment number)
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the address was created, in RFC 3339 format
- `date_updated` (String) The date and time the address was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this address by Twilio
- `validated` (Boolean) Whether the address has been validated by Twilio
- `verified` (Boolean) Whether the address has been verified by the customer

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An address can be imported using the `/Accounts/{addressSid}/Addresses/{sid}` format, e.g.

```shell
terraform import twilio_account_address.address /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
