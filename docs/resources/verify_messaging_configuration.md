---
page_title: "Twilio Verify Messaging Configuration"
subcategory: "Verify"
---

# twilio_verify_messaging_configuration Resource

Manages a Verify messaging configuration

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_messaging_service" "service" {
  friendly_name = "Test Messaging Service"
}

resource "twilio_verify_messaging_configuration" "messaging_configuration" {
  service_sid           = twilio_verify_service.service.sid
  messaging_service_sid = twilio_messaging_service.service.sid
  country_code          = "GB"
}
```

## Argument Reference

The following arguments are supported:

- `country_code` - (Mandatory) The country code the messaging configuration will be used for. Changing this forces a new resource to be created
- `messaging_service_sid` - (Mandatory) The messaging service SID which will be associated with the messaging configuration
- `service_sid` - (Mandatory) The service SID to associate the messaging configuration with. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the messaging configuration
- `account_sid` - The account SID the messaging configuration is associated into
- `country_code` - The country code the messaging configuration is associated with
- `messaging_service_sid` - The messaging service SID which will be associated with the messaging configuration
- `service_sid` - The service SID the messaging configuration is associated with
- `date_created` - The date in RFC3339 format that the messaging configuration was created
- `date_updated` - The date in RFC3339 format that the messaging configuration was updated
- `url` - The URL of the messaging configuration

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the messaging configuration
- `update` - (Defaults to 10 minutes) Used when updating the messaging configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the messaging configuration
- `delete` - (Defaults to 10 minutes) Used when deleting the messaging configuration

## Import

A messaging configuration can be imported using the `/Services/{serviceSid}/MessagingConfigurations/{countryCode}` format, e.g.

```shell
terraform import twilio_verify_rate_limit.rate_limit /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/MessagingConfigurations/GB
```
