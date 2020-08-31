---
page_title: "Twilio Serverless Environment"
subcategory: "Serverless"
---

# twilio_serverless_environment Resource

Manages a Serverless environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/environment) for more information

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
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The serverless service SID to associate the environment with. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the environment
- `domain_suffix` - (Optional) The domain suffix of the environment

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the environment (Same as the SID)
- `sid` - The SID of the environment (Same as the ID)
- `account_sid` - The account SID of the environment is deployed into
- `service_sid` - The service SID of the environment is managed under
- `build_sid` - The build SID of the current build deployed to the environment
- `unique_name` - The unique name of the environment
- `domain_suffix` - The domain suffix of the environment
- `domain_name` - The domain name of the environment
- `date_created` - The date in RFC3339 format that the environment was created
- `date_updated` - The date in RFC3339 format that the environment was updated
- `url` - The URL of the environment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the environment
- `read` - (Defaults to 5 minutes) Used when retrieving the environment
- `delete` - (Defaults to 10 minutes) Used when deleting the environment

## Import

A environment can be imported using the `/Services/{serviceSid}/Environments/{sid}` format, e.g.

```shell
terraform import twilio_serverless_environment.environment /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Environments/ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
