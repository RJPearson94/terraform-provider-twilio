---
page_title: "twilio_sync_service Data Source - twilio"
subcategory: "Sync"
description: |-
  
---

# twilio_sync_service Data Source

Use this data source to access information about an existing Sync service. See the [API docs](https://www.twilio.com/docs/sync/api/service) for more information

For more information on Sync, see the product [page](https://www.twilio.com/sync)

## Example Usage

```hcl
data "twilio_sync_service" "service" {
  sid = "ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_sync_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the Sync service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Sync service
- `acl_enabled` (Boolean) Whether token identities in the Sync service must be granted access to Sync objects via the Permissions API
- `date_created` (String) The date and time the Sync service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Sync service was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the Sync service
- `id` (String) The ID of this resource.
- `reachability_debouncing_enabled` (Boolean) Whether every endpoint_disconnected event fires after a configurable delay
- `reachability_debouncing_window` (Number) The reachability event delay, in milliseconds
- `reachability_webhooks_enabled` (Boolean) Whether the service instance calls the webhook_url when client endpoints connect or disconnect from Sync
- `url` (String) The absolute URL of the Sync service resource
- `webhook_url` (String) The URL to which Sync sends webhooks
- `webhooks_from_rest_enabled` (Boolean) Whether the service instance calls the webhook_url when the REST API is used to update Sync objects

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
