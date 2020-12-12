---
page_title: "Twilio SIP Trunking Trunk"
subcategory: "SIP Trunking"
---

# twilio_sip_trunking_trunk Resource

Manages a SIP trunk. See the [API docs](https://www.twilio.com/docs/sip-trunking/api/trunk-resource) for more information

For more information on SIP Trunking, see the product [page](https://www.twilio.com/docs/sip-trunking)

## Example Usage

```hcl
resource "twilio_sip_trunking_trunk" "trunk" {
  friendly_name = "twilio-test"
}
```

## Argument Reference

The following arguments are supported:

- `cnam_lookup_enabled` - (Optional) Whether Caller ID Name is enabled on the SIP trunk
- `disaster_recovery_url` - (Optional) The URL to call in event of disaster recovery.Valid values are `POST` or `GET`
- `disaster_recovery_method` - (Optional) The HTTP method which should be used to call the disaster recovery URL
- `domain_name` - (Optional) The domain name of the SIP trunk
- `friendly_name` - (Optional) The friendly name of the SIP trunk
- `recording` - (Optional) A `recording` block as documented below
- `secure` - (Optional) Whether secure trunking is enabled on the SIP trunk
- `transfer_mode` - (Optional) The call transfer configuration on the SIP trunk

---

A `recording` block supports the following:

- `mode` - (Optional) The recording mode configuration for the SIP trunk. Valid values are `do-not-record`, `record-from-ringing`, `record-from-answer`, `record-from-ringing-dual` or `record-from-answer-dual`
- `trim` - (Optional) The recording trim configuration for the SIP trunk. Valid values are `trim-silence` or `do-not-trim`

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the SIP trunk (Same as the SID)
- `sid` - The SID of the SIP trunk (Same as the ID)
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
- `auth_type_set` - The auth type set on the SIP trunk
- `date_created` - The date in RFC3339 format that the SIP trunk was created
- `date_updated` - The date in RFC3339 format that the SIP trunk was updated
- `url` - The URL of the SIP trunk resource

---

A `recording` block supports the following:

- `mode` - The recording mode configuration for the SIP trunk
- `trim` - The recording trim configuration for the SIP trunk

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the SIP trunk
- `update` - (Defaults to 10 minutes) Used when updating the SIP trunk
- `read` - (Defaults to 5 minutes) Used when retrieving the SIP trunk
- `delete` - (Defaults to 10 minutes) Used when deleting the SIP trunk

## Import

A SIP trunk can be imported using the `/Trunks/{sid}` format, e.g.

```shell
terraform import twilio_sip_trunking_trunk.trunk /Trunks/TKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
