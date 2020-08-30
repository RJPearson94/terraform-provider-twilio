---
page_title: "Twilio Serverless Variables"
subcategory: "Serverless"
---

# twilio_serverless_variables Data Source

Use this data source to access information about the variables associated with an existing Serverless service and environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/variable) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_variables" "variables" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "variables" {
  value = data.twilio_serverless_variables.variables
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the variables are associated with
- `environment_sid` - (Mandatory) The SID of the environment the variables are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `service_sid/environment_sid`
- `service_sid` - The SID of the service the variables are associated with
- `environment_sid` - The SID of the environment the variables are associated with
- `account_sid` - The account SID associated with the variables
- `variables` - A list of `variable` blocks as documented below

---

A `variable` block supports the following:

- `sid` - The SID of the environment variable
- `key` - The key of the environment variable
- `value` - The value of the environment variable
- `date_created` - The date in RFC3339 format that the environment variable was created
- `date_updated` - The date in RFC3339 format that the environment variable was updated
- `url` - The URL of the environment variable

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving environment variables
