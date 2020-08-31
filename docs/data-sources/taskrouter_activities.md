---
page_title: "Twilio TaskRouter Activities"
subcategory: "TaskRouter"
---

# twilio_taskrouter_activities Data Source

Use this data source to access information about the activities associated with an existing TaskRouter workspace. See the [API docs](https://www.twilio.com/docs/taskrouter/api/activity) for more information

For more information on TaskRouter, see the product [page](https://www.twilio.com/taskrouter)

## Example Usage

```hcl
data "twilio_taskrouter_activities" "activities" {
  workspace_sid = "WSXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "activities" {
  value = data.twilio_taskrouter_activities.activities
}
```

## Argument Reference

The following arguments are supported:

- `workspace_sid` - (Mandatory) The SID of the workspace activities are associated with

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource (Same as the workspace SID)
- `workspace_sid` - The SID of the workspace the activities are associated with
- `account_sid` - The account SID associated with the activities
- `activities` - A list of `activity` blocks as documented below

---

An `activity` block supports the following:

- `sid` - The SID of the activity
- `friendly_name` - The name of the activity
- `available` - Whether the activity is available to accept tasks in Task Router
- `date_created` - The date in RFC3339 format that the activity was created
- `date_updated` - The date in RFC3339 format that the activity was updated
- `url` - The URL of the activity

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving the activities
