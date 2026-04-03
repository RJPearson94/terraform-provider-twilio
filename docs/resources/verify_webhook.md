---
page_title: "twilio_verify_webhook Resource - twilio"
subcategory: "Verify"
description: |-
  
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

## Schema

### Required

- `event_types` (List of String) The list of events that trigger the webhook. Valid values: `*`, `factor.created`, `factor.verified`, `factor.deleted`, `challenge.approved`, `challenge.denied`
- `friendly_name` (String) A human-readable label for the webhook
- `service_sid` (String) The SID of the Verify service. Changing this forces a new resource
- `webhook_url` (String) The HTTPS URL that Twilio calls when an event occurs

### Optional

- `status` (String) The status of the webhook. Valid values: `enabled`, `disabled`. Defaults to `enabled`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `version` (String) The webhook version. Valid values: `v1`, `v2`. Defaults to `v2`

### Read-Only

- `account_sid` (String) The SID of the account that owns this webhook
- `date_created` (String) The date and time the webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the webhook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this webhook by Twilio
- `url` (String) The absolute URL of the webhook resource
- `webhook_method` (String) The HTTP method used when calling the webhook URL

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A webhook can be imported using the `/Services/{serviceSid}/Webhooks/{sid}` format, e.g.

```shell
terraform import twilio_verify_webhook.webhook /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Webhooks/YWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
