---
page_title: "Twilio Autopilot Task Samples"
subcategory: "Autopilot"
---

# twilio_autopilot_task_samples Data Source

Use this data source to access information about the task samples associated with an existing Autopilot assistant and task. See the [API docs](https://www.twilio.com/docs/autopilot/api/task-sample) for more information

For more information on Autopilot, see the product [page](https://www.twilio.com/autopilot)

## Example Usage

```hcl
data "twilio_autopilot_task_samples" "task_samples" {
  assistant_sid = "UAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
  task_sid      = "UDXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "task_samples" {
  value = data.twilio_autopilot_task_samples.task_samples
}
```

## Argument Reference

The following arguments are supported:

- `assistant_sid` - (Mandatory) The SID of the assistant the samples are associated with
- `task_sid` - (Mandatory) The SID of the task the samples are associated with
- `language` - (Optional) Search for all samples which have the language specified

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the resource in the format `assistant_sid/task_sid`
- `account_sid` - The SID of the account the samples are associated with
- `assistant_sid` - The SID of the assistant the samples are associated with
- `task_sid` - The SID of the task the samples are associated with
- `samples` - A list of `sample` blocks as documented below

---

A `sample` block supports the following:

- `sid` - The SID of the sample
- `language` - The language of the sample
- `tagged_text` - The labelled/ tagged sample text
- `source_channel` - The channel the sample was captured on
- `date_created` - The date in RFC3339 format that the sample was created
- `date_updated` - The date in RFC3339 format that the sample was updated
- `url` - The URL of the sample

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 10 minutes) Used when retrieving samples
