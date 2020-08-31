---
page_title: "Twilio Programmable Messaging Alpha Senders"
subcategory: "Programmable Messaging"
---

# twilio_messaging_alpha_senders Data Source

Use this data source to access information about the alphanumeric senders associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api/alphasender-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_alpha_senders" "alpha_senders" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "alpha_senders" {
  value = data.twilio_messaging_alpha_senders.alpha_senders
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the alpha senders are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the messaging service the alpha senders are associated with
- `account_sid` - The account SID associated with the alpha senders
- `alpha_senders` - A list of `alpha_sender` blocks as documented below

---

An `alpha_sender` block supports the following:

- `sid` - The SID of the alpha sender resource
- `capabilities` - The capabilities that are enabled for the alpha sender
- `alpha_sender` - The alpha sender name associated with the messaging service
- `date_created` - The date in RFC3339 format that the messaging alpha sender resource was created
- `date_updated` - The date in RFC3339 format that the messaging alpha sender resource was updated
- `url` - The URL of the messaging alpha sender resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the alpha sender
