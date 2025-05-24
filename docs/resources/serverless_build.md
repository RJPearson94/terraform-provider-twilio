---
page_title: "Twilio Serverless Build"
subcategory: "Serverless"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The serverless service SID to associate the build with. Changing this forces a new resource to be created
- `asset_version` - (Optional) A `asset_version` block as documented below. Changing this forces a new resource to be created
- `function_version` - (Optional) A `function_version` block as documented below. Changing this forces a new resource to be created
- `dependencies` - (Optional) Map of dependencies to be included in the build. Changing this forces a new resource to be created
- `runtime` - (Optional) The target runtime of the serverless functions and assets. Valid values are `node16`, `node18`, `node20` or `node22`. Changing this forces a new resource to be created
- `polling` - (Optional) A `polling` block as documented below.
- `triggers` - (Optional) A map of key-value pairs which can be used to determine if changes have occurred and redeployment is necessary. Changing this forces a new resource to be created
  ~> An alternative strategy is to use the [taint](https://www.terraform.io/docs/commands/taint.html) functionality of Terraform.

---

An `asset_version` block supports the following:

- `sid` - (Required) The SID of the asset version. Changing this forces a new resource to be created

---

A `function_version` block supports the following:

- `sid` - (Required) The SID of the function version. Changing this forces a new resource to be created

---

A `polling` block supports the following:

- `enabled` - (Required) Enable or disable polling of the build.
- `max_attempts` - (Optional) The maximum number of polling attempts. Default is 30
- `delay_in_ms` - (Optional) The time in milliseconds to wait between polling attempts. Default is 1000ms

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the build (Same as the `sid`)
- `sid` - The SID of the build (Same as the `id`)
- `account_sid` - The account SID associated with the build
- `service_sid` - The service SID associated with the build
- `asset_version` - A `asset_version` block as documented below.
- `function_version` - A `function_version` block as documented below.
- `triggers` - A map of key-value pairs which can be used to determine if changes have occurred and redeployment is necessary.
- `dependencies` - Map of dependencies to be included in the build
- `runtime` - The target runtime of the serverless functions and assets
- `status` - The current status of the build job
- `date_created` - The date in RFC3339 format that the build was created
- `date_updated` - The date in RFC3339 format that the build was updated
- `url` - The URL of the build

---

An `asset_version` block supports the following:

- `sid` - The SID of the asset version
- `account_sid` - The account SID of the asset version is deployed into
- `service_sid` - The service SID of the asset version is deployed into
- `asset_sid` - The asset SID of the version is managed under
- `date_created` - The date in RFC3339 format that the asset version was created
- `path` - The request URI path
- `visibility` - The visibility of the asset

---

A `function_version` block supports the following:

- `sid` - The SID of the function version
- `account_sid` - The account SID of the function version is deployed into
- `service_sid` - The service SID of the function version is deployed into
- `function_sid` - The function SID of the version is managed under
- `date_created` - The date in RFC3339 format that the function version was created
- `path` - The request URI path
- `visibility` - The visibility of the function

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the build
- `read` - (Defaults to 5 minutes) Used when retrieving the build
- `delete` - (Defaults to 10 minutes) Used when deleting the build

!> When polling is enabled, each request is constrained by the read timeout defined above

## Import

A build can be imported using the `/Services/{serviceSid}/Builds/{sid}` format, e.g.

```shell
terraform import twilio_serverless_build.build /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Builds/ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
