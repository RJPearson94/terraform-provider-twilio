---
page_title: "twilio_sip_trunking_credential_lists Data Source - twilio"
subcategory: "SIP Trunking"
description: |-
  
---

# twilio_sip_trunking_credential_lists Data Source

Use this data source to access information about the credential lists associated with an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_credential_lists" "credential_lists" {
  trunk_sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_lists" {
  value = data.twilio_sip_trunking_credential_lists.credential_lists
}
```

## Schema

### Required

- `trunk_sid` (String) The SID of the SIP trunk to retrieve credential lists for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns the SIP trunk credential lists
- `credential_lists` (List of Object) A list of credential lists associated with the SIP trunk (see [below for nested schema](#nestedatt--credential_lists))
- `id` (String) The ID of this resource.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--credential_lists"></a>
### Nested Schema for `credential_lists`

Read-Only:

- `date_created` (String)
- `date_updated` (String)
- `friendly_name` (String)
- `sid` (String)
- `url` (String)
