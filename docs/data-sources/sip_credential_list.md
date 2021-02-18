---
page_title: "Twilio SIP Credential List"
subcategory: "SIP"
---

# twilio_sip_credential_list Data Source

Use this data source to access information about an existing credential list. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/credentiallist-resource) for more information

## Example Usage

```hcl
data "twilio_sip_credential_list" "credential_list" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "CLXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "credential_list" {
  value = data.twilio_sip_credential_list.credential_list
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the credential list is associated with
- `sid` - (Mandatory) The SID of the credential list

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the credential list (Same as the `sid`)
- `sid` - The SID of the credential list (Same as the `id`)
- `account_sid` - The account SID associated with the credential list
- `friendly_name` - The friendly name of the credential list
- `date_created` - The date in RFC3339 format that the credential list was created
- `date_updated` - The date in RFC3339 format that the credential list was updated

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the credential list details
