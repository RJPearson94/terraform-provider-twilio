---
page_title: "twilio_sip_domain Data Source - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_domain Data Source

Use this data source to access information about an existing domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-domain-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "domain" {
  value = data.twilio_sip_domain.domain
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP domain
- `sid` (String) The SID of the SIP domain to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `auth_type` (String) The authentication type configured for the SIP domain
- `byoc_trunk_sid` (String) The SID of the BYOC trunk associated with this SIP domain
- `date_created` (String) The date and time the SIP domain was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP domain was last updated, in RFC 3339 format
- `domain_name` (String) The fully qualified domain name for the SIP domain
- `emergency` (List of Object) The emergency calling settings for the SIP domain (see [below for nested schema](#nestedatt--emergency))
- `friendly_name` (String) A human-readable label for the SIP domain
- `id` (String) The ID of this resource.
- `secure` (Boolean) Whether secure SIP (SIPS) is enabled for the domain
- `sip_registration` (Boolean) Whether SIP registration is allowed for the domain
- `voice` (List of Object) The voice settings for the SIP domain (see [below for nested schema](#nestedatt--voice))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--emergency"></a>
### Nested Schema for `emergency`

Read-Only:

- `caller_sid` (String)
- `calling_enabled` (Boolean)


<a id="nestedatt--voice"></a>
### Nested Schema for `voice`

Read-Only:

- `fallback_method` (String)
- `fallback_url` (String)
- `method` (String)
- `status_callback_method` (String)
- `status_callback_url` (String)
- `url` (String)
