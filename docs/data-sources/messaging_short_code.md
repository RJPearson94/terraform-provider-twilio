---
page_title: "Twilio Programmable Messaging Short Code"
subcategory: "Programmable Messaging"
---

# twilio_messaging_short_code Resource

Use this data source to access information about an existing Programmable Messaging short code resource. See the [API docs](https://www.twilio.com/docs/sms/services/api/shortcode-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_short_code" "short_code" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_code" {
  value = data.twilio_messaging_short_code.short_code
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the short code is associated with
- `sid` - (Mandatory) The SID of the short code

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the short code (Same as the SID)
- `sid` - The SID of the Twilio short code to associated with the messaging service (Same as the ID)
- `service_sid` - The messaging service SID associated with the short code
- `account_sid` - The account SID associated with the short code
- `capabilities` - The capabilities that are enabled for the short code
- `country_code` - The country code of the short code
- `short_code` - The short code
- `date_created` - The date in RFC3339 format that the messaging short code resource was created
- `date_updated` - The date in RFC3339 format that the messaging short code resource was updated
- `url` - The url of the messaging short code resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the short code
