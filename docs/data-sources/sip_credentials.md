---
page_title: "twilio_sip_credentials Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_credentials Data Source

Use this data source to access information about the credentials associated with an existing account and credential list. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-credential-resource) for more information

## Example Usage

```hcl
data "twilio_sip_credentials" "credentials" {
  account_sid         = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  credential_list_sid = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credentials" {
  value = data.twilio_sip_credentials.credentials
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns the SIP credentials
- `credential_list_sid` (String) The SID of the credential list to retrieve credentials from

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `credentials` (List of Object) A list of SIP credentials in the credential list (see [below for nested schema](#nestedatt--credentials))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--credentials"></a>
### Nested Schema for `credentials`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `sid` (String)
- `username` (String)
