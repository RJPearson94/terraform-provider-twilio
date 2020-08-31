---
page_title: "Twilio Serverless Variable"
subcategory: "Serverless"
---

# twilio_serverless_variable Resource

Manages a Serverless environment variable. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "test"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "test"
}

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "test-key"
  value           = "test-value"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The serverless service SID to associate the environment variable with. Changing this forces a new resource to be created
- `environment_sid` - (Mandatory) The serverless environment SID to associate the environment variable with. Changing this forces a new resource to be created
- `key` - (Mandatory) The environment variable key
- `value` - (Mandatory) The environment variable value

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the environment variable (Same as the SID)
- `sid` - The SID of the environment variable (Same as the ID)
- `account_sid` - The account SID of the environment variable is deployed into
- `service_sid` - The service SID of the environment variable is deployed into
- `environment_sid` - The environment SID of the environment variable is managed under
- `key` - The environment variable key
- `value` - The environment variable value
- `date_created` - The date in RFC3339 format that the environment variable was created
- `date_updated` - The date in RFC3339 format that the environment variable was updated
- `url` - The URL of the environment variable

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the environment variable
- `update` - (Defaults to 10 minutes) Used when updating the environment variable
- `read` - (Defaults to 5 minutes) Used when retrieving the environment variable
- `delete` - (Defaults to 10 minutes) Used when deleting the environment variable

## Import

A variable can be imported using the `/Services/{serviceSid}/Environments/{environmentSid}/Variables/{sid}` format, e.g.

```shell
terraform import twilio_serverless_variable.variable /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Variables/ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
