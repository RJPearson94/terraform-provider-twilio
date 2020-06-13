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
  service_sid  = twilio_serverless_service.service.sid
  function_sid = twilio_serverless_function.function.sid
  content_body = "ZXhwb3J0cy5oYW5kbGVyID0gZnVuY3Rpb24gKGNvbnRleHQsIGV2ZW50LCBjYWxsYmFjaykgewogIGNhbGxiYWNrKG51bGwsICJIZWxsbyBXb3JsZCIpOwp9Owo="
  content_type = "application/javascript"
  file_name    = "helloWorld.js"
  path         = "/test-function"
  visibility   = "private"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the function version is deployed into. Changing this forces a new resource to be created
- `function_sid` - (Mandatory) The Service SID of the function version is managed under. Changing this forces a new resource to be created
- `file_name` - (Mandatory) The name of the file
- `content_body` - (Mandatory) Base64 encoded file contents
- `content_type` - (Mandatory) The file MINE type
- `path` - (Mandatory) The request uri path
- `visibility` - (Mandatory) The visibility of the function. Options are `public` or `protected` or `private`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function version (Same as the SID)
- `sid` - The SID of the function version (Same as the ID)
- `account_sid` - The Account SID of the function version is deployed into
- `service_sid` - The Service SID of the function version is deployed into
- `function_sid` - The Service SID of the function version is managed under
- `file_name` - The name of the file
- `content_body` - Base64 encoded file contents
- `content_type` - The file MINE type
- `path` - The request uri path
- `visibility` - The visibility of the function
- `date_created` - The date that the function version was created
- `date_updated` - The date that the function version was updated
- `url` - The url of the function version
