---
page_title: "Twilio Verify Service"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the service. The length of the string must be between `1` and `30` characters (inclusive)
- `code_length` - (Optional) The length of the verification code to send. The value must be between `4` and `10` (inclusive). The default value is `6`
- `custom_code_enabled` - (Optional) Whether to allow the code to be generated via custom code. The default value is `false`
- `do_not_share_warning_enabled` - (Optional) Whether to include a warning not share the code with anyone. The default value is `false`
- `dtmf_input_required` - (Optional) Whether to request the user presses a key when the code is delivered over the phone. The default value is `true`
- `default_template_sid` - (Optional) The default template SID that will be used using the verification process
- `lookup_enabled` - (Optional) Whether to perform a lookup when starting the verification process. The default value is `true`
- `mailer_sid` - (Optional) The mailer SID to associate with the service
- `psd2_enabled` - (Optional) Whether to pass PSD2 parameters when starting the verification process. The default value is `false`
- `push` - (Optional) A `push` block as documented below.
- `skip_sms_to_landlines` - (Optional) Whether to skip sending SMS's to landline numbers. The default value is `false`
  ~> To use this feature, `lookup_enabled` must be set to `true`
- `totp` - (Optional) A `totp` block as documented below.
- `tts_name` - (Optional) The name of the Text to Speech service to use when sending verification code over the phone

---

A `push` block supports the following:

- `apn_credential_sid` - (Optional) The APN credentials SID to associate with the service
- `fcm_credential_sid` - (Optional) The FCM credentials SID to associate with the service

---

A `totp` block supports the following:

- `issuer` - (Optional) The TOTP issuer for the service
- `time_step` - (Optional) The number of seconds between generating new TOTP code. The value must be between `20` and `60` (inclusive). The default value is `30`
- `code_length` - (Optional) The number of digits in the generated TOTP code. The value must be between `3` and `8` (inclusive). The default value is `6`
- `skew` - (Optional) The number of codes in the past and future that will be accepted. The value must be between `0` and `2` (inclusive). The default value is `1`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The account SID the service is associated with
- `friendly_name` - The friendly name of the service
- `code_length` - The length of the verification code to send
- `custom_code_enabled` - Whether to allow the code to be generated via custom code
- `do_not_share_warning_enabled` - Whether to include a warning to not share the code with anyone
- `dtmf_input_required` - Whether to request the user presses a key when the code is delivered over the phone
- `default_template_sid` - The default template SID that will be used using the verification process
- `lookup_enabled` - Whether to perform a lookup when starting the verification process
- `mailer_sid` - The mailer SID to associate with the service
- `psd2_enabled` - Whether to pass PSD2 parameters when starting the verification process
- `push` - A `push` block as documented below.
- `skip_sms_to_landlines` - Whether to skip sending SMS's to landline numbers
- `totp` - A `totp` block as documented below
- `tts_name` - The name of the Text to Speech service to use when sending verification code over the phone
- `date_created` - The date in RFC3339 format that the service was created
- `date_updated` - The date in RFC3339 format that the service was updated
- `url` - The URL of the service

---

A `push` block supports the following:

- `apn_credential_sid` - The APN credentials SID to associate with the service
- `fcm_credential_sid` - The FCM credentials SID to associate with the service

---

A `totp` block supports the following:

- `issuer` - The TOTP issuer for the service
- `time_step` - The number of seconds between generating new TOTP code
- `code_length` - The number of digits in the generated TOTP code
- `skew` - The number of codes in the past and future that will be accepted

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the service
- `update` - (Defaults to 10 minutes) Used when updating the service
- `read` - (Defaults to 5 minutes) Used when retrieving the service
- `delete` - (Defaults to 10 minutes) Used when deleting the service

## Import

A service can be imported using the `/Service/{sid}` format, e.g.

```shell
terraform import twilio_verify_service.service /Service/VAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
