---
page_title: "twilio_serverless_functions Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_functions Data Source

Use this data source to access information about the functions associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/function) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_functions" "functions" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "functions" {
  value = data.twilio_serverless_functions.functions
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these functions
- `functions` (List of Object) A list of functions belonging to the Serverless service (see [below for nested schema](#nestedatt--functions))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--functions"></a>
### Nested Schema for `functions`

Read-Only:

- `content` (String)
- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `latest_version_sid` (String)
- `path` (String)
- `sid` (String)
- `url` (String)
- `visibility` (String)
