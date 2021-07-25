package video

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Video"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_video_composition_hook":     dataSourceVideoCompositionHook(),
		"twilio_video_composition_settings": dataSourceVideoCompositionSettings(),
		"twilio_video_recording_settings":   dataSourceVideoRecordingSettings(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_video_composition_hook":     resourceVideoCompositionHook(),
		"twilio_video_composition_settings": resourceVideoCompositionSettings(),
		"twilio_video_recording_settings":   resourceVideoRecordingSettings(),
	}
}
