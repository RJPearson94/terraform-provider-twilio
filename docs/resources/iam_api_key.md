---
page_title: "twilio_iam_api_key Resource - twilio"
subcategory: "IAM"
description: |-
  
---

# twilio_iam_api_key Resource

Manages an API Key for a Twilio Account. See the [API docs](https://www.twilio.com/docs/iam/keys/api-key-resource) for more information

!> Only Standard API Keys can be created via the API. If you require a Master API Key then you will need to create this manually in the Twilio console

## Example Usage

```hcl
resource "twilio_iam_api_key" "api_key" {
  account_sid   = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  friendly_name = "Test API Key"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this API key. Changing this forces a new resource

### Optional

- `friendly_name` (String) A human-readable label for the API key. Must be between 0 and 64 characters
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the API key was created, in RFC 3339 format
- `date_updated` (String) The date and time the API key was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `secret` (String, Sensitive) The secret for the API key, used for authentication. Sensitive -- will not be shown in logs or plans. Only available on initial creation
- `sid` (String) The unique SID assigned to this API key by Twilio

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

Not supported
