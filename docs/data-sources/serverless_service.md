---
page_title: "Twilio Serverless Service"
subcategory: "Serverless"
---

# twilio_serverless_service Data Source

Use this data source to access information about an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/service) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_build" "build" {
  sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The account SID of the service is deployed into
- `unique_name` - The unique name of the service
- `friendly_name` - The name of the service
- `include_credentials` - Whether or not credentials are included in the service runtime
- `ui_editable` - Whether or not the service is editable in the console
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the service
