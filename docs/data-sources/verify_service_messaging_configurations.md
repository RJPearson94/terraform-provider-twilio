---
page_title: "Twilio Verify Messaging Configurations"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID the messaging configurations are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the messaging configurations (Same as the `service_sid`)
- `account_sid` - The account SID of the messaging configurations are associated with
- `service_sid` - The service SID the messaging configurations are associated with (Same as the `id`)
- `messaging_configurations` - A list of `messaging_configuration` blocks as documented below

---

A `messaging_configuration` block supports the following:

- `country_code` - The country code the messaging configuration is associated with
- `messaging_service_sid` - The messaging service SID associated with the messaging configuration
- `date_created` - The date in RFC3339 format that the messaging configuration was created
- `date_updated` - The date in RFC3339 format that the messaging configuration was updated
- `url` - The URL of the messaging configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configurations/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the messaging configurations
