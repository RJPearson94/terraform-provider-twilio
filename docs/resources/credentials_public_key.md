---
page_title: "twilio_credentials_public_key Resource - twilio"
subcategory: "Credentials"
description: |-
  
---

# twilio_credentials_public_key Resource

Manages a public key resource. This resource allows you to upload a public key for various Twilio services to use

!> If the `account_sid` is managed via Terraform and the `account_sid` is removed from the configuration file. The old value will be retained on the next apply.

## Example Usage

```hcl
resource "twilio_credentials_public_key" "public_key" {
  friendly_name = "Test Public Key Resource"
  public_key    = "-----BEGIN PUBLIC KEY-----....-----END PUBLIC KEY-----"
}
```

## Schema

### Required

- `public_key` (String) The PEM-encoded public key content to store as a credential. Changing this forces a new resource

### Optional

- `account_sid` (String) The SID of the account that owns this public key credential. Changing this forces a new resource
- `friendly_name` (String) A human-readable label for the public key credential
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `date_created` (String) The date and time the public key credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the public key credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this public key credential by Twilio
- `url` (String) The absolute URL of the public key credential resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)
