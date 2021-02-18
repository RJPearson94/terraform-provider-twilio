---
page_title: "Twilio SIP Domain"
subcategory: "SIP"
---

# twilio_sip_domain Data Source

Use this data source to access information about an existing domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-domain-resource) for more information

## Example Usage

```hcl
data "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid         = "DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "domain" {
  value = data.twilio_sip_domain.domain
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The SID of the account the domain is associated with
- `sid` - (Mandatory) The SID of the domain

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the domain (Same as the `sid`)
- `sid` - The SID of the domain (Same as the `id`)
- `account_sid` - The account SID associated with the domain
- `domain_name` - The domain name of the resource
- `friendly_name` - The friendly name of the domain
- `voice` - A `voice` block as documented below
- `emergency` - A `emergency` block as documented below
- `byoc_trunk_sid` - The BYOC trunk SID to associate the domain with
- `secure` - Whether secure SIP is enabled
- `sip_registration` - Whether the SIP endpoint is allowed to register with the domain
- `auth_type` - The authentication for the domain
- `date_created` - The date in RFC3339 format that the domain was created
- `date_updated` - The date in RFC3339 format that the domain was updated

---

A `voice` block supports the following:

- `url` - The URL which should be called on each incoming call
- `method` - The HTTP method which should be used to call the URL
- `fallback_url` - The URL which should be called when the URL request fails
- `fallback_method` - The HTTP method which should be used to call the fallback URL
- `status_callback_url` - The URL to call on each status change
- `status_callback_method` - The HTTP method which should be used to call the status callback URL

---

An `emergency` block supports the following:

- `calling_enabled` - Whether emergency calling is enabled for the domain
- `caller_sid` - The caller SID to associate with the domain

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the domain details
