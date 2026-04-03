---
page_title: "twilio_serverless_environment Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_environment Resource

Manages a Serverless environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/environment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "test"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "test"
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service. Changing this forces a new resource
- `unique_name` (String) A unique, developer-assigned name for the environment. Changing this forces a new resource

### Optional

- `domain_suffix` (String) A URL-friendly suffix appended to the environment's domain name. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this environment
- `build_sid` (String) The SID of the build currently deployed to this environment
- `date_created` (String) The date and time the environment was created, in RFC 3339 format
- `date_updated` (String) The date and time the environment was last updated, in RFC 3339 format
- `domain_name` (String) The fully qualified domain name for this environment
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this environment by Twilio
- `url` (String) The absolute URL of the environment resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A environment can be imported using the `/Services/{serviceSid}/Environments/{sid}` format, e.g.

```shell
terraform import twilio_serverless_environment.environment /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
