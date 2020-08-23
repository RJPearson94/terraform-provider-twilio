---
page_title: "Twilio Account Details"
subcategory: "Account"
---

# twilio_account_details Data Source

Use this data source to access information about an existing account

## Example Usage

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

- `sid` - (Mandatory) The sid of the account

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
