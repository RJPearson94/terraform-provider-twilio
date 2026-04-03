---
page_title: "twilio_serverless_builds Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_builds Data Source

Use this data source to access information about the builds associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/build) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_builds" "builds" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "builds" {
  value = data.twilio_serverless_builds.builds
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these builds
- `builds` (List of Object) A list of builds belonging to the Serverless service (see [below for nested schema](#nestedatt--builds))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--builds"></a>
### Nested Schema for `builds`

Read-Only:

- `asset_versions` (List of Object) (see [below for nested schema](#nestedobjatt--builds--asset_versions))
- `date_created` (String)
- `date_updated` (String)
- `dependencies` (Map of String)
- `function_versions` (List of Object) (see [below for nested schema](#nestedobjatt--builds--function_versions))
- `runtime` (String)
- `sid` (String)
- `status` (String)
- `url` (String)

<a id="nestedobjatt--builds--asset_versions"></a>
### Nested Schema for `builds.asset_versions`

Read-Only:

- `account_sid` (String)
- `asset_sid` (String)
- `date_created` (String)
- `path` (String)
- `service_sid` (String)
- `sid` (String)
- `visibility` (String)


<a id="nestedobjatt--builds--function_versions"></a>
### Nested Schema for `builds.function_versions`

Read-Only:

- `account_sid` (String)
- `date_created` (String)
- `function_sid` (String)
- `path` (String)
- `service_sid` (String)
- `sid` (String)
- `visibility` (String)
