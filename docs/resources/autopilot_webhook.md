---
page_title: "Twilio Autopilot Webhook"
subcategory: "Autopilot"
---

# twilio_autopilot_webhook Resource

Manages a Autopilot webhook. See the [API docs](https://www.twilio.com/docs/autopilot/api/event-webhooks) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
resource "twilio_autopilot_assistant" "assistant" {
  friendly_name = "twilio-test"
}

resource "twilio_autopilot_webhook" "webhook" {
  assistant_sid = twilio_autopilot_assistant.assistant.sid
  unique_name   = "twilio-test-webhook"
  webhook_url   = "http://localhost/webhook"
  events = [
    "onDialogueEnd"
  ]
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant to attach the webhook to. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the webhook
- `webhook_url` - (Mandatory) The webhook url
- `events` - (Mandatory) A list of webhook events strings which trigger the webhook
- `webhook_method` - (Optional) The HTTP method to trigger the webhook. Valid values are `POST` or `GET`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the webhook (Same as the SID)
- `sid` - The SID of the webhook (Same as the ID)
- `account_sid` - The account SID associated with the webhook
- `assistant_sid` - The SID of the assistant to attach the webhook to
- `unique_name` - The unique name of the webhook
- `webhook_url` - The webhook url
- `events` - A list of webhook events strings which trigger the webhook
- `webhook_method` - The HTTP method to trigger the webhook
- `date_created` - The date in RFC3339 format that the webhook was created
- `date_updated` - The date in RFC3339 format that the webhook was updated
- `url` - The url of the webhook resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the webhook
- `update` - (Defaults to 10 minutes) Used when updating the webhook
- `read` - (Defaults to 5 minutes) Used when retrieving the webhook
- `delete` - (Defaults to 10 minutes) Used when deleting the webhook

## Import

A webhook can be imported using the `/Assistants/{assistantSid}/Webhooks/{sid}}` format, e.g.

```shell
terraform import twilio_autopilot_webhook.webhook /Assistants/UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/UMXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
