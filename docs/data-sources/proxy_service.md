---
page_title: "twilio_proxy_service Data Source - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_service Data Source

Use this data source to access information about an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/service) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_service" "service" {
  sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_proxy_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the Proxy service to read

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy service
- `callback_url` (String) The URL to receive callback events for the Proxy service
- `chat_instance_sid` (String) The SID of the Chat service instance associated with the Proxy service
- `date_created` (String) The date and time the Proxy service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Proxy service was last updated, in RFC 3339 format
- `default_ttl` (Number) The default time-to-live (TTL) for sessions in the Proxy service, in seconds
- `geo_match_level` (String) The geographic area for matching proxy numbers
- `id` (String) The ID of this resource.
- `intercept_callback_url` (String) The URL to receive intercept callback events for the Proxy service
- `number_selection_behavior` (String) The behavior for selecting proxy numbers
- `out_of_session_callback_url` (String) The URL to receive out-of-session callback events for the Proxy service
- `unique_name` (String) A unique, developer-assigned name for the Proxy service
- `url` (String) The absolute URL of the Proxy service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
