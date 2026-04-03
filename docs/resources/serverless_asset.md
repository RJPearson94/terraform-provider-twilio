---
page_title: "twilio_serverless_asset Resource - twilio"
subcategory: "Serverless"
description: |-
  
---

# twilio_serverless_asset Resource

Manages a versioned Serverless asset. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/asset) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid       = twilio_serverless_service.service.sid
  friendly_name     = "test"
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}
```

## Schema

### Required

- `content_type` (String) The MIME type of the asset content (e.g. `text/html`)
- `friendly_name` (String) A human-readable label for the asset
- `path` (String) The URL path at which the asset will be accessible
- `service_sid` (String) The SID of the Serverless service. Changing this forces a new resource
- `visibility` (String) The access control for the asset. Valid values are `public`, `protected`, `private`

### Optional

- `content` (String) The inline content of the asset. Conflicts with `source`
- `content_file_name` (String) The file name to use when uploading inline content. Conflicts with `source`
- `source` (String) The path to the file containing the asset content. Conflicts with `content`
- `source_hash` (String) A hash of the source file, used to detect changes. Conflicts with `content`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this asset
- `date_created` (String) The date and time the asset was created, in RFC 3339 format
- `date_updated` (String) The date and time the asset was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `latest_version_sid` (String) The SID of the latest version of the asset
- `sid` (String) The unique SID assigned to this asset by Twilio
- `url` (String) The absolute URL of the asset resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A asset can be imported using the `/Services/{serviceSid}/Assets/{sid}` format, e.g.

```shell
terraform import twilio_serverless_asset.asset /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments `content`, `content_file_name`, `content_type` and `source_hash` cannot be imported, as the API doesn't return this data
