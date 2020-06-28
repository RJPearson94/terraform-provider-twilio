---
page_title: "Twilio API Keys"
subcategory: "IAM"
---

# twilio_iam_api_key Resource

Manages a API Key for a Twilio Account

## Example Usage

```hcl
resource "twilio_iam_api_key" "api_key" {
  friendly_name = "Test API Key"
}
```

## Argument Reference

The following arguments are supported:

* `friendly_name` - (Optional) The name of the API Key

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the API Key (Same as the SID)
* `sid` - The SID of the API Key (Same as the ID)
* `friendly_name` - The name of the API Key
* `secret` - The API Key Secret
* `date_created` - The date in RFC3339 format that the API Key was created
* `date_updated` - The date in RFC3339 format that the API Key was updated
