---
page_title: "Twilio Account Details"
subcategory: "Account"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Optional) The SID of the account

~> If an account SID is not supplied the account SID configured on the provider is used instead

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the account (Same as the `sid`)
- `sid` - The SID of the account (Same as the `id`)
- `friendly_name` - The friendly name of the account
- `status` - The status of the account
- `owner_account_sid` - The SID of the parent/ owner account
- `type` - The type of account
- `auth_token` - The auth token for the account
- `date_created` - The date in RFC3339 format that the account was created
- `date_updated` - The date in RFC3339 format that the account was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the account details
