---
page_title: "Twilio Serverless Deployments"
subcategory: "Serverless"
---

# twilio_serverless_deployments Data Source

Use this data source to access information about the deployments associated with an existing Serverless service and environment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_deployments" "deployments" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "deployments" {
  value = data.twilio_serverless_deployments.deployments
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the deployments are associated with
- `environment_sid` - (Mandatory) The SID of the environment the deployments are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `service_sid/environment_sid`
- `service_sid` - The SID of the service the deployments are associated with
- `environment_sid` - The SID of the environment the deployments are associated with
- `account_sid` - The account SID associated with the deployments
- `deployments` - A list of `deployment` blocks as documented below

---

A `deployment` block supports the following:

- `sid` - The SID of the deployment
- `build_sid` - The build SID to be deployed to the environment
- `date_created` - The date in RFC3339 format that the deployment was created
- `date_updated` - The date in RFC3339 format that the deployment was updated
- `url` - The URL of the deployment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving deployments
