---
page_title: "Twilio Serverless Asset"
subcategory: "Serverless"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The serverless service SID to associate the asset with. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the asset. The length of the string must be between `1` and `255` characters (inclusive)
- `content_file_name` - (Optional) The name of the file. Conflicts with `source`
- `content` - (Optional) The file contents as string. Conflicts with `source`
- `source` - (Optional) The relative path to the asset file. Conflicts with `content`
- `source_hash` - (Optional) A hash of the asset file to trigger deployments. Conflicts with `content`
- `content_type` - (Mandatory) The file MIME-type
- `path` - (Mandatory) The request URI path. The length of the string must be between `1` and `255` characters (inclusive)
- `visibility` - (Mandatory) The visibility of the asset. Valid values are `public` or `protected` or `private`

~> Either `source` or `content` need to be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the asset (Same as the `sid`)
- `sid` - The SID of the asset (Same as the `id`)
- `account_sid` - The account SID of the asset is deployed into
- `service_sid` - The service SID of the asset is managed under
- `friendly_name` - The name of the asset
- `content_file_name` - The name of the file
- `latest_version_sid` - The SID of the latest asset version
- `source` - The relative path to the asset file
- `source_hash` - A hash of the asset file to trigger deployments
- `content_type` - The file MIME-type
- `path` - The request URI path
- `visibility` - The visibility of the asset
- `date_created` - The date in RFC3339 format that the asset was created
- `date_updated` - The date in RFC3339 format that the asset was updated
- `url` - The URL of the asset

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the asset
- `update` - (Defaults to 10 minutes) Used when updating the asset
- `read` - (Defaults to 5 minutes) Used when retrieving the asset
- `delete` - (Defaults to 10 minutes) Used when deleting the asset

## Import

A asset can be imported using the `/Services/{serviceSid}/Assets/{sid}` format, e.g.

```shell
terraform import twilio_serverless_asset.asset /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Assets/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments `content`, `content_file_name`, `content_type` and `source_hash` cannot be imported, as the API doesn't return this data
