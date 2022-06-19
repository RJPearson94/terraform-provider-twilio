---
page_title: "Twilio Verify Webhook"
subcategory: "Verify"
---

# twilio_verify_webhook Resource

Manages a Verify webhook. See the [API docs](https://www.twilio.com/docs/verify/api/webhooks) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}

resource "twilio_verify_webhook" "webhook" {
  service_sid   = twilio_verify_service.service.sid
  friendly_name = "Test Verify Webhook"
  event_types   = ["*"]
  webhook_url   = "https://localhost.com/webhook"
}
```

## Argument Reference

The following arguments are supported:

- `event_types` - (Mandatory) The list of events which trigger a webhook call. Valid values are `*`, `factor.created`, `factor.verified`, `factor.deleted`, `challenge.approved` or `challenge.denied`
- `friendly_name` - (Mandatory) The friendly name of the webhook
- `service_sid` - (Mandatory) The service SID to associate the webhook with. Changing this forces a new resource to be created
- `webhook_url` - (Mandatory) The webhook URL
  ~> Webhook URL must use HTTPS
- `status` - (Optional) The webhook status. Valid values are `enabled` or `disabled`. The default value is `enabled`
- `version` - (Optional) The webhook version. Valid values are `v1` or `v2`. The default value is `v2`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the webhook (Same as the `sid`)
- `sid` - The SID of the webhook (Same as the `id`)
- `account_sid` - The account SID the webhook is associated with
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

- `create` - (Defaults to 10 minutes) Used when creating the webhook
- `update` - (Defaults to 10 minutes) Used when updating the webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
- `delete` - (Defaults to 10 minutes) Used when deleting the webhook

## Import

A webhook can be imported using the `/Services/{serviceSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_verify_webhook.webhook /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
