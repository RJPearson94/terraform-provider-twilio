---
page_title: "twilio_verify_service Resource - twilio"
subcategory: "Verify"
description: |-
  
---

# twilio_verify_service Resource

Manages a Verify service. See the [API docs](https://www.twilio.com/docs/verify/api/service) for more information

For more information on Verify, see the product [page](https://www.twilio.com/verify)

!> If the `totp issuer` is managed via Terraform and the `issuer` is removed from the configuration file. The old value will be retained on the next apply.

## Example Usage

### Basic

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
}
```

### With TOTP

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
  totp {
    issuer      = "Test"
    time_step   = 60
    code_length = 4
    skew        = 2
  }
}
```

### With Mailer Config (Twilio Sendgrid Integration)

```hcl
resource "twilio_verify_service" "service" {
  friendly_name = "Test Verify Service"
  mailer_sid    = "MDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

~> The Sendgrid mailer is not currently supported by the provider

### With Default Template

```hcl
resource "twilio_verify_service" "service" {
  friendly_name        = "Test Verify Service"
  default_template_sid = "HJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the Verify service. Must be between 1 and 30 characters

### Optional

- `code_length` (Number) The length of the verification code to generate. Must be between 4 and 10. Defaults to `6`
- `custom_code_enabled` (Boolean) Whether to allow sending verifications with a custom code instead of a randomly generated one. Defaults to `false`
- `default_template_sid` (String) The SID of the default verification template to use for this Verify service
- `do_not_share_warning_enabled` (Boolean) Whether to include a warning not to share the verification code in the SMS body. Defaults to `false`
- `dtmf_input_required` (Boolean) Whether to require the user to press a key to deliver the verification code via phone call. Defaults to `true`
- `lookup_enabled` (Boolean) Whether to perform a phone number lookup with each verification. Defaults to `false`
- `mailer_sid` (String) The SID of the mailer service to associate with this Verify service for email verifications
- `psd2_enabled` (Boolean) Whether to pass PSD2 transaction parameters when starting a verification. Defaults to `false`
- `push` (Block List, Max: 1) Push notification configuration for the Verify service (see [below for nested schema](#nestedblock--push))
- `skip_sms_to_landlines` (Boolean) Whether to skip sending SMS verifications to landlines. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `totp` (Block List, Max: 1) Time-based one-time password (TOTP) configuration for the Verify service (see [below for nested schema](#nestedblock--totp))
- `tts_name` (String) The name of the text-to-speech voice to use for phone call verifications

### Read-Only

- `account_sid` (String) The SID of the account that owns this Verify service
- `date_created` (String) The date and time the Verify service was created, in RFC 3339 format
- `date_updated` (String) The date and time the Verify service was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this Verify service by Twilio
- `url` (String) The absolute URL of the Verify service resource

<a id="nestedblock--push"></a>
### Nested Schema for `push`

Optional:

- `apn_credential_sid` (String) The SID of the Apple Push Notification Service (APN) credential to use for push notifications
- `fcm_credential_sid` (String) The SID of the Firebase Cloud Messaging (FCM) credential to use for push notifications


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedblock--totp"></a>
### Nested Schema for `totp`

Optional:

- `code_length` (Number) The number of digits in the generated TOTP code. Must be between 3 and 8. Defaults to `6`
- `issuer` (String) The name that appears in the user's authenticator app as the issuer of the TOTP code
- `skew` (Number) The number of past and future time steps to allow during TOTP code validation. Must be between 0 and 2. Defaults to `1`
- `time_step` (Number) The number of seconds each TOTP code is valid for. Must be between 20 and 60. Defaults to `30`

## Import

A service can be imported using the `/Services/{sid}` format, e.g.

```shell
terraform import twilio_verify_service.service /Services/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
