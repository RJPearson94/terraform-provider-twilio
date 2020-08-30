---
page_title: "Twilio Serverless Asset"
subcategory: "Serverless"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the asset is associated with
- `sid` - (Mandatory) The SID of the asset

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the asset (Same as the SID)
- `sid` - The SID of the asset (Same as the ID)
- `account_sid` - The account SID of the asset is deployed into
- `service_sid` - The service SID of the asset is managed under
- `friendly_name` - The name of the asset
- `content_file_name` - The name of the file
- `latest_version_sid` - The SID of the latest asset version
- `source` - The relative path to the asset file
- `source_hash` - A hash of the asset file to trigger deployments
- `content_type` - The file MIME type
- `path` - The request uri path
- `visibility` - The visibility of the asset
- `date_created` - The date in RFC3339 format that the asset was created
- `date_updated` - The date in RFC3339 format that the asset was updated
- `url` - The url of the asset

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the asset
