---
page_title: "Twilio Serverless Asset"
subcategory: "Serverless"
---

# twilio_serverless_variable Resource

Manages a Serverless Environment Variable

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "test"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid   = twilio_serverless_service.service.sid
  unique_name   = "test"
}

resource "twilio_serverless_variable" "variable" {
  service_sid     = twilio_serverless_service.service.sid
  environment_sid = twilio_serverless_environment.environment.sid
  key             = "test-key"
  value           = "test-value"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the environment variable is deployed into. Changing this forces a new resource to be created
- `environment_sid` - (Mandatory) The Environment SID of the environment variable is managed under. Changing this forces a new resource to be created
- `key` - (Mandatory) The key of the environment variable
- `value` - (Mandatory) The value of the environment variable

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the environment variable (Same as the SID)
- `sid` - The SID of the environment variable (Same as the ID)
- `account_sid` - The Account SID of the environment variable is deployed into
- `service_sid` - The Service SID of the environment variable is deployed into
- `environment_sid` - The Environment SID of the environment variable is managed under
- `key` - The key of the environment variable
- `value` - The value of the environment variable
- `date_created` - The date in RFC3339 format that the environment variable was created
- `date_updated` - The date in RFC3339 format that the environment variable was updated
- `url` - The url of the environment
