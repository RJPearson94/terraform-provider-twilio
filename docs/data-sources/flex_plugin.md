---
page_title: "Twilio Flex Plugin"
subcategory: "Flex"
---

# twilio_flex_plugin Resource

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

## Argument Reference

The following arguments are supported:

- `sid` - (Optional) The SID of the plugin
- `unique_name` - (Optional) The unique name of the plugin

~> Either `sid` or `unique_name` must be specified

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

- `read` - (Defaults to 5 minutes) Used when retrieving the plugin
