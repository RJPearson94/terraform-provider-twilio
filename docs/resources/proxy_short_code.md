---
page_title: "Twilio Proxy Short Code"
subcategory: "Proxy"
---

# twilio_proxy_short_code Resource

Manages a proxy short code

## Example Usage

```hcl
resource "twilio_proxy_service" "service" {
  unique_name = "Test Proxy Service"
}

resource "twilio_proxy_short_code" "short_code" {
  service_sid = twilio_proxy_service.service.sid
  sid         = "SCXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  is_reserved = true
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of a Twilio proxy service. Changing this forces a new resource to be created
- `sid` - (Optional) The SID of a Twilio short code to associate with the proxy. Changing this forces a new resource to be created
- `is_reserved` - (Optional) Whether the short code is reserved

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the proxy short code resource (Same as the SID)
- `sid` - The SID of a Twilio short code to associate with the proxy (Same as the ID)
- `account_sid` - The Account SID of the short code resource is deployed into
- `service_sid` - The SID of a Twilio proxy service
- `is_reserved` - Whether the short code is reserved
- `short_code` - The short code associated with the SID
- `iso_country` - The ISO country associated with the SID
- `capabilities` - A `capabilities` block as documented below.
- `date_created` - The date in RFC3339 format that the proxy short code resource was created
- `date_updated` - The date in RFC3339 format that the proxy short code resource was updated
- `url` - The url of the proxy short code resource

---

A `capabilities` block supports the following:

- `fax_inbound` - Whether the short code is able to accept inbound faxes
- `fax_outbound` -  Whether the short code is able to send outbound faxes
- `mms_inbound` - Whether the short code is able to accept inbound MMS's
- `mms_outbound` -  Whether the short code is able to send outbound MMS's
- `restriction_fax_domestic` - Whether the short code is restricted to domestic faxes
- `restriction_mms_domestic` - Whether the short code is restricted to domestic MMS'
- `restriction_sms_domestic` - Whether the short code is restricted to domestic SMS's
- `restriction_voice_domestic` - Whether the short code is restricted to domestic voice calls
- `sip_trunking` -  Whether the short code supports SIP trunking
- `sms_inbound` - Whether the short code is able to accept inbound SMS's
- `sms_outbound` - Whether the short code is able to send outbound SMS's
- `voice_inbound` - Whether the short code is able to accept inbound voice calls
- `voice_outbound` - Whether the short code is able to make outbound voice calls

---
