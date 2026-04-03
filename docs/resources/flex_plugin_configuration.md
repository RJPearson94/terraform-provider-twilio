---
page_title: "twilio_flex_plugin_configuration Resource - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin_configuration Resource

Manages a Flex plugin configuration resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/plugin-configuration) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

~> To allow terraform to correctly manage the lifecycle of the configuration, it is recommended that use the lifecycle meta-argument `create_before_destroy` with this resource. The docs can be found [here](https://www.terraform.io/docs/configuration/resources.html#create_before_destroy)

## Example Usage

### With no plugins

```hcl
resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "twilio-test"
}
```

### With 1 plugin

```hcl
resource "twilio_flex_plugin" "plugin" {
  unique_name = "twilio-test"
  version     = "1.0.0"
  plugin_url  = "https://example.com"
}

resource "twilio_flex_plugin_configuration" "plugin_configuration" {
  name = "twilio-test"
  plugins {
    plugin_version_sid = twilio_flex_plugin.plugin.latest_version_sid
  }

  lifecycle {
    create_before_destroy = true
  }
}
```

## Schema

### Required

- `name` (String) The name of the plugin configuration. Changing this forces a new resource

### Optional

- `description` (String) A description of the plugin configuration. Changing this forces a new resource
- `plugins` (Block List) A list of plugin versions to include in this configuration. Changing this forces a new resource (see [below for nested schema](#nestedblock--plugins))
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this plugin configuration
- `archived` (Boolean) Whether the plugin configuration has been archived
- `date_created` (String) The date and time the plugin configuration was created, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this plugin configuration by Twilio
- `url` (String) The absolute URL of the plugin configuration resource

<a id="nestedblock--plugins"></a>
### Nested Schema for `plugins`

Required:

- `plugin_version_sid` (String) The SID of the plugin version to include in the configuration. Changing this forces a new resource

Read-Only:

- `date_created` (String) The date and time the plugin was created, in RFC 3339 format
- `phase` (Number) The load order phase of the plugin
- `plugin_sid` (String) The SID of the plugin
- `plugin_url` (String) The hosted URL of the plugin bundle
- `private` (Boolean) Whether the plugin version is private
- `unique_name` (String) The unique name of the plugin
- `url` (String) The absolute URL of the plugin resource
- `version` (String) The version string of the plugin


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A plugin configuration can be imported using the `/PluginService/Configurations/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin_configuration.plugin_configuration /PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
