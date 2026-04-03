---
page_title: "twilio_serverless_asset Data Source - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_asset Data Source

Use this data source to access information about an existing Serverless asset. See the [API docs](https://www.twilio.com/docs/proxy/api/phone-number) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_asset" "asset" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "asset" {
  value = data.twilio_serverless_asset.asset
}
```

## Schema

### Required

- `service_sid` (String) The SID of the Serverless service
- `sid` (String) The SID of the Serverless asset

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this asset
- `date_created` (String) The date and time the asset was created, in RFC 3339 format
- `date_updated` (String) The date and time the asset was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the asset
- `id` (String) The ID of this resource.
- `latest_version_sid` (String) The SID of the latest version of the asset
- `path` (String) The URL path at which the asset is accessible
- `url` (String) The absolute URL of the asset resource
- `visibility` (String) The access control for the asset

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
