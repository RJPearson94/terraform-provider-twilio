---
page_title: "twilio_verify_service Data Source - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service Data Source

Use this data source to access information about an existing Verify service. See the [API docs](https://www.twilio.com/docs/verify/api/service) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

## Example Usage

```hcl
data "twilio_verify_service" "service" {
  sid = "VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "service" {
  value = data.twilio_verify_service.service
}
```

## Schema

### Required

- `sid` (String) The SID of the Verify service to fetch

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Verify service
- `code_length` (Number) The length of the verification code
- `custom_code_enabled` (Boolean) Whether sending verifications with a custom code is enabled
- `date_created` (String) The date and time the Verify service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Verify service was last updated, in RFC 3339 format
- `default_template_sid` (String) The SID of the default verification template used by this Verify service
- `do_not_share_warning_enabled` (Boolean) Whether a warning not to share the verification code is included in the SMS body
- `dtmf_input_required` (Boolean) Whether the user must press a key to deliver the verification code via phone call
- `friendly_name` (String) A human-readable label for the Verify service
- `id` (String) The ID of this resource.
- `lookup_enabled` (Boolean) Whether a phone number lookup is performed with each verification
- `mailer_sid` (String) The SID of the mailer service associated with this Verify service
- `psd2_enabled` (Boolean) Whether PSD2 transaction parameters are passed when starting a verification
- `push` (List of Object) Push notification configuration for the Verify service (see [below for nested schema](#nestedatt--push))
- `skip_sms_to_landlines` (Boolean) Whether SMS verifications to landlines are skipped
- `totp` (List of Object) Time-based one-time password (TOTP) configuration for the Verify service (see [below for nested schema](#nestedatt--totp))
- `tts_name` (String) The name of the text-to-speech voice used for phone call verifications
- `url` (String) The absolute URL of the Verify service resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--push"></a>
### Nested Schema for `push`

Read-Only:

- `apn_credential_sid` (String)
- `fcm_credential_sid` (String)


<a id="nestedatt--totp"></a>
### Nested Schema for `totp`

Read-Only:

- `code_length` (Number)
- `issuer` (String)
- `skew` (Number)
- `time_step` (Number)
