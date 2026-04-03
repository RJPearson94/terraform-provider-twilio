---
page_title: "twilio_flex_plugin Resource - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin Resource

Manages a versioned Flex plugin. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/plugin) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

## Example Usage

```hcl
resource "twilio_flex_plugin" "plugin" {
  unique_name = "twilio-test"
  version     = "1.0.0"
  plugin_url  = "https://example.com"
}
```

## Schema

### Required

- `plugin_url` (String) The hosted URL of the plugin bundle. Must use HTTP or HTTPS
- `unique_name` (String) The unique name of the Flex plugin. Changing this forces a new resource
- `version` (String) The version string for the plugin (e.g., `1.0.0`)

### Optional

- `changelog` (String) The changelog for this version of the plugin
- `description` (String) A description of the Flex plugin
- `friendly_name` (String) A human-readable label for the Flex plugin
- `private` (Boolean) Whether the plugin version is private
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this Flex plugin
- `archived` (Boolean) Whether the Flex plugin has been archived
- `date_created` (String) The date and time the Flex plugin was created, in RFC 3339 format
- `date_updated` (String) The date and time the Flex plugin was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `latest_version_sid` (String) The SID of the latest plugin version
- `sid` (String) The unique SID assigned to this Flex plugin by Twilio
- `url` (String) The absolute URL of the Flex plugin resource
- `version_archived` (Boolean) Whether the latest version of the plugin has been archived

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A plugin can be imported using the `/PluginService/Plugins/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin.plugin /PluginService/Plugins/FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
