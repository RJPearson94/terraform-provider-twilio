---
page_title: "Twilio Autopilot Webhooks"
subcategory: "Autopilot"
---

# twilio_autopilot_webhooks Data Source

Use this data source to access information about the webhooks associated with an existing Autopilot assistant. See the [API docs](https://www.twilio.com/docs/autopilot/api/event-webhooks) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_webhooks" "webhooks" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhooks" {
  value = data.twilio_autopilot_webhooks.webhooks
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the webhooks are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the assistant SID)
- `account_sid` - The SID of the account the webhooks are associated with
- `assistant_sid` - The SID of the assistant the webhooks are associated with
- `webhooks` - A list of `webhook` blocks as documented below

---

A `webhook` block supports the following:

- `sid` - The SID of the webhook (Same as the ID)
- `unique_name` - The unique name of the webhook
- `webhook_url` - The webhook url
- `events` - A list of webhook events strings which trigger the webhook
- `webhook_method` - The HTTP method to trigger the webhook
- `date_created` - The date in RFC3339 format that the webhook was created
- `date_updated` - The date in RFC3339 format that the webhook was updated
- `url` - The url of the webhook resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving webhooks
