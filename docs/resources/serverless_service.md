---
page_title: "twilio_serverless_service Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_service Resource

Manages a Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/service) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the service
- `unique_name` (String) A unique, developer-assigned name for the Serverless service. Changing this forces a new resource

### Optional

- `include_credentials` (Boolean) Whether to inject account credentials into a function invocation context. Defaults to `true`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `ui_editable` (Boolean) Whether the service's properties and sub-resources can be edited in the Twilio Console. Defaults to `false`

### Read-Only

- `account_sid` (String) The SID of the account that owns this service
- `date_created` (String) The date and time the service was created, in RFC 3339 format
- `date_updated` (String) The date and time the service was last updated, in RFC 3339 format
- `domain_base` (String) The base domain name for this service, used to compose the URLs of the service's environments
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this service by Twilio
- `url` (String) The absolute URL of the service resource

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
terraform import twilio_serverless_service.service /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
