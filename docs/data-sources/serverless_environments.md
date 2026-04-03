---
page_title: "twilio_serverless_environments Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_environments Data Source

Use this data source to access information about the environments associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/environment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_environments" "environments" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "environments" {
  value = data.twilio_serverless_environments.environments
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these environments
- `environments` (List of Object) A list of environments belonging to the Serverless service (see [below for nested schema](#nestedatt--environments))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--environments"></a>
### Nested Schema for `environments`

Read-Only:

- `build_sid` (String)
- `date_created` (String)
- `date_updated` (String)
- `domain_name` (String)
- `domain_suffix` (String)
- `sid` (String)
- `unique_name` (String)
- `url` (String)
