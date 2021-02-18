---
page_title: "Twilio Flex Plugin Configuration"
subcategory: "Flex"
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

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the plugin configuration

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

- `read` - (Defaults to 5 minutes) Used when retrieving the plugin configuration
