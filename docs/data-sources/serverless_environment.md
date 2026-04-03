---
page_title: "twilio_serverless_environment Data Source - twilio"
subcategory: "Serverless"
description: |-
  
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

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless environment

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this environment
- `build_sid` (String) The SID of the build currently deployed to this environment
- `date_created` (String) The date and time the environment was created, in RFC 3339 format
- `date_updated` (String) The date and time the environment was last updated, in RFC 3339 format
- `domain_name` (String) The fully qualified domain name for this environment
- `domain_suffix` (String) A URL-friendly suffix appended to the environment's domain name
- `id` (String) The ID of this resource.
- `unique_name` (String) A unique, developer-assigned name for the environment
- `url` (String) The absolute URL of the environment resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
