---
page_title: "Twilio Programmable Chat Channel Webhook"
subcategory: "Programmable Chat"
---

# twilio_chat_channel_webhook Data Source

Use this data source to access information about an existing Programmable Chat channel webhook. See the [API docs](https://www.twilio.com/docs/chat/rest/channel-webhook-resource) for more information

~> This is a generic data source which can be used to retrieve channel webhook info regardless of the type (webhook, trigger, studio)

For more information on Programmable Chat, see the product [page](https://www.twilio.com/chat)

## Example Usage

```hcl
data "twilio_chat_channel_webhook" "webhook" {
  service_sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  channel_sid = "CHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "WHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhook" {
  value = data.twilio_chat_channel_webhook.webhook
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the channel webhook is associated with
- `channel_sid` - (Mandatory) The SID of the channel the webhook is associated with
- `sid` - (Mandatory) The SID of the channel webhook

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the channel webhook (Same as the SID)
- `sid` - The SID of the channel webhook (Same as the ID)
- `account_sid` - The account SID associated with the channel webhook
- `service_sid` - The service SID associated with the channel webhook
- `channel_sid` - The channel SID associated with the channel webhook
- `type` - The type of webhook
- `configuration` - The `configuration` block as documented below
- `date_created` - The date in RFC3339 format that the channel webhook was created
- `date_updated` - The date in RFC3339 format that the channel webhook was updated
- `url` - The url of the channel webhook

---

A `configuration` block supports the following:

- `method` - The HTTP method to trigger the channel webhook
- `webhook_url` - The webhook url
- `filters` - The filter conditions that trigger the channel webhook
- `retry_count` - The number of attempt to retry a failed channel webhook call
- `flow_sid` - The SID for the studio flow which will be called
- `retry_count` - The number of attempt to retry a failed channel webhook call
- `triggers` - The keywords which trigger the channel webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the channel webhook
