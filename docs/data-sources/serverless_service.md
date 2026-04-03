---
page_title: "twilio_serverless_service Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_service Data Source

Use this data source to access information about an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/service) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

### SID

```hcl
data "twilio_serverless_service" "service" {
  sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

### Unique Name

```hcl
data "twilio_serverless_service" "service" {
  unique_name = "UniqueName"
}
```

## Schema

### Optional

- `sid` (String) The SID of the Serverless service. Exactly one of `sid` or `unique_name` must be specified
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `unique_name` (String) The unique name of the Serverless service. Exactly one of `sid` or `unique_name` must be specified

### Read-Only

- `account_sid` (String) The SID of the account that owns this service
- `date_created` (String) The date and time the service was created, in RFC 3339 format
- `date_updated` (String) The date and time the service was last updated, in RFC 3339 format
- `domain_base` (String) The base domain name for this service, used to compose the URLs of the service's environments
- `friendly_name` (String) A human-readable label for the service
- `id` (String) The ID of this resource.
- `include_credentials` (Boolean) Whether account credentials are injected into function invocation contexts
- `ui_editable` (Boolean) Whether the service's properties and sub-resources can be edited in the Twilio Console
- `url` (String) The absolute URL of the service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
