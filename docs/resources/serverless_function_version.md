# twilio_serverless_function_version

Manages a Serverless function version

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}

resource "twilio_serverless_function" "function" {
  service_sid   = twilio_serverless_service.service.sid
  friendly_name = "test"
}

resource "twilio_serverless_function_version" "function_version" {
  service_sid       = twilio_serverless_service.service.sid
  function_sid      = twilio_serverless_function.function.sid
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

- `service_sid` - (Mandatory) The Service SID of the function version is deployed into. Changing this forces a new resource to be created
- `function_sid` - (Mandatory) The Service SID of the function version is managed under. Changing this forces a new resource to be created
- `content_file_name` - (Optional) The name of the file. Conflicts with source. Changing this forces a new resource to be created
- `content` - (Optional) The file contents as string. Conflicts with source. Changing this forces a new resource to be created
- `source` - (Optional) The relative path to the function file. Conflicts with content. Changing this forces a new resource to be created
- `source_hash` - (Optional) A hash of the function file to trigger deployments. Conflicts with content. Changing this forces a new resource to be created
- `content_type` - (Mandatory) The file MIME type. Changing this forces a new resource to be created
- `path` - (Mandatory) The request uri path. Changing this forces a new resource to be created
- `visibility` - (Mandatory) The visibility of the function. Options are `public` or `protected` or `private`. Changing this forces a new resource to be created

*NOTE:* Either source or content need to be specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function version (Same as the SID)
- `sid` - The SID of the function version (Same as the ID)
- `account_sid` - The Account SID of the function version is deployed into
- `service_sid` - The Service SID of the function version is deployed into
- `function_sid` - The Service SID of the function version is managed under
- `content_file_name` - The name of the file
- `source` - The relative path to the function file
- `source_hash` - A hash of the function file to trigger deployments
- `content_type` - The file MIME type
- `path` - The request uri path
- `visibility` - The visibility of the function
- `date_created` - The date that the function version was created
- `date_updated` - The date that the function version was updated
- `url` - The url of the function version
