---
page_title: "Twilio Programmable Messaging Phone Numbers"
subcategory: "Programmable Messaging"
---

# twilio_messaging_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api/phonenumber-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_phone_numbers" "phone_numbers" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_messaging_phone_numbers.phone_numbers
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the messaging service the phone numbers are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `service_sid`)
- `service_sid` - The SID of the messaging service the phone numbers are associated with (Same as the `id`)
- `account_sid` - The account SID associated with the phone number
- `phone_numbers` - A list of `phone_number` blocks as documented below

---

A `phone_number` block supports the following:

- `sid` - The SID of the Twilio phone number associated with the messaging service
- `capabilities` - The capabilities that are enabled for the phone number
- `country_code` - The country code of the phone number
- `phone_number` - The phone number
- `date_created` - The date in RFC3339 format that the messaging phone number resource was created
- `date_updated` - The date in RFC3339 format that the messaging phone number resource was updated
- `url` - The URL of the messaging phone number resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving phone numbers resource
