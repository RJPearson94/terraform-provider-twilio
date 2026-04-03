---
page_title: "twilio_sip_domain Resource - twilio"
subcategory: "SIP"
description: |-
  
---

# twilio_sip_domain Resource

Manages a SIP domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-domain-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_name = "test.sip.twilio.com"
}
```

## Schema

### Required

- `account_sid` (String) The SID of the account that owns this SIP domain. Changing this forces a new resource
- `domain_name` (String) The fully qualified domain name for the SIP domain, ending with `.sip.twilio.com`

### Optional

- `byoc_trunk_sid` (String) The SID of the BYOC trunk to associate with this SIP domain
- `emergency` (Block List, Max: 1) A block to configure emergency calling settings for the SIP domain (see [below for nested schema](#nestedblock--emergency))
- `friendly_name` (String) A human-readable label for the SIP domain
- `secure` (Boolean) Whether secure SIP (SIPS) is enabled for the domain. Defaults to `false`
- `sip_registration` (Boolean) Whether SIP registration is allowed for the domain. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `voice` (Block List, Max: 1) A block to configure voice settings for the SIP domain (see [below for nested schema](#nestedblock--voice))

### Read-Only

- `auth_type` (String) The authentication type configured for the SIP domain
- `date_created` (String) The date and time the SIP domain was created, in RFC 3339 format
- `date_updated` (String) The date and time the SIP domain was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this SIP domain by Twilio

<a id="nestedblock--emergency"></a>
### Nested Schema for `emergency`

Optional:

- `caller_sid` (String) The SID of the phone number to use as the emergency caller ID
- `calling_enabled` (Boolean) Whether emergency calling is enabled for this SIP domain. Defaults to `false`


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedblock--voice"></a>
### Nested Schema for `voice`

Optional:

- `fallback_method` (String) The HTTP method used to call the voice fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `fallback_url` (String) The URL to call when an error occurs retrieving or executing the TwiML for voice calls
- `method` (String) The HTTP method used to call the voice URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_method` (String) The HTTP method used to call the voice status callback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_url` (String) The URL to call for voice status callback events
- `url` (String) The URL to call when the SIP domain receives an incoming voice call

## Import

An domain can be imported using the `Accounts/{accountSid}/Domains/{sid}` format, e.g.

```shell
terraform import twilio_sip_domain.domain /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Domains/DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
