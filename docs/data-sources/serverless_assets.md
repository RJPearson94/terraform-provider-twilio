---
page_title: "Twilio Serverless Assets"
subcategory: "Serverless"
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

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the assets are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the service the assets are associated with
- `account_sid` - The account SID associated with the assets
- `assets` - A list of `asset` blocks as documented below

---

A `asset` block supports the following:

- `sid` - The SID of the asset
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

- `read` - (Defaults to 10 minutes) Used when retrieving assets
