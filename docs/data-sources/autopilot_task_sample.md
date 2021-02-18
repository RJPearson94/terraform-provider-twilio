---
page_title: "Twilio Autopilot Task Sample"
subcategory: "Autopilot"
---

# twilio_autopilot_task_sample Data Source

Use this data source to access information about an existing Autopilot task sample. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-sample) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_task_sample" "task_sample" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  task_sid      = "UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  sid           = "UEXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_sample" {
  value = data.twilio_autopilot_task_sample.task_sample
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the sample is associated with
- `task_sid` - (Mandatory) The SID of the task the sample is associated with
- `sid` - (Mandatory) The SID of the sample

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the sample (Same as the `sid`)
- `sid` - The SID of the sample (Same as the `id`)
- `account_sid` - The account SID associated with the sample
- `assistant_sid` - The SID of the assistant to attach the sample to
- `task_sid` - The SID of the task to attach the sample to
- `language` - The language of the sample
- `tagged_text` - The labelled/ tagged sample text
- `source_channel` - The channel the sample was captured on
- `date_created` - The date in RFC3339 format that the sample was created
- `date_updated` - The date in RFC3339 format that the sample was updated
- `url` - The URL of the sample

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the sample
