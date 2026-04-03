---
page_title: "twilio_account_sub_account Resource - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_sub_account Resource

Manages a Twilio sub-account. See the [API docs](https://www.twilio.com/docs/iam/api/account) for more information

~> Currently only sub-accounts can be created via the API. Parent accounts have to be created via the Twilio console

!> If the `friendly_name` is managed via Terraform and the `friendly_name` is removed from the configuration file. The old value will be retained on the next apply.

## Example Usage

```hcl
resource "twilio_account_sub_account" "sub_account" {
  friendly_name = "twilio-test"
}
```

## Schema

### Optional

- `friendly_name` (String) A human-readable label for the sub-account
- `status` (String) The status of the sub-account. Valid values are `active`, `suspended`, or `closed`. Defaults to `active`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `auth_token` (String, Sensitive) The authorization token for the sub-account. Sensitive -- will not be shown in logs or plans
- `date_created` (String) The date and time the sub-account was created, in RFC 3339 format
- `date_updated` (String) The date and time the sub-account was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `owner_account_sid` (String) The SID of the parent account that owns this sub-account
- `sid` (String) The unique SID assigned to this sub-account by Twilio
- `type` (String) The type of the account (e.g., `Trial` or `Full`)

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

An account can be imported using the `/Accounts/{sid}` format, e.g.

```shell
terraform import twilio_account_sub_account.sub_account /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
