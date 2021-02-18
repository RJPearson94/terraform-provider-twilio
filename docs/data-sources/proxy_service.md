---
page_title: "Twilio Proxy Service"
subcategory: "Proxy"
---

# twilio_proxy_service Data Source

Use this data source to access information about an existing Proxy service. See the [API docs](https://www.twilio.com/docs/proxy/api/service) for more information

For more information on Proxy, see the product [page](https://www.twilio.com/docs/proxy)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_proxy_service" "service" {
  sid = "KSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_proxy_service.service
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
- `chat_instance_sid` - The chat instance SID of the service
- `chat_service_sid` - The chat service SID of the service
- `unique_name` - The unique name of the service
- `default_ttl` - The default TTL of the service
- `geo_match_level` - Where the proxy number and participant must be relatively located
- `number_selection_behavior` - How the proxy service selects proxy numbers
- `callback_url` - The callback URL for the service
- `intercept_callback_url` - The intercept callback URL for the service
- `out_of_session_callback_url` - The out of session callback URL for the service
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the service
