---
page_title: "Twilio Serverless Deployment"
subcategory: "Serverless"
---

# twilio_serverless_deployment Resource

Manages a Serverless deployment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

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
  }

  polling {
    enabled = true
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
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID associated with the deployment. Changing this forces a new resource to be created
- `environment_sid` - (Mandatory) The Environment SID associated with the deployment. Changing this forces a new resource to be created
- `build_sid` - (Optional) The Build SID to be deployed to the environment. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the deployment (Same as the SID)
- `sid` - The SID of the deployment (Same as the ID)
- `account_sid` - The Account SID associated with the deployment
- `service_sid` - The Service SID associated with the deployment
- `environment_sid` - The Environment SID associated with the deployment
- `build_sid` - The Build SID to be deployed to the environment
- `date_created` - The date in RFC3339 format that the deployment was created
- `date_updated` - The date in RFC3339 format that the deployment was updated
- `url` - The url of the deployment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the deployment
- `read` - (Defaults to 5 minutes) Used when retrieving the deployment
- `delete` - (Defaults to 10 minutes) Used when deleting the deployment

## Import

A deployment can be imported using the `/Services/{serviceSid}/Environments/{environmentSid}/Deployments/{sid}` format, e.g.

```shell
terraform import twilio_serverless_deployment.deployment /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Deployments/ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
