---
page_title: "Twilio Sub Account"
subcategory: "Account"
---

# twilio_account_sub_account Resource

Manages a Twilio sub account. See the [API docs](https://www.twilio.com/docs/iam/api/account) for more information

~> Currently only sub accounts can be created via the API. Parent accounts have to be created via the Twilio console

## Example Usage

```hcl
resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the account
- `status` - (Optional) The status of the account

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the account (Same as the SID)
- `sid` - The SID of the account (Same as the ID)
- `friendly_name` - The friendly name of the account
- `status` - The status of the account
- `owner_account_sid` - The SID of the parent/ owner account
- `type` - The type of account
- `auth_token` - The auth token for the account
- `date_created` - The date in RFC3339 format that the account was created
- `date_updated` - The date in RFC3339 format that the account was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the account
- `update` - (Defaults to 10 minutes) Used when updating the account
- `read` - (Defaults to 5 minutes) Used when retrieving the account
- `delete` - (Defaults to 10 minutes) Used when deleting the account

## Import

A account can be imported using the `/Accounts/{sid}` format, e.g.

```shell
terraform import twilio_account_sub_account.sub_account /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
