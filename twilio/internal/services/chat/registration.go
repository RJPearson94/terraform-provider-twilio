package chat

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Chat"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_chat_service":                 resourceChatService(),
		"twilio_chat_role":                    resourceChatRole(),
		"twilio_chat_channel":                 resourceChatChannel(),
		"twilio_chat_channel_webhook":         resourceChatChannelWebhook(),
		"twilio_chat_channel_studio_webhook":  resourceChatChannelStudioWebhook(),
		"twilio_chat_channel_trigger_webhook": resourceChatChannelTriggerWebhook(),
	}
}
