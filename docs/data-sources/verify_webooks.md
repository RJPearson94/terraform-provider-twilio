---
page_title: "Twilio Verify Webhooks"
subcategory: "Verify"
---

# twilio_verify_webhooks Data Source

Use this data source to access information about existing Verify webhooks. See the [API docs](https://www.twilio.com/docs/verify/api/webhooks) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_webhooks" "webhooks" {
  service_sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "webhooks" {
  value = data.twilio_verify_webhooks.webhooks
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID the webhooks are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the webhooks (Same as the `service_sid`)
- `account_sid` - The account SID of the webhooks are associated with
- `service_sid` - The service SID the webhooks are associated with (Same as the `id`)
- `webhooks` - A list of `webhook` blocks as documented below

---

A `webhook` block supports the following:

- `sid` - The SID of the webhook
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

- `read` - (Defaults to 10 minutes) Used when retrieving the webhooks
