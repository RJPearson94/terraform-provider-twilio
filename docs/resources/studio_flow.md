---
page_title: "twilio_studio_flow Resource - twilio"
subcategory: "Studio"
description: |-
  
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

## Schema

### Required

- `definition` (String) A JSON string defining the flow structure. Use the `twilio_studio_flow_definition` data source to build this value
- `friendly_name` (String) A human-readable label for the flow, unique within the account
- `status` (String) The lifecycle status of the flow. Valid values: `draft`, `published`

### Optional

- `commit_message` (String) A description of the changes made in this flow revision. Defaults to "Updated via Terraform"
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `validate` (Boolean) Whether to validate the flow definition against the Twilio API before saving. Defaults to false

### Read-Only

- `account_sid` (String) The SID of the account that owns this flow
- `date_created` (String) The date and time the flow was created, in RFC 3339 format
- `date_updated` (String) The date and time the flow was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `revision` (Number) The current revision number of the flow
- `sid` (String) The unique SID assigned to this flow by Twilio
- `url` (String) The absolute URL of the flow resource
- `valid` (Boolean) Whether the flow definition passed Twilio's validation checks
- `webhook_url` (String) The webhook URL used to trigger this flow via the REST API

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A flow can be imported using the `/Flows/{sid}` format, e.g.

```shell
terraform import twilio_studio_flow.flow /Flows/FWXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

!> `validate` cannot be imported
