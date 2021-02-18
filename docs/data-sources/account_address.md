---
page_title: "Twilio Account Address"
subcategory: "Account"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the address is associated with
- `sid` - (Mandatory) The SID of the address

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

- `read` - (Defaults to 5 minutes) Used when retrieving the address details
