---
page_title: "Twilio Verify Webhook"
subcategory: "Verify"
---

# twilio_verify_webhook Data Source

Use this data source to access information about an existing Verify webhook. See the [API docs](https://www.twilio.com/docs/verify/api/webhooks) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_webhook" "webhook" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhook" {
  value = data.twilio_verify_webhook.webhook
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the webhook
- `service_sid` - (Mandatory) The service SID the webhook is associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the webhook (Same as the `sid`)
- `sid` - The SID of the webhook (Same as the `id`)
- `account_sid` - The account SID of the webhook is associated with
- `service_sid` - The service SID the webhook is associated with
- `friendly_name` - The friendly name of the webhook
- `event_types` - The list of events which trigger a webhook call
- `status` - The webhook status
- `version` - The webhook version.
- `webhook_url` - The webhook URL
- `webhook_method` - The HTTP method to trigger the webhook
- `date_created` - The date in RFC3339 format that the webhook was created
- `date_updated` - The date in RFC3339 format that the webhook was updated
- `url` - The URL of the webhook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
