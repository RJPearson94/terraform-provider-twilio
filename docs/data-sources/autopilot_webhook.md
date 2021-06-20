---
page_title: "Twilio Autopilot Webhook"
subcategory: "Autopilot"
---

# twilio_autopilot_webhook Data Source

Use this data source to access information about an existing Autopilot webhook. See the [API docs](https://www.twilio.com/docs/autopilot/api/event-webhooks) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

### SID

```hcl
data "twilio_autopilot_webhook" "webhook" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "UMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhook" {
  value = data.twilio_autopilot_webhook.webhook
}
```

### Unique Name

```hcl
data "twilio_autopilot_webhook" "webhook" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  unique_name   = "UniqueName"
}

output "webhook" {
  value = data.twilio_autopilot_webhook.webhook
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the webhook is associated with
- `sid` - (Optional) The SID of the webhook
- `unique_name` - (Optional) The unique name of the webhook

~> Either `sid` or `unique_name` must be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the webhook (Same as the `sid`)
- `sid` - The SID of the webhook (Same as the `id`)
- `account_sid` - The account SID associated with the webhook
- `assistant_sid` - The SID of the assistant to attach the webhook to
- `unique_name` - The unique name of the webhook
- `webhook_url` - The webhook URL
- `events` - A list of webhook events strings which trigger the webhook
- `webhook_method` - The HTTP method to trigger the webhook
- `date_created` - The date in RFC3339 format that the webhook was created
- `date_updated` - The date in RFC3339 format that the webhook was updated
- `url` - The URL of the webhook resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
