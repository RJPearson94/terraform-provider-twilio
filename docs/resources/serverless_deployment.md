---
page_title: "twilio_serverless_deployment Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_deployment Resource

Manages a Serverless deployment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

~> Serverless deployments cannot be removed, they can only be superseded. To allow a build to be deleted, on the destruction of the resource, the provider will check if the `build_sid` is deployed to the environment. If the `build_sid` matches the environment config, a new deployment will be created without a `build_sid` to remove the active deployment. Once the deployment has been completed or if the `build_sid` doesn't match the environment, the state is removed and the deployment is orphaned.

~> To allow terraform to correctly manage the lifecycle of the deployment, it is recommended that use the lifecycle meta-argument `create_before_destroy` with this resource. The docs can be found [here](https://www.terraform.io/docs/configuration/resources.html#create_before_destroy)

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
    "twilio" : "3.6.3"
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

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "test"
}

resource "twilio_serverless_deployment" "deployment" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  build_sid       = twilio_serverless_build.build.sid

  lifecycle {
    create_before_destroy = true
  }
}
```

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment. Changing this forces a new resource
- `service_sid` (String) The SID of the Serverless service. Changing this forces a new resource

### Optional

- `build_sid` (String) The SID of the build to deploy to the environment. Changing this forces a new resource
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `triggers` (Map of String) A map of arbitrary keys and values that, when changed, will trigger a new deployment. Changing this forces a new resource

### Read-Only

- `account_sid` (String) The SID of the account that owns this deployment
- `date_created` (String) The date and time the deployment was created, in RFC 3339 format
- `date_updated` (String) The date and time the deployment was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `is_latest_deployment` (Boolean) Whether this deployment is the most recent deployment for the environment
- `sid` (String) The unique SID assigned to this deployment by Twilio
- `url` (String) The absolute URL of the deployment resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A deployment can be imported using the `/Services/{serviceSid}/Environments/{environmentSid}/Deployments/{sid}` format, e.g.

```shell
terraform import twilio_serverless_deployment.deployment /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `triggers` cannot be imported
