---
page_title: "Twilio Studio Flow"
subcategory: "Studio"
---

# twilio_studio_flow Resource

Manages a Studio flow. See the [API docs](https://www.twilio.com/docs/studio/rest-api/v2/flow) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

!> This API used to manage this resource is currently in beta and is subject to change

## Example Usage

```hcl
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test studio flow"
  status        = "draft"
  definition    = <<EOF
{
  "description": "A New Flow",
  "flags": {
    "allow_concurrent_calls": true
  },
  "initial_state": "Trigger",
  "states": [
    {
      "name": "Trigger",
      "properties": {
        "offset": {
          "x": 0,
          "y": 0
        }
      },
      "transitions": [],
      "type": "trigger"
    }
  ]
}
EOF
  validate = true
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The name of the studio flow
- `status` - (Mandatory) The status of the studio flow. Valid values include `draft` and `published`
- `definition` - (Mandatory) The Flow Definition JSON
- `validate` - (Optional) Whether to validate the Flow Definition JSON before creating a new revision. The default is false
- `commit_message` - (Optional) Description of the changes made

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the studio flow (Same as the SID)
- `sid` - The SID of the studio flow (Same as the ID)
- `friendly_name` - The name of the studio flow
- `definition` - The Flow Definition JSON
- `status` - The status of the studio flow
- `revision` - The revision number of the studio flow
- `valid` - Whether the studio flow is valid
- `validate` - Whether the studio flow has been validated on creation and on updates
- `date_created` - The date in RFC3339 format that the studio flow was created
- `date_updated` - The date in RFC3339 format that the studio flow was updated
- `url` - The url of the studio flow
- `webhook_url` - The webhook url of the studio flow

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 10 minutes) Used when creating the studio flow
- `update` - (Defaults to 10 minutes) Used when updating the studio flow
- `read` - (Defaults to 5 minutes) Used when retrieving the studio flow
- `delete` - (Defaults to 10 minutes) Used when deleting the studio flow

!> When request validation is enabled, the request is constrained by its own create timeout as defined above

## Import

A flow can be imported using the `/Flows/{sid}` format, e.g.

```shell
terraform import twilio_studio_flow.flow /Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
