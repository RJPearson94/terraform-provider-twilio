---
page_title: "twilio_video_composition_hook Resource - twilio"
subcategory: "Video"
description: |-
  
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

## Schema

### Required

- `friendly_name` (String) A human-readable label for the video composition hook

### Optional

- `audio_sources` (List of String) A list of audio source track names to include in the composition
- `audio_sources_excluded` (List of String) A list of audio source track names to exclude from the composition
- `enabled` (Boolean) Whether the composition hook is enabled. Defaults to `true`
- `format` (String) The file format for the composition. Valid values are `mp4` or `webm`. Defaults to `webm`
- `resolution` (String) The resolution of the composition in the format `WIDTHxHEIGHT`. Defaults to `640x480`
- `status_callback_method` (String) The HTTP method used to call the status callback URL. Valid values are `GET` or `POST`. Defaults to `POST`
- `status_callback_url` (String) The URL to call for composition status callback events
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))
- `trim` (Boolean) Whether to remove intervals with no media from the composition. Defaults to `true`
- `video_layout` (String) A JSON string describing the video layout of the composition

### Read-Only

- `account_sid` (String) The SID of the account that owns this video composition hook
- `date_created` (String) The date and time the video composition hook was created, in RFC 3339 format
- `date_updated` (String) The date and time the video composition hook was last updated, in RFC 3339 format
- `id` (String) The ID of this resource.
- `sid` (String) The unique SID assigned to this video composition hook by Twilio
- `url` (String) The absolute URL of the video composition hook resource

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)

## Import

A composition hook can be imported using the `/CompositionHooks/{sid}` format, e.g.

```shell
terraform import twilio_video_composition_hook.composition_hook /CompositionHooks/HKXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
