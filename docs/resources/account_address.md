---
page_title: "Twilio Account Address"
subcategory: "Account"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the address with. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the address. The default value is an empty string/ no configuration specified
- `customer_name` - (Mandatory) The customer/ business name. The value cannot be an empty string
- `street` - (Mandatory) The address street. The value cannot be an empty string
- `street_secondary` - (Optional) The address secondary street. The default value is an empty string/ no configuration specified
- `city` - (Mandatory) The address city. The value cannot be an empty string
- `region` - (Mandatory) The address region. The value cannot be an empty string
- `postal_code` - (Mandatory) The address postal code. The value cannot be an empty string
- `iso_country` - (Mandatory) The address ISO country. The value cannot be an empty string
- `emergency_enabled` - (Optional) Whether emergency calling is enabled for the address. The default value is `false`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the address (Same as the `sid`)
- `sid` - The SID of the address (Same as the `id`)
- `account_sid` - The account SID the address is associated with
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

- `create` - (Defaults to 10 minutes) Used when creating the address
- `update` - (Defaults to 10 minutes) Used when updating the address
- `read` - (Defaults to 5 minutes) Used when retrieving the address
- `delete` - (Defaults to 10 minutes) Used when deleting the address

## Import

An address can be imported using the `/Accounts/{addressSid}/Addresses/{sid}` format, e.g.

```shell
terraform import twilio_account_address.address /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Addresses/ADXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
