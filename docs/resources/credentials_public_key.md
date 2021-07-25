---
page_title: "Twilio Public Key"
subcategory: "Credentials"
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

## Argument Reference

The following arguments are supported:

- `account_sid` - (Optional) The SID of a sub account to associate the public key resource with. Changing this forces a new resource to be created
- `public_key` - (Mandatory) The public key to associate with the public key resource. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the public key resource

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the public key resource (Same as the `sid`)
- `sid` - The SID of the public key resource (Same as the `id`)
- `account_sid` - The account SID associated with the public key resource
- `friendly_name` - The friendly name of the public key resource
- `date_created` - The date in RFC3339 format that the public key resource was created
- `date_updated` - The date in RFC3339 format that the public key resource was updated
- `url` - The URL of the public key resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the public key resource
- `update` - (Defaults to 10 minutes) Used when updating the public key resource
- `read` - (Defaults to 5 minutes) Used when retrieving the public key resource
- `delete` - (Defaults to 10 minutes) Used when deleting the public key resource
