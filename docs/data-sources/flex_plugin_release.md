---
page_title: "twilio_flex_plugin_release Data Source - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin_release Data Source

Use this data source to access information about an existing Twilio Flex plugin release resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/release) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_flex_plugin_release" "plugin_release" {
  sid = "FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "plugin_release" {
  value = data.twilio_flex_plugin_release.plugin_release
}
```

## Schema

### Required

- `sid` (String) The SID of the Flex plugin release to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this plugin release
- `configuration_sid` (String) The SID of the plugin configuration associated with this release
- `date_created` (String) The date and time the plugin release was created, in RFC 3339 format
- `id` (String) The ID of this resource.
- `url` (String) The absolute URL of the plugin release resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
