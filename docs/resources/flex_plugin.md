---
page_title: "Twilio Flex Plugin"
subcategory: "Flex"
---

# twilio_flex_plugin Resource

Manages a versioned Flex plugin. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/plugin) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> This API used to manage this resource is currently in beta and is subject to change

!> The plugin API does not support deleting or archiving of plugin resources therefore the provider will not remove or modify any resources when running a destroy, only the state for the resource will be removed.

## Example Usage

```hcl
resource "twilio_flex_plugin" "plugin" {
  unique_name = "twilio-test"
  version     = "1.0.0"
  plugin_url  = "https://example.com"
}
```

## Argument Reference

The following arguments are supported:

- `unique_name` - (Mandatory) The unique name of the plugin. Changing this forces a new resource to be created
- `friendly_name` - (Optional) The friendly name of the plugin
- `description` - (Optional) The description of the plugin
- `changelog` - (Optional) The changelog for the plugin
- `version` - (Mandatory) The version of the plugin
- `plugin_url` - (Mandatory) The URL of the hosted plugin bundle
- `private` - (Optional) Whether credentials are required to access the plugin

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the plugin (Same as the `sid`)
- `sid` - The SID of the plugin (Same as the `id`)
- `account_sid` - The account SID associated with the plugin
- `unique_name` - The unique name of the plugin
- `friendly_name` - The friendly name of the plugin
- `description` - The description of the plugin
- `changelog` - The changelog for the plugin
- `version` - The version of the plugin
- `plugin_url` - The URL of the hosted plugin bundle
- `private` - Whether credentials are required to access the plugin
- `archived` - Whether the plugin has been archived
- `version_archived` - Whether the latest plugin version has been archived
- `latest_version_sid` - The SID of the latest plugin version
- `date_created` - The date in RFC3339 format that the plugin was created
- `date_updated` - The date in RFC3339 format that the plugin was updated
- `url` - The URL of the plugin

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the plugin
- `update` - (Defaults to 10 minutes) Used when updating the plugin
- `read` - (Defaults to 5 minutes) Used when retrieving the plugin

## Import

A plugin can be imported using the `/PluginService/Plugins/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin.plugin /PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
