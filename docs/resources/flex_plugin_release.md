---
page_title: "Twilio Flex Plugin Release"
subcategory: "Flex"
---

# twilio_flex_plugin_release Resource

Manages a Flex plugin release resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/release) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> This API used to manage this resource is currently in beta and is subject to change

!> The plugin release API does not support deleting or archiving of release resources therefore the provider will not remove or modify any resources when running a destroy, only the state for the resource will be removed.

!> If you need to remove a current release from Flex you will need to create a new release resource with the desired configuration

## Example Usage

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
}

resource "twilio_flex_plugin_release" "plugin_release" {
  configuration_sid = twilio_flex_plugin_configuration.plugin_configuration.sid
}
```

## Argument Reference

The following arguments are supported:

- `configuration_sid` - (Mandatory) The SID of the configuration to associate with the release. Changing this forces a new resource to be created

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the plugin release (Same as the `sid`)
- `sid` - The SID of the plugin release (Same as the `id`)
- `account_sid` - The account SID associated with the plugin release
- `configuration_sid` - The SID of the configuration associated with the release
- `date_created` - The date in RFC3339 format that the plugin release was created
- `url` - The URL of the plugin release

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/release/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the plugin release
- `read` - (Defaults to 5 minutes) Used when retrieving the plugin release

## Import

A plugin release can be imported using the `/PluginService/Releases/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin_release.plugin_release /PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
