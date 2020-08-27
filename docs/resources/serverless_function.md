---
page_title: "Twilio Serverless Function"
subcategory: "Serverless"
---

# twilio_serverless_function Resource

Manages a Serverless function

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
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The Service SID of the function is managed under. Changing this forces a new resource to be created
- `friendly_name` - (Mandatory) The name of the function

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the function (Same as the SID)
- `sid` - The SID of the function (Same as the ID)
- `account_sid` - The Account SID of the function is deployed into
- `service_sid` - The Service SID of the function is managed under
- `friendly_name` - The name of the function
- `date_created` - The date in RFC3339 format that the function was created
- `date_updated` - The date in RFC3339 format that the function was updated
- `url` - The url of the function

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
