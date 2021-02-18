---
page_title: "Twilio SIP Trunking Trunk"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_trunk Data Source

Use this data source to access information about an existing SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/trunk-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {
  sid = "TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "trunk" {
  value = data.twilio_sip_trunking_trunk.trunk
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the SIP trunk

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the SIP trunk (Same as the `sid`)
- `sid` - The SID of the SIP trunk (Same as the `id`)
- `account_sid` - The account SID the SIP trunk is associated with
- `cnam_lookup_enabled` - Whether Caller ID Name is enabled on the SIP trunk
- `disaster_recovery_url` - The URL to call in event of disaster recovery
- `disaster_recovery_method` - The HTTP method which should be used to call the disaster recovery URL
- `domain_name` - The domain name of the SIP trunk
- `friendly_name` - The friendly name of the SIP trunk
- `recording` - A `recording` block as documented below
- `secure` - Whether secure trunking is enabled on the SIP trunk
- `transfer_mode` - The call transfer configuration on the SIP trunk
- `auth_type` - The auth configuration on the SIP trunk
- `auth_type_set` - The auth typeset on the SIP trunk
- `date_created` - The date in RFC3339 format that the SIP trunk was created
- `date_updated` - The date in RFC3339 format that the SIP trunk was updated
- `url` - The URL of the SIP trunk resource

---

A `recording` block supports the following:

- `mode` - The recording mode configuration for the SIP trunk
- `trim` - The recording trim configuration for the SIP trunk

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the SIP trunk details
