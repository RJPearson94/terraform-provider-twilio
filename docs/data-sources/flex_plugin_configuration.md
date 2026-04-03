---
page_title: "twilio_flex_plugin_configuration Data Source - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin_configuration Data Source

Use this data source to access information about an existing Twilio Flex plugin configuration resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_flex_plugin_configuration" "plugin_configuration" {
  sid = "FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "plugin_configuration" {
  value = data.twilio_flex_plugin_configuration.plugin_configuration
}
```

## Schema

### Required

- `sid` (String) The SID of the Flex plugin configuration to look up

### Read-Only

- `account_sid` (String) The SID of the account that owns this plugin configuration
- `archived` (Boolean) Whether the plugin configuration has been archived
- `date_created` (String) The date and time the plugin configuration was created, in RFC 3339 format
- `description` (String) A description of the plugin configuration
- `id` (String) The ID of this resource.
- `name` (String) The name of the plugin configuration
- `plugins` (List of Object) A list of plugins included in this configuration (see [below for nested schema](#nestedatt--plugins))
- `url` (String) The absolute URL of the plugin configuration resource

<a id="nestedatt--plugins"></a>
### Nested Schema for `plugins`

Read-Only:

- `date_created` (String)
- `phase` (Number)
- `plugin_sid` (String)
- `plugin_url` (String)
- `plugin_version_sid` (String)
- `private` (Boolean)
- `unique_name` (String)
- `url` (String)
- `version` (String)
