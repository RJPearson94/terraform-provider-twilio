---
page_title: "twilio_conversations_push_credential_apn Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_push_credential_apn Resource

Manages push credentials to allow Twilio Conversations to send push notifications via Apple Push Notification Service (APN). See the [API docs](https://www.twilio.com/docs/conversations/api/credential-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_push_credential_apn" "push_credential_apn" {
  friendly_name = "apn-credential"
  certificate   = "<<certificate>>"
  private_key   = "<<private_key>>"
}
```

## Schema

### Required

- `certificate` (String) The APN service certificate
- `friendly_name` (String) A human-readable label for the APN push credential
- `private_key` (String, Sensitive) The APN service private key

### Optional

- `sandbox` (Boolean) Whether to use the APN sandbox environment. Defaults to `false`
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this APN push credential
- `date_created` (String) The date and time the APN push credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the APN push credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this APN push credential by Twilio
- `type` (String) The type of push credential
- `url` (String) The absolute URL of the APN push credential resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

APN push credentials can be imported using the `/Credentials/{sid}` format, e.g.

```shell
terraform import twilio_conversations_push_credential_apn.push_credential_apn /Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
