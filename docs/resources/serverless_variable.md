---
page_title: "twilio_serverless_variable Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_variable Resource

Manages a Serverless environment variable. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

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

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "test-key"
  value           = "test-value"
}
```

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment. Changing this forces a new resource
- `key` (String) The name of the environment variable
- `service_sid` (String) The SID of the Serverless service. Changing this forces a new resource
- `value` (String) The value of the environment variable

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this variable
- `date_created` (String) The date and time the variable was created, in RFC 3339 format
- `date_updated` (String) The date and time the variable was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this variable by Twilio
- `url` (String) The absolute URL of the variable resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A variable can be imported using the `/Services/{serviceSid}/Environments/{environmentSid}/Variables/{sid}` format, e.g.

```shell
terraform import twilio_serverless_variable.variable /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
