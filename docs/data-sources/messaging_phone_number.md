---
page_title: "Twilio Programmable Messaging Phone Number"
subcategory: "Programmable Messaging"
---

# twilio_messaging_phone_number Data Source

Use this data source to access information about an existing Programmable Messaging phone number. See the [API docs](https://www.twilio.com/docs/sms/services/api/phonenumber-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_phone_number" "phone_number" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "PNXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_number" {
  value = data.twilio_messaging_phone_number.phone_number
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the phone number is associated with
- `sid` - (Mandatory) The SID of the phone number

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the phone number (Same as the SID)
- `sid` - The SID of the Twilio phone number to associated with the messaging service (Same as the ID)
- `service_sid` - The messaging service SID associated with the phone number
- `account_sid` - The account SID associated with the phone number
- `capabilities` - The capabilities that are enabled for the phone number
- `country_code` - The country code of the phone number
- `phone_number` - The phone number
- `date_created` - The date in RFC3339 format that the messaging phone number resource was created
- `date_updated` - The date in RFC3339 format that the messaging phone number resource was updated
- `url` - The url of the messaging phone number resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the phone number resource
