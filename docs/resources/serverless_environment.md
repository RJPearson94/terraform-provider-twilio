---
page_title: "Twilio Serverless Resource"
subcategory: "Serverless"
---

# twilio_serverless_service Resource

Manages a Serverless service

!> This resource is in beta

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "test"
  friendly_name = "test"
}

resource "twilio_serverless_environment" "environment" {
  service_sid = twilio_serverless_service.service.sid
  unique_name = "test"
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the environment is managed under. Changing this forces a new resource to be created
- `unique_name` - (Mandatory) The unique name of the environment
- `domain_suffix` - (Optional) The domain suffix of the environment

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the environment (Same as the SID)
- `sid` - The SID of the environment (Same as the ID)
- `account_sid` - The Account SID of the environment is deployed into
- `service_sid` - The Service SID of the environment is managed under
- `build_sid` - The Build SID of the current build deployed to the environment
- `unique_name` - The unique name of the environment
- `domain_suffix` - The domain suffix of the environment
- `domain_name` - The domain name of the environment
- `date_created` - The date in RFC3339 format that the environment was created
- `date_updated` - The date in RFC3339 format that the environment was updated
- `url` - The url of the environment
