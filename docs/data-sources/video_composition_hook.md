---
page_title: "Twilio Video Composition Hook"
subcategory: "Video"
---

# twilio_video_composition_hook Data Source

Use this data source to access information about an existing composition hook. See the [API docs](https://www.twilio.com/docs/video/api/composition-hooks) for more information

## Example Usage

```hcl
data "twilio_video_composition_hook" "composition_hook" {
  sid = "HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

output "composition_hook" {
  value = data.twilio_video_composition_hook.composition_hook
}
```

## Argument Reference

The following arguments are supported:

- `sid` - (Mandatory) The SID of the composition hook

## Attributes Reference

The following attributes are exported:

- `id` - The ID of the composition hook (Same as the `sid`)
- `sid` - The SID of the composition hook (Same as the `id`)
- `account_sid` - The account SID the composition hook is associated with
- `friendly_name` - The friendly name of the composition hook
- `audio_sources` - A list of audio sources to include in the compositions
- `audio_sources_excluded` - A list of audio sources to exclude from the compositions
- `enabled` - Whether the composition hook is enabled
- `format` - The media file format of the compositions
- `resolution` - The pixel dimensions for the video
- `status_callback_url` - The URL to call on each status change
- `status_callback_method` - The HTTP method which should be used to call the status callback URL
- `trim` - Whether sections with no audio or media content should be trimmed from the composition
- `video_layout` - JSON string of the composition video layout
- `date_created` - The date in RFC3339 format that the composition hook was created
- `date_updated` - The date in RFC3339 format that the composition hook was updated
- `url` - The URL of the composition hook

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `read` - (Defaults to 5 minutes) Used when retrieving the composition hook details
