---
page_title: "twilio_video_composition_hook Data Source - twilio"
subcategory: "Video"
description: |-
  
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

## Schema

### Required

- `sid` (String) The SID of the video composition hook to look up

### Optional

- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `account_sid` (String) The SID of the account that owns this video composition hook
- `audio_sources` (List of String) A list of audio source track names included in the composition
- `audio_sources_excluded` (List of String) A list of audio source track names excluded from the composition
- `date_created` (String) The date and time the video composition hook was created, in RFC 3339 format
- `date_updated` (String) The date and time the video composition hook was last updated, in RFC 3339 format
- `enabled` (Boolean) Whether the composition hook is enabled
- `format` (String) The file format for the composition
- `friendly_name` (String) A human-readable label for the video composition hook
- `id` (String) The ID of this resource.
- `resolution` (String) The resolution of the composition in the format `WIDTHxHEIGHT`
- `status_callback_method` (String) The HTTP method used to call the status callback URL
- `status_callback_url` (String) The URL called for composition status callback events
- `trim` (Boolean) Whether intervals with no media are removed from the composition
- `url` (String) The absolute URL of the video composition hook resource
- `video_layout` (String) A JSON string describing the video layout of the composition

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `read` (String)
