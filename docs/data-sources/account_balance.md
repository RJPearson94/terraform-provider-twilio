---
page_title: "twilio_account_balance Data Source - twilio"
subcategory: "Account"
description: |-
  
---

# twilio_account_balance Data Source

Use this data source to access balance information for an existing account. See the [API docs](https://www.twilio.com/docs/iam/api/account) for more information

~> This balance can only be retrieved for parent/ owner accounts

## Example Usage

```hcl
data "twilio_account_balance" "balance" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "balance" {
  value = data.twilio_account_balance.balance.balance
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account to retrieve the balance for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `balance` (String) The current balance of the account
- `currency` (String) The currency unit of the balance (e.g., USD)
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
