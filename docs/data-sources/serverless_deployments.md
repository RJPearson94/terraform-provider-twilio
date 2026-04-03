---
page_title: "twilio_serverless_deployments Data Source - twilio"
subcategory: "Serverless"
description: |-
  
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

## Schema

### Required

- `environment_sid` (String) The SID of the Serverless environment
- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these deployments
- `deployments` (List of Object) A list of deployments belonging to the Serverless environment (see [below for nested schema](#nestedatt--deployments))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--deployments"></a>
### Nested Schema for `deployments`

Read-Only:

- `build_sid` (String)
- `date_created` (String)
- `date_updated` (String)
- `sid` (String)
- `url` (String)
