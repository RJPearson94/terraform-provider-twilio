---
page_title: "Twilio Sync Service"
subcategory: "Sync"
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

## Argument Reference

The following arguments are supported:

- `acl_enabled` - (Optional) Whether access control lists are enabled. The default value is `false`
- `friendly_name` - (Optional) The friendly name of the service
- `reachability_debouncing_enabled` - (Optional) Whether endpoint_disconnected event should occur when the reachability_debouncing_window is reached. The default value is `false`
- `reachability_debouncing_window` - (Optional) The reachability event delay in milliseconds. The value must be between `1000` and `30000` (inclusive). The default value is `5000`
- `reachability_webhooks_enabled` - (Optional) Whether the service should call the webhook url on client connections. The default value is `false`
- `webhook_url` - (Optional) The URL called when Sync objects are changed
- `webhooks_from_rest_enabled` - (Optional) Whether the service should call the webhook url on updates via the REST API. The default value is `false`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The account SID the service is associated with
- `acl_enabled` - Whether access control lists are enabled
- `friendly_name` - The friendly name of the service
- `reachability_debouncing_enabled` - Whether endpoint_disconnected event should occur when the reachability_debouncing_window is reached
- `reachability_debouncing_window` - The reachability event delay in milliseconds
- `reachability_webhooks_enabled` - Whether the service should call the webhook url on client connections
- `webhook_url` - The URL called when Sync objects are changed
- `webhooks_from_rest_enabled` - Whether the service should call the webhook url on updates via the REST API
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the service
- `update` - (Defaults to 10 minutes) Used when updating the service
- `read` - (Defaults to 5 minutes) Used when retrieving the service
- `delete` - (Defaults to 10 minutes) Used when deleting the service

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_sync_service.service /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
