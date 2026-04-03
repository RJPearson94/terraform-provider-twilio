---
page_title: "twilio_flex_plugin Data Source - twilio"
subcategory: "Flex"
description: |-
  
---

# twilio_flex_plugin Data Source

Use this data source to access information about an existing versioned Twilio Flex plugin. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/plugin) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

## Example Usage

### SID

```hcl
data "twilio_flex_plugin" "plugin" {
  sid = "FPXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "plugin" {
  value = data.twilio_flex_plugin.plugin
}
```

### Unique Name

```hcl
data "twilio_flex_plugin" "plugin" {
  unique_name = "UniqueName"
}

output "plugin" {
  value = data.twilio_flex_plugin.plugin
}
```

## Schema

### Optional

- `sid` (String) The SID of the Flex plugin to look up. Exactly one of `sid` or `unique_name` must be specified
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `unique_name` (String) The unique name of the Flex plugin to look up. Exactly one of `sid` or `unique_name` must be specified

### Read-Only

- `account_sid` (String) The SID of the account that owns this Flex plugin
- `archived` (Boolean) Whether the Flex plugin has been archived
- `changelog` (String) The changelog for the latest version of the plugin
- `date_created` (String) The date and time the Flex plugin was created, in RFC 3339 format
- `date_updated` (String) The date and time the Flex plugin was last updated, in RFC 3339 format
- `description` (String) A description of the Flex plugin
- `friendly_name` (String) A human-readable label for the Flex plugin
- `id` (String) The ID of this resource.
- `latest_version_sid` (String) The SID of the latest plugin version
- `plugin_url` (String) The hosted URL of the latest version of the plugin bundle
- `private` (Boolean) Whether the latest version of the plugin is private
- `url` (String) The absolute URL of the Flex plugin resource
- `version` (String) The version string of the latest plugin version
- `version_archived` (Boolean) Whether the latest version of the plugin has been archived

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
