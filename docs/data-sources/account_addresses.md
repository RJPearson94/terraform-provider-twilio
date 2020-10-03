---
page_title: "Twilio Account Addresses"
subcategory: "Account"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the addresses are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the account SID)
- `account_sid` - The SID of the account the addresses are associated with
- `addresses` - A list of `address` blocks as documented below

---

An `address` block supports the following:

- `sid` - The SID of the address
- `friendly_name` - The friendly name of the address
- `customer_name` - The customer/ business name
- `street` - The address street
- `street_secondary` - The address secondary street
- `city` - The address city
- `region` - The address region
- `postal_code` - The address postal code
- `iso_country` - The address ISO country
- `emergency_enabled` - Whether emergency calling is enabled for the address
- `validated` - Whether the address has been validated
- `verified` - Whether the address has been verified
- `date_created` - The date in RFC3339 format that the address was created
- `date_updated` - The date in RFC3339 format that the address was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving addresses
