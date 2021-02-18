---
page_title: "Twilio Programmable Messaging Alpha Sender"
subcategory: "Programmable Messaging"
---

# twilio_messaging_alpha_sender Resource

Manages a Programmable Messaging alphanumeric sender resource. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_messaging_service" "service" {
  friendly_name = "twilio-test"
}

resource "twilio_messaging_alpha_sender" "alpha_sender" {
  service_sid  = twilio_messaging_service.service.sid
  alpha_sender = "test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The messaging service SID to associate the alpha sender with. Changing this forces a new resource to be created
- `alpha_sender` - (Mandatory) The alpha sender name to associate with the messaging service. Changing this forces a new resource to be created

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

- `create` - (Defaults to 10 minutes) Used when creating the alpha sender
- `read` - (Defaults to 5 minutes) Used when retrieving the alpha sender
- `delete` - (Defaults to 10 minutes) Used when deleting the alpha sender

## Import

A alpha sender can be imported using the `/Services/{serviceSid}/AlphaSenders/{sid}` format, e.g.

```shell
terraform import twilio_messaging_alpha_sender.alpha_sender /Services/MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/AlphaSenders/AIXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
