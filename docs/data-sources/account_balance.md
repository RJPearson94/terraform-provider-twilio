---
page_title: "Twilio Account Balance"
subcategory: "Account"
---

# twilio_account_balance Data Source

Use this data source to access balance information for an existing account.

!> This balance can only be retrieved for parent/ owner accounts

## Example Usage

```hcl
data "twilio_account_balance" "balance" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "balance" {
  value = data.twilio_account_balance.balance.balance
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The sid of the account

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the account (Same as the SID)
- `account_sid` - The SID of the account (Same as the ID)
- `balance` - The balance of the account
- `currency` - The currency of the account

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the account balance
