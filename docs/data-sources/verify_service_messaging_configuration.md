---
page_title: "Twilio Verify Messaging Configuration"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `country_code` - (Mandatory) The country code the messaging configuration is for
- `service_sid` - (Mandatory) The service SID the messaging configuration is apart of

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the messaging configuration
- `account_sid` - The account SID of the messaging configuration is apart of
- `country_code` - The country code the messaging configuration is associated with
- `messaging_service_sid` - The messaging service SID which will be associated with the messaging configuration
- `service_sid` - The service SID the messaging configuration is associated with
- `date_created` - The date in RFC3339 format that the messaging configuration was created
- `date_updated` - The date in RFC3339 format that the messaging configuration was updated
- `url` - The URL of the messaging configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the messaging configuration
