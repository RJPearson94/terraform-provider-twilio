---
page_title: "Twilio SIP Trunking Credential Lists"
subcategory: "SIP Trunking"
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

## Argument Reference

The following arguments are supported:

- `trunk_sid` - (Mandatory) The SID of the SIP trunk the credential lists are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the `trunk_sid`)
- `trunk_sid` - The SID of the SIP trunk the credential lists are associated with (Same as `id`)
- `account_sid` - The account SID associated with the credential lists
- `credential_lists` - A list of `credential_list` blocks as documented below

---

A `credential_list` block supports the following:

- `sid` - The SID of the credential list
- `friendly_name` - The friendly name of the credential list
- `date_created` - The date in RFC3339 format that the credential list was created
- `date_updated` - The date in RFC3339 format that the credential list was updated
- `url` - The URL of the credential list resource

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving credential lists
