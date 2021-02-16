---
page_title: "Twilio Serverless Builds"
subcategory: "Serverless"
---

# twilio_serverless_builds Data Source

Use this data source to access information about the builds associated with an existing Serverless service. See the [API docs](https://www.twilio.com/docs/runtime/functions-assets-api/api/build) for more information

For more information on Serverless (also known as Runtime), see the product [page](https://www.twilio.com/runtime)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
data "twilio_serverless_builds" "builds" {
  service_sid = "ZSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "builds" {
  value = data.twilio_serverless_builds.builds
}
```

## Argument Reference

The following arguments are supported:

- `service_sid` - (Mandatory) The SID of the service the builds are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the service SID)
- `service_sid` - The SID of the service the builds are associated with
- `account_sid` - The account SID associated with the builds
- `builds` - A list of `build` blocks as documented below

---

A `build` block supports the following:

- `sid` - The SID of the build
- `asset_version` - A `asset_version` block as documented below.
- `function_version` - A `function_version` block as documented below.
- `dependencies` - Map of dependencies to be included in the build
- `runtime` - The target runtime of the serverless functions and assets
- `status` - The current status of the build job
- `date_created` - The date in RFC3339 format that the build was created
- `date_updated` - The date in RFC3339 format that the build was updated
- `url` - The URL of the build

---

An `asset_version` block supports the following:

- `sid` - The SID of the asset version
- `account_sid` - The account SID of the asset version is deployed into
- `service_sid` - The account SID of the asset version is deployed into
- `asset_sid` - The asset SID of the version is managed under
- `date_created` - The date in RFC3339 format that the asset version was created
- `path` - The request URI path
- `visibility` - The visibility of the asset

---

A `function_version` block supports the following:

- `sid` - The SID of the function version
- `account_sid` - The account SID of the function version is deployed into
- `service_sid` - The account SID of the function version is deployed into
- `function_sid` - The function SID of the version is managed under
- `date_created` - The date in RFC3339 format that the function version was created
- `path` - The request URI path
- `visibility` - The visibility of the function

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving builds
