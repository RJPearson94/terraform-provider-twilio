---
page_title: "twilio_serverless_deployment Data Source - twilio"
subcategory: "Serverless"
description: |-
  
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

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment
- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless deployment

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this deployment
- `build_sid` (String) The SID of the build deployed in this deployment
- `date_created` (String) The date and time the deployment was created, in RFC 3339 format
- `date_updated` (String) The date and time the deployment was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the deployment resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
