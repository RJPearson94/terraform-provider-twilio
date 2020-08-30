---
page_title: "Twilio Serverless Environment"
subcategory: "Serverless"
---

# twilio_serverless_environment Data Source

Use this data source to access information about an existing Serverless environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/environment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_environment" "environment" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "environment" {
  value = data.twilio_serverless_environment.environment
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the environment is associated with
- `sid` - (Mandatory) The SID of the environment

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
- `url` - The url of the environment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the environment
