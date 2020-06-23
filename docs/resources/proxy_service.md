# twilio_proxy_service

Manages a Proxy service

## Example Usage

```hcl
resource "twilio_proxy_service" "service" {
  unique_name = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `unique_name` - (Mandatory) The unique name of the service
- `chat_instance_sid` - (Optional) The Chat Instance SID of the service
- `chat_service_sid` - (Optional) The Chat Service SID of the service
- `default_ttl` - (Optional) The default ttl of the service
- `geo_match_level` - (Optional) Where the proxy number and participant must be relatively located
- `number_selection_behavior` - (Optional) How the proxy service selects proxy numbers
- `callback_url` - v The callback url for the service
- `intercept_callback_url` - (Optional) The intercept callback url for the service
- `out_of_session_callback_url` - (Optional) The out of session callback url for the service

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the SID)
- `sid` - The SID of the service (Same as the ID)
- `account_sid` - The Account SID of the service is deployed into
- `chat_instance_sid` - The Chat Instance SID of the service
- `chat_service_sid` - The Chat Service SID of the service
- `unique_name` - The unique name of the service
- `default_ttl` - The default ttl of the service
- `geo_match_level` - Where the proxy number and participant must be relatively located
- `number_selection_behavior` - How the proxy service selects proxy numbers
- `callback_url` - The callback url for the service
- `intercept_callback_url` - The intercept callback url for the service
- `out_of_session_callback_url` - The out of session callback url for the service
- `date_created` - The date that the service was created
- `date_updated` - The date that the service was updated
- `url` - The url of the service
