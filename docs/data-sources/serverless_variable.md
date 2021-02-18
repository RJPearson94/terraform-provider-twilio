---
page_title: "Twilio Serverless Variable"
subcategory: "Serverless"
---

# twilio_serverless_variable Data Source

Use this data source to access information about an existing Serverless environment variable. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_variable" "variable" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid             = "ZVXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "variable" {
  value = data.twilio_serverless_variable.variable
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the variable is associated with
- `environment_sid` - (Mandatory) The SID of the environment the variable is associated with
- `sid` - (Mandatory) The SID of the variable

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the environment variable (Same as the `sid`)
- `sid` - The SID of the environment variable (Same as the `id`)
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

- `read` - (Defaults to 5 minutes) Used when retrieving the environment variable
