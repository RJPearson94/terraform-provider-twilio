---
page_title: "twilio_conversations_push_credential_fcm Resource - twilio"
subcategory: "Conversations"
description: |-
  
---

# twilio_conversations_push_credential_fcm Resource

Manages push credentials to allow Twilio Conversations to send push notifications via Firebase Cloud Messaging (FCM). See the [API docs](https://www.twilio.com/docs/conversations/api/credential-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_push_credential_fcm" "push_credential_fcm" {
  friendly_name = "fcm-credential"
  secret        = "<<fcm_secret>>"
}
```

## Schema

### Required

- `friendly_name` (String) A human-readable label for the FCM push credential
- `secret` (String, Sensitive) The FCM server key or secret

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this FCM push credential
- `date_created` (String) The date and time the FCM push credential was created, in RFC 3339 format
- `date_updated` (String) The date and time the FCM push credential was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this FCM push credential by Twilio
- `type` (String) The type of push credential
- `url` (String) The absolute URL of the FCM push credential resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

FCM push credentials can be imported using the `/Credentials/{sid}` format, e.g.

```shell
terraform import twilio_conversations_push_credential_fcm.push_credential_fcm /Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
