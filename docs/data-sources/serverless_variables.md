---
page_title: "twilio_serverless_variables Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_variables Data Source

Use this data source to access information about the variables associated with an existing Serverless service and environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_variables" "variables" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "variables" {
  value = data.twilio_serverless_variables.variables
}
```

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment
- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these variables
- `id` (String) The ID of this resource.
- `variables` (List of Object) A list of variables belonging to the Serverless environment (see [below for nested schema](#nestedatt--variables))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `key` (String)
- `sid` (String)
- `url` (String)
- `value` (String)
