---
page_title: "twilio_verify_webhook Data Source - twilio"
subcategory: "Verify"
description: |-
  
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


## Schema

### Required

- `service_sid` (String) The SID of the Verify service
- `sid` (String) The SID of the webhook to fetch

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this webhook
- `date_created` (String) The date and time the webhook was created, in RFC 3339 format
- `date_updated` (String) The date and time the webhook was last updated, in RFC 3339 format
- `event_types` (List of String) The list of events that trigger the webhook
- `friendly_name` (String) A human-readable label for the webhook
- `id` (String) The ID of this resource.
- `status` (String) The status of the webhook
- `url` (String) The absolute URL of the webhook resource
- `version` (String) The webhook version
- `webhook_method` (String) The HTTP method used when calling the webhook URL
- `webhook_url` (String) The HTTPS URL that Twilio calls when an event occurs

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
