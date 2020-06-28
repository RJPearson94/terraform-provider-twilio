---
page_title: "Twilio Serverless Build"
subcategory: "Serverless"
---

# twilio_serverless_build Resource

Manages a Serverless build.
If polling is enabled then the create step will poll until the build status is either `completed` or `failed` or the max attempts threshold is reached.

!> This resource is in beta

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
  visibility        = "public"
}

resource "twilio_serverless_build" "build" {
  service_sid           = twilio_serverless_service.service.sid
  function_version_sids = [twilio_serverless_function_version.function_version.sid]
  dependencies = {
    "twilio" : "3.6.3"
  }

  polling {
    enabled = true
  }
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID associated with the build. Changing this forces a new resource to be created
- `asset_version_sids` - (Optional) List of Asset Version Sids to be included in the build. Changing this forces a new resource to be created
- `function_version_sids` - (Optional) List of Function Version Sids to be included in the build. Changing this forces a new resource to be created
- `dependencies` - (Optional) Map of dependencies to be included in the build. Changing this forces a new resource to be created
- `polling` - (Optional) A `polling` block as documented below.

---

A `polling` block supports the following:

- `enabled` - (Required) Enable or or disable polling of the build.
- `max_attempts` - (Optional) The maximum number of polling attempts. Default is 30
- `delay_in_ms` - (Optional) The time in milliseconds to wait between polling attempts.

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the build (Same as the SID)
- `sid` - The SID of the build (Same as the ID)
- `account_sid` - The Account SID associated with the build
- `service_sid` - The Service SID associated with the build
- `asset_version_sids` - List of Asset Version Sids to be included in the build
- `asset_versions` - List of `asset_version` blocks as documented below.
- `function_version_sids` - List of Function Version Sids to be included in the build
- `function_versions` - List of `function_version` blocks as documented below.
- `dependencies` - Map of dependencies to be included in the build
- `status` - The current status of the build job
- `date_created` - The date that the build was created
- `date_updated` - The date that the build was updated
- `url` - The url of the build

---

A `asset_version` block supports the following:

- `sid` - The SID of the asset version (Same as the ID)
- `account_sid` - The Account SID of the asset version is deployed into
- `service_sid` - The Service SID of the asset version is deployed into
- `asset_sid` - The Service SID of the asset version is managed under
- `date_created` - The date that the asset version was created
- `path` - The request uri path
- `visibility` - The visibility of the asset

---

A `function_version` block supports the following:

- `sid` - The SID of the function version (Same as the ID)
- `account_sid` - The Account SID of the function version is deployed into
- `service_sid` - The Service SID of the function version is deployed into
- `function_sid` - The Service SID of the function version is managed under
- `date_created` - The date that the function version was created
- `path` - The request uri path
- `visibility` - The visibility of the function
