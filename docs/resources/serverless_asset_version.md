# twilio_serverless_asset_version

Manages a Serverless asset version

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_asset" "asset" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}

resource "twilio_serverless_asset_version" "asset_version" {
  service_sid       = twilio_serverless_service.service.sid
  asset_sid         = twilio_serverless_asset.asset.sid
  content           = "{}"
  content_type      = "application/json"
  content_file_name = "test.json"
  path              = "/test-asset"
  visibility        = "private"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the asset version is deployed into. Changing this forces a new resource to be created
- `asset_sid` - (Mandatory) The Service SID of the asset version is managed under. Changing this forces a new resource to be created
- `content_file_name` - (Optional) The name of the file. Conflicts with source
- `content` - (Optional) The file contents as string.Conflicts with source
- `source` - (Optional) The relative path to the asset file. Conflicts with content
- `source_hash` - (Optional) A hash of the asset file to trigger deployments. Conflicts with content
- `content_type` - (Mandatory) The file MINE type
- `path` - (Mandatory) The request uri path
- `visibility` - (Mandatory) The visibility of the asset. Options are `public` or `protected` or `private`

*NOTE:* Either source or content need to be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the asset version (Same as the SID)
- `sid` - The SID of the asset version (Same as the ID)
- `account_sid` - The Account SID of the asset version is deployed into
- `service_sid` - The Service SID of the asset version is deployed into
- `asset_sid` - The Service SID of the asset version is managed under
- `content_file_name` - The name of the file
- `source` - The relative path to the asset file
- `source_hash` - A hash of the asset file to trigger deployments
- `content_type` - The file MINE type
- `path` - The request uri path
- `visibility` - The visibility of the asset
- `date_created` - The date that the asset version was created
- `date_updated` - The date that the asset version was updated
- `url` - The url of the asset version
