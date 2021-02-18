---
page_title: "Twilio TaskRouter Activity"
subcategory: "TaskRouter"
---

# twilio_taskrouter_activity Data Source

Use this data source to access information about an existing TaskRouter activity. See the [API docs](https://www.twilio.com/docs/taskrouter/api/activity) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_activity" "activity" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "WAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "activity" {
  value = data.twilio_taskrouter_activity.activity
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace the activity is associated with
- `sid` - (Mandatory) The SID of the activity

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the activity (Same as the `sid`)
- `sid` - The SID of the activity (Same as the `id`)
- `account_sid` - The account SID of the activity is deployed into
- `workspace_sid` - The workspace SID to create the activity under.
- `friendly_name` - The name of the activity
- `available` - Whether the activity is available to accept tasks in Task Router
- `date_created` - The date in RFC3339 format that the activity was created
- `date_updated` - The date in RFC3339 format that the activity was updated
- `url` - The URL of the activity

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the activity
