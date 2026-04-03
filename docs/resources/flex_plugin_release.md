---
page_title: "twilio_flex_plugin_release Resource - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin_release Resource

Manages a Flex plugin release resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/release) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> If this resource is deleted and the release is the latest Twilio Flex Plugin release. A new configuration without any plugins will be created. This configuration will then be deployed as a new release to supersede the existing release

~> To allow terraform to correctly manage the lifecycle of the release, it is recommended that use the lifecycle meta-argument `create_before_destroy` with this resource. The docs can be found [here](https://www.terraform.io/docs/configuration/resources.html#create_before_destroy)

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

  lifecycle {
    create_before_destroy = true
  }
}

resource "twilio_flex_plugin_release" "plugin_release" {
  configuration_sid = twilio_flex_plugin_configuration.plugin_configuration.sid

  lifecycle {
    create_before_destroy = true
  }
}
```

## Schema

### Required

- `configuration_sid` (String) The SID of the plugin configuration to deploy as a release. Changing this forces a new resource

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this plugin release
- `date_created` (String) The date and time the plugin release was created, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this plugin release by Twilio
- `url` (String) The absolute URL of the plugin release resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)

## Import

A plugin release can be imported using the `/PluginService/Releases/{sid}` format, e.g.

```shell
terraform import twilio_flex_plugin_release.plugin_release /PluginService/Releases/FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
