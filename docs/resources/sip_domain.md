---
page_title: "Twilio SIP Domain"
subcategory: "SIP"
---

# twilio_sip_domain Resource

Manages a SIP domain. See the [API docs](https://www.twilio.com/docs/voice/sip/api/sip-domain-resource) for more information

## Example Usage

```hcl
resource "twilio_sip_domain" "domain" {
  account_sid = "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  domain_name = "test.sip.twilio.com"
}
```

## Argument Reference

The following arguments are supported:

- `account_sid` - (Mandatory) The account SID to associate the domain with. Changing this forces a new resource to be created
- `domain_name` - (Mandatory) The domain name of the resource. The domain name must end with `.sip.twilio.com`
- `friendly_name` - (Optional) The friendly name of the domain
- `voice` - (Optional) A `voice` block as documented below
- `emergency` - (Optional) A `emergency` block as documented below
- `byoc_trunk_sid` - (Optional) The BYOC trunk SID to associate the domain with
- `secure` - (Optional) Whether secure SIP is enabled. The default value is `false`
- `sip_registration` - (Optional) Whether the SIP endpoint is allowed to register with the domain. The default value is `false`

---

A `voice` block supports the following:

- `url` - (Optional) The URL which should be called on each incoming call
- `method` - (Optional) The HTTP method which should be used to call the URL. Valid values are `GET` or `POST`. The default value is `POST`
- `fallback_url` - (Optional) The URL which should be called when the URL request fails
- `fallback_method` - (Optional) The HTTP method which should be used to call the fallback URL. Valid values are `GET` or `POST`. The default value is `POST`
- `status_callback_url` - (Optional) The URL to call on each status change
- `status_callback_method` - (Optional) The HTTP method which should be used to call the status callback URL. Valid values are `GET` or `POST`. The default value is `POST`

---

An `emergency` block supports the following:

- `calling_enabled` - (Optional) Whether emergency calling is enabled for the domain. The default value is `false`
- `caller_sid` - (Optional) The caller SID to associate with the domain

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

- `create` - (Defaults to 10 minutes) Used when creating the domain
- `update` - (Defaults to 10 minutes) Used when updating the domain
- `read` - (Defaults to 5 minutes) Used when retrieving the domain
- `delete` - (Defaults to 10 minutes) Used when deleting the domain

## Import

An domain can be imported using the `Accounts/{accountSid}/Domains/{sid}` format, e.g.

```shell
terraform import twilio_sip_domain.domain /Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Domains/DSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
