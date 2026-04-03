---
page_title: "twilio_serverless_build Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_build Data Source

Use this data source to access information about an existing Serverless build. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/build) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_build" "build" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "build" {
  value = data.twilio_serverless_build.build
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless build

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this build
- `asset_versions` (List of Object) A list of asset versions included in this build (see [below for nested schema](#nestedatt--asset_versions))
- `date_created` (String) The date and time the build was created, in RFC 3339 format
- `date_updated` (String) The date and time the build was last updated, in RFC 3339 format
- `dependencies` (Map of String) A map of npm package names to version ranges included in the build
- `function_versions` (List of Object) A list of function versions included in this build (see [below for nested schema](#nestedatt--function_versions))
- `id` (String) The ID of this resource.
- `runtime` (String) The Node.js runtime version used for the build
- `status` (String) The current status of the build
- `url` (String) The absolute URL of the build resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--asset_versions"></a>
### Nested Schema for `asset_versions`

Read-Only:

- `account_sid` (String)
- `asset_sid` (String)
- `date_created` (String)
- `path` (String)
- `service_sid` (String)
- `sid` (String)
- `visibility` (String)


<a id="nestedatt--function_versions"></a>
### Nested Schema for `function_versions`

Read-Only:

- `account_sid` (String)
- `date_created` (String)
- `function_sid` (String)
- `path` (String)
- `service_sid` (String)
- `sid` (String)
- `visibility` (String)
