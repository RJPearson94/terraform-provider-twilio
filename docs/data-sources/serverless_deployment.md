---
page_title: "Twilio Serverless Deployment"
subcategory: "Serverless"
---

# twilio_serverless_deployment Data Source

Use this data source to access information about an existing Serverless deployment. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/deployment) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_deployment" "deployment" {
  service_sid     = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  environment_sid = "ZEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid             = "ZDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "deployment" {
  value = data.twilio_serverless_deployment.deployment
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the deployment is associated with
- `environment_sid` - (Mandatory) The SID of the environment the deployment is associated with
- `sid` - (Mandatory) The SID of the deployment

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the deployment (Same as the `sid`)
- `sid` - The SID of the deployment (Same as the `id`)
- `account_sid` - The account SID associated with the deployment
- `service_sid` - The service SID associated with the deployment
- `environment_sid` - The environment SID associated with the deployment
- `build_sid` - The build SID to be deployed to the environment
- `date_created` - The date in RFC3339 format that the deployment was created
- `date_updated` - The date in RFC3339 format that the deployment was updated
- `url` - The URL of the deployment

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the deployment
