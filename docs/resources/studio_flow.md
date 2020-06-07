# twilio_studio_flow

Manages a Studio Flow

**NOTE:** This resource is in beta

## Example Usage

```hcl
resource "twilio_studio_flow" "flow" {
  friendly_name = "Test Studio Flow"
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

* `friendly_name` - (Mandatory) The name of the Studio Flow
* `status` - (Mandatory) The status of the Studio Flow. Valid values include `draft` and `published`
* `definition` - (Mandatory) The Flow Definition JSON
* `validate` - (Optional) Whether to validate the Flow Definition JSON before creating a new revision. The default is false
* `commit_message` - (Optional) Description of the changes made

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Studio Flow (Same as the SID)
* `sid` - The SID of the Studio Flow (Same as the ID)
* `friendly_name` - The name of the Studio Flow
* `definition` - The Flow Definition JSON
* `status` -  The status of the Studio Flow
* `revision` - The revision number of the Studio Flow
* `valid` -  Whether the Studio Flow is valid
* `validate` -  Whether the Studio Flow has been validated on creation and on updates
* `date_created` - The date that the Studio Flow was created
* `date_updated` - The date that the Studio Flow was updated
* `url` - The url of the Studio Flow
* `webhook_url` - The webhook url of the Studio Flow
