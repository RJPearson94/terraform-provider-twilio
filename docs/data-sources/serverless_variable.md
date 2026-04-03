---
page_title: "twilio_serverless_variable Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_variable Data Source

Use this data source to access information about an existing Serverless environment variable. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_variable" "variable" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid             = "ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "variable" {
  value = data.twilio_serverless_variable.variable
}
```

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment
- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless variable

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this variable
- `date_created` (String) The date and time the variable was created, in RFC 3339 format
- `date_updated` (String) The date and time the variable was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `key` (String) The name of the environment variable
- `url` (String) The absolute URL of the variable resource
- `value` (String) The value of the environment variable

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
