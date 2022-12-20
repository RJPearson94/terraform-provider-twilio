---
page_title: "Twilio Serverless Service"
subcategory: "Serverless"
---

# twilio_serverless_service Resource

Manages a Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/service) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_serverless_service" "service" {
  unique_name   = "twilio-test"
  friendly_name = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `unique_name` - (Mandatory) The unique name of the service. The length of the string must be between `1` and `50` characters (inclusive)
- `friendly_name` - (Mandatory) The name of the service. The length of the string must be between `1` and `255` characters (inclusive)
- `include_credentials` - (Optional) Whether or not credentials are included in the service runtime. The default value is `true`
- `ui_editable` - (Optional) Whether or not the service is editable in the console. The default value is `false`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The account SID of the service is deployed into
- `unique_name` - The unique name of the service
- `friendly_name` - The name of the service
- `include_credentials` - Whether or not credentials are included in the service runtime
- `ui_editable` - Whether or not the service is editable in the console
- `domain_base` - The base name of the service
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the service
- `update` - (Defaults to 10 minutes) Used when updating the service
- `read` - (Defaults to 5 minutes) Used when retrieving the service
- `delete` - (Defaults to 10 minutes) Used when deleting the service

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_serverless_service.service /Services/ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
