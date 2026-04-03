---
page_title: "twilio_proxy_service Resource - twilio"
subcategory: "Proxy"
description: |-
  
---

# twilio_proxy_service Resource

Manages a Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/service) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_proxy_service" "service" {
  unique_name = "twilio-test"
}
```

## Schema

### Required

- `unique_name` (String) A unique, developer-assigned name for the Proxy service. Must be between 1 and 191 characters

### Optional

- `callback_url` (String) The URL to receive callback events for the Proxy service
- `chat_instance_sid` (String) The SID of the Chat service instance to associate with the Proxy service
- `default_ttl` (Number) The default time-to-live (TTL) for sessions in the Proxy service, in seconds. Defaults to `0`
- `geo_match_level` (String) The geographic area for matching proxy numbers. Valid values are `area-code`, `country`, or `extended-area-code`. Defaults to `country`
- `intercept_callback_url` (String) The URL to receive intercept callback events for the Proxy service
- `number_selection_behavior` (String) The behavior for selecting proxy numbers. Valid values are `avoid-sticky` or `prefer-sticky`. Defaults to `prefer-sticky`
- `out_of_session_callback_url` (String) The URL to receive out-of-session callback events for the Proxy service
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Proxy service
- `date_created` (String) The date and time the Proxy service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Proxy service was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this Proxy service by Twilio
- `url` (String) The absolute URL of the Proxy service resource

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
terraform import twilio_proxy_service.service /Services/KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
