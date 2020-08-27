---
page_title: "Twilio API Keys"
subcategory: "IAM"
---

# twilio_iam_api_key Resource

Manages a API Key for a Twilio Account

## Example Usage

```hcl
resource "twilio_iam_api_key" "api_key" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test API Key"
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The Account SID associated with the API Key
- `friendly_name` - (Optional) The name of the API Key

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the API Key (Same as the SID)
- `sid` - The SID of the API Key (Same as the ID)
- `account_sid` - The Account SID associated with the API Key
- `friendly_name` - The name of the API Key
- `secret` - The API Key Secret
- `date_created` - The date in RFC3339 format that the API Key was created
- `date_updated` - The date in RFC3339 format that the API Key was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the API Key
- `update` - (Defaults to 10 minutes) Used when updating the API Key
- `read` - (Defaults to 5 minutes) Used when retrieving the API Key
- `delete` - (Defaults to 10 minutes) Used when deleting the API Key

## Import

Not supported
