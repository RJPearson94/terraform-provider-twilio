package chat

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Chat"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_chat_channel":          dataSourceChatChannel(),
		"twilio_chat_channel_member":   dataSourceChatChannelMember(),
		"twilio_chat_channel_members":  dataSourceChatChannelMembers(),
		"twilio_chat_channel_webhook":  dataSourceChatChannelWebhook(),
		"twilio_chat_channel_webhooks": dataSourceChatChannelWebhooks(),
		"twilio_chat_channels":         dataSourceChatChannels(),
		"twilio_chat_role":             dataSourceChatRole(),
		"twilio_chat_roles":            dataSourceChatRoles(),
		"twilio_chat_service":          dataSourceChatService(),
		"twilio_chat_user":             dataSourceChatUser(),
		"twilio_chat_users":            dataSourceChatUsers(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_chat_channel":                 resourceChatChannel(),
		"twilio_chat_channel_member":          resourceChatChannelMember(),
		"twilio_chat_channel_webhook":         resourceChatChannelWebhook(),
		"twilio_chat_channel_studio_webhook":  resourceChatChannelStudioWebhook(),
		"twilio_chat_channel_trigger_webhook": resourceChatChannelTriggerWebhook(),
		"twilio_chat_role":                    resourceChatRole(),
		"twilio_chat_service":                 resourceChatService(),
		"twilio_chat_user":                    resourceChatUser(),
	}
}
