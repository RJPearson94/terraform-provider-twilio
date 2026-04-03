---
page_title: "twilio_serverless_build Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_build Resource

Manages a Serverless build. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/build) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

~> If polling is enabled then the create step will poll until the build status is either `completed` or `failed` or the max attempts threshold is reached.

~> To allow terraform to correctly manage the lifecycle of the deployment, it is recommended that use the lifecycle meta-argument `create_before_destroy` with this resource. The docs can be found [here](https://www.terraform.io/docs/configuration/resources.html#create_before_destroy)

!> If the `dependencies` are managed via Terraform and the `dependencies` are removed from the configuration file. The old value will be retained on the next apply

!> If the `runtime` is managed via Terraform and the `runtime` is removed from the configuration file. The old value will be retained on the next apply.

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"

  content           = <<EOF
exports.handler = function (context, event, callback) {
  callback(null, "Hello World");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}

resource "twilio_serverless_build" "build" {
  service_sid = twilio_serverless_service.service.sid

  function_version {
    sid = twilio_serverless_function.function.latest_version_sid
  }

  dependencies = {
    "twilio"                  = "3.6.3"
    "fs"                      = "0.0.1-security"
    "lodash"                  = "4.17.11"
    "util"                    = "0.11.0"
    "xmldom"                  = "0.1.27"
    "@twilio/runtime-handler" = "1.0.1"
  }

  polling {
    enabled = true
  }

  lifecycle {
    create_before_destroy = true
  }
}
```

~> When creating the build Twilio are currently injecting dependencies (at the time of writing it is `twilio`, `fs`, `lodash`, `util`, `xmldom` & `@twilio/runtime-handler`). If you need custom dependencies please ensure all dependencies (the ones needed for your app and Twilio supplied) are added to your terraform config otherwise the terraform config and state will not match after applying the changes

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service. Changing this forces a new resource

### Optional

- `asset_version` (Block List) A list of asset versions to include in the build. Changing this forces a new resource (see [below for nested schema](#nestedblock--asset_version))
- `dependencies` (Map of String) A map of npm package names to version ranges to include in the build. Changing this forces a new resource
- `function_version` (Block List) A list of function versions to include in the build. Changing this forces a new resource (see [below for nested schema](#nestedblock--function_version))
- `polling` (Block List, Max: 1) Configuration for polling the build status until it completes (see [below for nested schema](#nestedblock--polling))
- `runtime` (String) The Node.js runtime version for the build. Valid values are `node16`, `node18`, `node20`, `node22`. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `triggers` (Map of String) A map of arbitrary keys and values that, when changed, will trigger a new build. Changing this forces a new resource

### Read-Only

- `account_sid` (String) The SID of the account that owns this build
- `date_created` (String) The date and time the build was created, in RFC 3339 format
- `date_updated` (String) The date and time the build was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this build by Twilio
- `status` (String) The current status of the build
- `url` (String) The absolute URL of the build resource

<a id="nestedblock--asset_version"></a>
### Nested Schema for `asset_version`

Required:

- `sid` (String) The SID of the asset version to include in the build. Changing this forces a new resource

Read-Only:

- `account_sid` (String) The SID of the account that owns this asset version
- `asset_sid` (String) The SID of the asset that this version belongs to
- `date_created` (String) The date and time the asset version was created, in RFC 3339 format
- `path` (String) The URL path of the asset version
- `service_sid` (String) The SID of the Serverless service that owns this asset version
- `visibility` (String) The access control for the asset version


<a id="nestedblock--function_version"></a>
### Nested Schema for `function_version`

Required:

- `sid` (String) The SID of the function version to include in the build. Changing this forces a new resource

Read-Only:

- `account_sid` (String) The SID of the account that owns this function version
- `date_created` (String) The date and time the function version was created, in RFC 3339 format
- `function_sid` (String) The SID of the function that this version belongs to
- `path` (String) The URL path of the function version
- `service_sid` (String) The SID of the Serverless service that owns this function version
- `visibility` (String) The access control for the function version


<a id="nestedblock--polling"></a>
### Nested Schema for `polling`

Required:

- `enabled` (Boolean) Whether to poll the build status until it completes or fails

Optional:

- `delay_in_ms` (Number) The delay in milliseconds between each polling attempt. Defaults to `1000`
- `max_attempts` (Number) The maximum number of polling attempts before timing out. Defaults to `30`


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A build can be imported using the `/Services/{serviceSid}/Builds/{sid}` format, e.g.

```shell
terraform import twilio_serverless_build.build /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
