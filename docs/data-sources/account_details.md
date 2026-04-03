---
page_title: "twilio_account_details Data Source - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_details Data Source

Use this data source to access information about an existing account. See the [API docs](https://www.twilio.com/docs/iam/api/account) for more information

## Example Usage

## Use the provider Account SID

```hcl
data "twilio_account_details" "account" {}

output "account_sid" {
  value = data.twilio_account_details.account.account_sid
}
```

## Specify Account SID

```hcl
data "twilio_account_details" "account" {
  sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "friendly_name" {
  value = data.twilio_account_details.account.friendly_name
}
```

## Schema

### Optional

- `sid` (String) The SID of the account to look up. If not specified, the provider's account SID is used
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `auth_token` (String, Sensitive) The authorization token for the account. Sensitive -- will not be shown in logs or plans
- `date_created` (String) The date and time the account was created, in RFC 3339 format
- `date_updated` (String) The date and time the account was last updated, in RFC 3339 format
- `friendly_name` (String) A human-readable label for the account
- `id` (String) The ID of this resource.
- `owner_account_sid` (String) The SID of the parent account that owns this account
- `status` (String) The status of the account. Valid values are `active`, `suspended`, or `closed`
- `type` (String) The type of the account (e.g., `Trial` or `Full`)

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
