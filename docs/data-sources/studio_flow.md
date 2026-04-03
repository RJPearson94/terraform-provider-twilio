---
page_title: "twilio_studio_flow Data Source - twilio"
subcategory: "Studio"
description: |-
  
---

# twilio_studio_flow Data Source

Use this data source to access information about an existing studio flow. See the [API docs](https://www.twilio.com/docs/studio/rest-api/v2/flow) for more information

For more information on Studio, see the product [page](https://www.twilio.com/studio)

## Example Usage

```hcl
data "twilio_studio_flow" "flow" {
  sid = "FWxxxxxxxxxxxxxxxx"
}

output "definition" {
  value = data.twilio_studio_flow.flow.definition
}
```

## Schema

### Required

- `sid` (String) The SID of the Studio flow to read

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this flow
- `commit_message` (String) The description of the changes made in the current flow revision
- `date_created` (String) The date and time the flow was created, in RFC 3339 format
- `date_updated` (String) The date and time the flow was last updated, in RFC 3339 format
- `definition` (String) A JSON string defining the flow structure
- `friendly_name` (String) The human-readable label for the flow
- `id` (String) The ID of this resource.
- `revision` (Number) The current revision number of the flow
- `status` (String) The lifecycle status of the flow. Either `draft` or `published`
- `url` (String) The absolute URL of the flow resource
- `valid` (Boolean) Whether the flow definition passed Twilio's validation checks
- `webhook_url` (String) The webhook URL used to trigger this flow via the REST API

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
