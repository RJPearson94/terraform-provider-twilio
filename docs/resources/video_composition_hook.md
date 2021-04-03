---
page_title: "Twilio Video Composition Hook"
subcategory: "Video"
---

# twilio_video_composition_hook Resource

Manages a Programmable Video composition hook. See the [API docs](https://www.twilio.com/docs/video/api/composition-hooks) for more information

## Example Usage

```hcl
resource "twilio_video_composition_hook" "composition_hook" {
  friendly_name = "Test Composition Hook"
  audio_sources = ["*"]
  format        = "mp4"
}
```

## Argument Reference

The following arguments are supported:

- `friendly_name` - (Mandatory) The friendly name of the composition hook. The value cannot be an empty string
- `audio_sources` - (Optional) A list of audio sources to include in the compositions
- `audio_sources_excluded` - (Optional) A List of audio sources to exclude from the compositions
- `enabled` - (Optional) Whether the composition hook is enabled. The default value is `true`
- `format` - (Optional) The media file format of the compositions. Valid values are `mp4` or `webm`. The default value is `webm`
- `resolution` - (Optional) The pixel dimensions for the video. The value must be in the format `{height in pixels}x{width in pixels}`. The default value is `640x480`
- `status_callback_url` - (Optional) The URL to call on each status change
- `status_callback_method` - (Optional) The HTTP method should be used to call the status callback URL. Valid values are `GET` or `POST`. The default value is `POST`
- `trim` - (Optional) Whether sections with no audio or media content should be trimmed from the composition. The default value is `true`
- `video_layout` - (Optional) JSON string of the composition video layout

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

- `create` - (Defaults to 10 minutes) Used when creating the composition hook
- `update` - (Defaults to 10 minutes) Used when updating the composition hook
- `read` - (Defaults to 5 minutes) Used when retrieving the composition hook
- `delete` - (Defaults to 10 minutes) Used when deleting the composition hook

## Import

A composition hook can be imported using the `/CompositionHooks/{sid}` format, e.g.

```shell
terraform import twilio_video_composition_hook.composition_hook /CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
