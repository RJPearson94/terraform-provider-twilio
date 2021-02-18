---
page_title: "Twilio Programmable Messaging Alpha Sender"
subcategory: "Programmable Messaging"
---

# twilio_messaging_alpha_sender Data Source

Use this data source to access information about an existing Programmable Messaging alphanumeric sender. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "alpha_sender" {
  value = data.twilio_messaging_alpha_sender.alpha_sender
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the alpha sender is associated with
- `sid` - (Mandatory) The SID of the alpha sender

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the alpha sender resource (Same as the `sid`)
- `sid` - The SID of the alpha sender resource (Same as the `id`)
- `service_sid` - The messaging service SID associated with the alpha sender
- `account_sid` - The account SID associated with the alpha sender
- `capabilities` - The capabilities that are enabled for the alpha sender
- `alpha_sender` - The alpha sender name associated with the messaging service
- `date_created` - The date in RFC3339 format that the messaging alpha sender resource was created
- `date_updated` - The date in RFC3339 format that the messaging alpha sender resource was updated
- `url` - The URL of the messaging alpha sender resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the alpha sender
