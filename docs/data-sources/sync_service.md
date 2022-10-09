---
page_title: "Twilio Sync Service"
subcategory: "Sync"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service

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

- `read` - (Defaults to 5 minutes) Used when retrieving the service
