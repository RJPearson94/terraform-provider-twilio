---
page_title: "Twilio Flex Plugin Release"
subcategory: "Flex"
---

# twilio_flex_plugin_release Resource

Use this data source to access information about an existing Twilio Flex plugin release resource. See the [API docs](https://www.twilio.com/docs/flex/developer/plugins/api/release) for more information

For more information on Twilio Flex, see the product [page](https://www.twilio.com/flex)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_flex_plugin_release" "plugin_release" {
  sid = "FKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "plugin_release" {
  value = data.twilio_flex_plugin_release.plugin_release
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the plugin release

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

- `read` - (Defaults to 5 minutes) Used when retrieving the plugin release
