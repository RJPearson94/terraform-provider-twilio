---
page_title: "Twilio Serverless Function"
subcategory: "Serverless"
---

# twilio_serverless_function Resource

Manages a versioned Serverless function. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/function) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"

  content           = <<EOF
exports.handler = function (context, event, callback) {
  callback(null, "Hello World");
};
EOF
  content_type      = "application/javascript"
  content_file_name = "helloWorld.js"
  path              = "/test-function"
  visibility        = "private"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The service SID of the function is managed under. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the function
- `content_file_name` - (Optional) The name of the file
- `content` - (Optional) The file contents as string
- `source` - (Optional) The relative path to the function file
- `source_hash` - (Optional) A hash of the function file to trigger deployments
- `content_type` - (Mandatory) The file MIME type
- `path` - (Mandatory) The request uri path
- `visibility` - (Mandatory) The visibility of the function. Options are `public` or `protected` or `private`

**NOTE:** Either source or content need to be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function (Same as the SID)
- `sid` - The SID of the function (Same as the ID)
- `account_sid` - The account SID of the function is deployed into
- `service_sid` - The service SID of the function is managed under
- `friendly_name` - The name of the function
- `content_file_name` - The name of the file
- `latest_version_sid` - The SID of the latest function version
- `source` - The relative path to the function file
- `source_hash` - A hash of the function file to trigger deployments
- `content_type` - The file MIME type
- `path` - The request URI path
- `visibility` - The visibility of the function
- `date_created` - The date in RFC3339 format that the function was created
- `date_updated` - The date in RFC3339 format that the function was updated
- `url` - The URL of the function

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the function
- `update` - (Defaults to 10 minutes) Used when updating the function
- `read` - (Defaults to 5 minutes) Used when retrieving the function
- `delete` - (Defaults to 10 minutes) Used when deleting the function

## Import

A function can be imported using the `/Services/{serviceSid}/Functions/{sid}` format, e.g.

```shell
terraform import twilio_serverless_function.function /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Functions/ZHXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> The following arguments "content_file_name", "content_type" and "source_hash" cannot be imported, as the API doesn't return this data
