---
page_title: "Twilio Flex Plugin Configuration"
subcategory: "Flex"
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

## Argument Reference

The following arguments are supported:

- `name` - (Mandatory) The name of the plugin configuration. Changing this forces a new resource to be created
- `description` - (Optional) The description of the plugin configuration. Changing this forces a new resource to be created
- `plugins` - (Optional) A list of `plugin` blocks as documented below. Changing this forces a new resource to be created

---

A `plugin` block supports the following:

- `plugin_version_sid` - (Mandatory) The SID of the plugin version to associate with the configuration. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the plugin configuration (Same as the `sid`)
- `sid` - The SID of the plugin configuration (Same as the `id`)
- `account_sid` - The account SID associated with the plugin configuration
- `name` - (Mandatory) The name of the plugin configuration
- `description` - (Optional) The description of the plugin configuration
- `plugins` - A `plugin` block as documented below
- `archived` - Whether the plugin configuration has been archived
- `date_created` - The date in RFC3339 format that the plugin configuration was created
- `url` - The URL of the plugin configuration

---

A `plugin` block supports the following:

- `plugin_version_sid` - The SID of the plugin version associated with the configuration
- `plugin_sid` - The SID of the plugin associated with the configuration
- `plugin_url` - The URL of the hosted plugin bundle
- `private` - Whether credentials are required to access the plugin
- `unique_name` - The unique name of the plugin
- `private` - Whether credentials are required to access the plugin
- `phase` - The phase number of the plugin
- `version` - The version of the plugin
- `date_created` - The date in RFC3339 format that the plugin was created
- `url` - The URL of the plugin

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the plugin configuration
- `read` - (Defaults to 5 minutes) Used when retrieving the plugin configuration

## Import

A plugin configuration can be imported using the `/PluginService/Configurations/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin_configuration.plugin_configuration /PluginService/Configurations/FJXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
