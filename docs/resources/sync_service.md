---
page_title: "twilio_sync_service Resource - twilio"
subcategory: "Sync"
description: |-
  
---

# twilio_sync_service Resource

Manages a Sync service. See the [API docs](https://www.twilio.com/docs/sync/api/service) for more information

For more information on Sync, see the product [page](https://www.twilio.com/sync)

## Example Usage

### Basic

```hcl
resource "twilio_sync_service" "service" {}
```

### With Friendly Name

```hcl
resource "twilio_sync_service" "service" {
  friendly_name = "Test Sync Service"
}
```

## Schema

### Optional

- `acl_enabled` (Boolean) Whether token identities in the Sync service must be granted access to Sync objects via the Permissions API. Defaults to `false`
- `friendly_name` (String) A human-readable label for the Sync service
- `reachability_debouncing_enabled` (Boolean) Whether every endpoint_disconnected event should fire after a configurable delay. Defaults to `false`
- `reachability_debouncing_window` (Number) The reachability event delay, in milliseconds, between 1000 and 30000. Defaults to `5000`
- `reachability_webhooks_enabled` (Boolean) Whether the service instance should call the webhook_url when client endpoints connect or disconnect from Sync. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `webhook_url` (String) The URL to which Sync will send webhooks. Must be a valid HTTP or HTTPS URL
- `webhooks_from_rest_enabled` (Boolean) Whether the service instance should call the webhook_url when the REST API is used to update Sync objects. Defaults to `false`

### Read-Only

- `account_sid` (String) The SID of the account that owns this Sync service
- `date_created` (String) The date and time the Sync service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Sync service was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this Sync service by Twilio
- `url` (String) The absolute URL of the Sync service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_sync_service.service /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
