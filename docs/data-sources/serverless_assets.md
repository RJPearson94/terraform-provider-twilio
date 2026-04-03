---
page_title: "twilio_serverless_assets Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_assets Data Source

Use this data source to access information about the assets associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_assets" "assets" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "assets" {
  value = data.twilio_serverless_assets.assets
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these assets
- `assets` (List of Object) A list of assets belonging to the Serverless service (see [below for nested schema](#nestedatt--assets))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--assets"></a>
### Nested Schema for `assets`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `latest_version_sid` (String)
- `path` (String)
- `sid` (String)
- `url` (String)
- `visibility` (String)
