---
page_title: "twilio_serverless_function Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_function Data Source

Use this data source to access information about an existing Serverless function. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/function) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_function" "function" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "function" {
  value = data.twilio_serverless_function.function
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless function

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this function
- `content` (String) The content of the latest function version
- `date_created` (String) The date and time the function was created, in RFC 3339 format
- `date_updated` (String) The date and time the function was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the function
- `id` (String) The ID of this resource.
- `latest_version_sid` (String) The SID of the latest version of the function
- `path` (String) The URL path at which the function is accessible
- `url` (String) The absolute URL of the function resource
- `visibility` (String) The access control for the function

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
