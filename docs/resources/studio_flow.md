---
page_title: "Twilio Studio Flow"
subcategory: "Studio"
---

# twilio_studio_flow Resource

Manages a Studio flow. See the [API docs](https://www.twilio.com/docs/studio/rest-api/v2/flow) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

```hcl
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test studio flow"
  status        = "draft"
  definition = jsonencode({
    "description" : "A New Flow",
    "flags" : {
      "allow_concurrent_calls" : true
    },
    "initial_state" : "Trigger",
    "states" : [
      {
        "name" : "Trigger",
        "properties" : {
          "offset" : {
            "x" : 0,
            "y" : 0
          }
        },
        "transitions" : [],
        "type" : "trigger"
      }
    ]
  })
  validate = true
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the Studio flow
- `status` - (Mandatory) The status of the Studio flow. Valid values include `draft` and `published`
- `definition` - (Mandatory) The flow definition JSON
- `validate` - (Optional) Whether to validate the flow definition JSON before creating a new revision. The default is `false`
- `commit_message` - (Optional) Description of the changes made

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the Studio flow (Same as the `sid`)
- `sid` - The SID of the Studio flow (Same as the `id`)
- `friendly_name` - The name of the Studio flow
- `definition` - The flow definition JSON
- `status` - The status of the Studio flow
- `revision` - The revision number of the Studio flow
- `valid` - Whether the Studio flow is valid
- `date_created` - The date in RFC3339 format that the Studio flow was created
- `date_updated` - The date in RFC3339 format that the Studio flow was updated
- `url` - The URL of the Studio flow
- `webhook_url` - The webhook URL of the Studio flow

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the Studio flow
- `update` - (Defaults to 10 minutes) Used when updating the Studio flow
- `read` - (Defaults to 5 minutes) Used when retrieving the Studio flow
- `delete` - (Defaults to 10 minutes) Used when deleting the Studio flow

!> When request validation is enabled, the request is constrained by its own create timeout as defined above

## Import

A flow can be imported using the `/Flows/{sid}` format, e.g.

```shell
terraform import twilio_studio_flow.flow /Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `validate` cannot be imported
