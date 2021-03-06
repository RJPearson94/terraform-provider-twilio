---
page_title: "Twilio Serverless Function"
subcategory: "Serverless"
---

# twilio_serverless_function Data Source

Use this data source to access information about an existing Serverless function. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/function) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_function" "function" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "ZBXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "function" {
  value = data.twilio_serverless_function.function
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the function is associated with
- `sid` - (Mandatory) The SID of the function

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function (Same as the `sid`)
- `sid` - The SID of the function (Same as the `id`)
- `account_sid` - The account SID of the function is deployed into
- `service_sid` - The service SID of the function is managed under
- `friendly_name` - The name of the function
- `content_file_name` - The name of the file
- `latest_version_sid` - The SID of the latest function version
- `source` - The relative path to the function file
- `source_hash` - A hash of the function file to trigger deployments
- `content_type` - The file MIME-type
- `path` - The request URI path
- `visibility` - The visibility of the function
- `date_created` - The date in RFC3339 format that the function was created
- `date_updated` - The date in RFC3339 format that the function was updated
- `url` - The URL of the function

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the function
