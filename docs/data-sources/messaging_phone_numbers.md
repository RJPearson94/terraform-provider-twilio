---
page_title: "twilio_messaging_phone_numbers Data Source - twilio"
subcategory: "Programmable Messaging"
description: |-
  
---

# twilio_messaging_phone_numbers Data Source

Use this data source to access information about the phone numbers associated with an existing Programmable Messaging service. See the [API docs](https://www.twilio.com/docs/sms/services/api/phonenumber-resource) for more information

For more information on Programmable Messaging, see the product [page](https://www.twilio.com/messaging)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_messaging_phone_numbers" "phone_numbers" {
  service_sid = "MGXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "phone_numbers" {
  value = data.twilio_messaging_phone_numbers.phone_numbers
}
```

## Schema

### Required

- `service_sid` (String) The SID of the messaging service to retrieve phone numbers for

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns these phone numbers
- `id` (String) The ID of this resource.
- `phone_numbers` (List of Object) The list of phone numbers associated with the messaging service (see [below for nested schema](#nestedatt--phone_numbers))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)


<a id="nestedatt--phone_numbers"></a>
### Nested Schema for `phone_numbers`

Read-Only:

- `capabilities` (List of String)
- `country_code` (String)
- `date_created` (String)
- `date_updated` (String)
- `phone_number` (String)
- `sid` (String)
- `url` (String)
