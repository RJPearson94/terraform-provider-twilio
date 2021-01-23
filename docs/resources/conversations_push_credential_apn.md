---
page_title: "Twilio Conversations Push Credentials (APN)"
subcategory: "Conversations"
---

# twilio_conversations_push_credential_apn Resource

Manages push credentials to allow Twilio Conversations to send push notifications via Apple Push Notification Service (APN). See the [API docs](https://www.twilio.com/docs/conversations/api/credential-resource) for more information

For more information on conversations, see the product [page](https://www.twilio.com/conversations)

## Example Usage

```hcl
resource "twilio_conversations_push_credential_apn" "push_credential_apn" {
  friendly_name = "apn-credential"
  certificate = "<<certificate>>"
  private_key = "<<private_key>>"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the push credentials
- `certificate` - (Mandatory) The APN certificate
- `private_key` - (Mandatory) The APN private key
- `sandbox` - (Optional) Whether to use the sandbox APN

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the push credentials (Same as the SID)
- `sid` - The SID of the push credentials  (Same as the ID)
- `account_sid` - The account SID associated with the push credentials 
- `friendly_name` - The friendly name of the push credentials
- `certificate` - The APN certificate
- `private_key` - The APN private key
- `type` - What notification service the credentials are for. The value will be `apn`
- `sandbox` - Whether to use the sandbox APN
- `date_created` - The date in RFC3339 format that the push credentials were created
- `date_updated` - The date in RFC3339 format that the push credentials were updated
- `url` - The URL of the push credentials

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the push credentials
- `update` - (Defaults to 10 minutes) Used when updating the push credentials
- `read` - (Defaults to 5 minutes) Used when retrieving the push credentials
- `delete` - (Defaults to 10 minutes) Used when deleting the push credentials

## Import

APN push credentials can be imported using the `/Credentials/{sid}` format, e.g.

```shell
terraform import twilio_conversations_push_credential_apn.push_credential_apn /Credentials/CRXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
