---
page_title: "Twilio Serverless Environments"
subcategory: "Serverless"
---

# twilio_serverless_environments Data Source

Use this data source to access information about the environments associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/environment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_environments" "environments" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "environments" {
  value = data.twilio_serverless_environments.environments
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the environments are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the service the environments are associated with
- `account_sid` - The account SID associated with the environments
- `environments` - A list of `environment` blocks as documented below

---

A `environment` block supports the following:

- `sid` - The SID of the environment
- `build_sid` - The build SID of the current build deployed to the environment
- `unique_name` - The unique name of the environment
- `domain_suffix` - The domain suffix of the environment
- `domain_name` - The domain name of the environment
- `date_created` - The date in RFC3339 format that the environment was created
- `date_updated` - The date in RFC3339 format that the environment was updated
- `url` - The url of the environment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving environments
