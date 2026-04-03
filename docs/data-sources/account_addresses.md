---
page_title: "twilio_account_addresses Data Source - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_addresses Data Source

Use this data source to access information about the addresses associated with an existing account. See the [API docs](https://www.twilio.com/docs/usage/api/address) for more information

## Example Usage

```hcl
data "twilio_account_addresses" "addresses" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "addresses" {
  value = data.twilio_account_addresses.addresses
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account to list addresses for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `addresses` (List of Object) A list of addresses associated with the account (see [below for nested schema](#nestedatt--addresses))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--addresses"></a>
### Nested Schema for `addresses`

Read-Only:

- `city` (String)
- `customer_name` (String)
- `date_created` (String)
- `date_updated` (String)
- `emergency_enabled` (Boolean)
- `friendly_name` (String)
- `iso_country` (String)
- `postal_code` (String)
- `region` (String)
- `sid` (String)
- `street` (String)
- `street_secondary` (String)
- `validated` (Boolean)
- `verified` (Boolean)
