---
page_title: "Twilio Programmable Messaging Short Codes"
subcategory: "Programmable Messaging"
---

# twilio_messaging_short_codes Resource

Use this data source to access information about the short codes associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api/short code-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_short_codes" "short_codes" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "short_codes" {
  value = data.twilio_messaging_short_codes.short_codes
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the messaging service the short codes are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the messaging service the short codes are associated with
- `account_sid` - The account SID associated with the short code
- `short_codes` - A list of `short_code` blocks as documented below

---

A `short_code` block supports the following:

- `sid` - The SID of the Twilio short code associated with the messaging service
- `capabilities` - The capabilities that are enabled for the short code
- `country_code` - The country code of the short code
- `short_code` - The short code
- `date_created` - The date in RFC3339 format that the messaging short code resource was created
- `date_updated` - The date in RFC3339 format that the messaging short code resource was updated
- `url` - The URL of the messaging short code resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving short codes
