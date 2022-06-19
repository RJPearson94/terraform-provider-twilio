---
page_title: "Twilio Verify Service"
subcategory: "Verify"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the service

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the service (Same as the `sid`)
- `sid` - The SID of the service (Same as the `id`)
- `account_sid` - The account SID of the service is associated with
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

- `read` - (Defaults to 5 minutes) Used when retrieving the service
