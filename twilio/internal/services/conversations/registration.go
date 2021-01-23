package conversations

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Conversations"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_conversations_configuration":         dataSourceConversationsConfiguration(),
		"twilio_conversations_conversation_webhook":  dataSourceConversationsConversationWebhook(),
		"twilio_conversations_conversation_webhooks": dataSourceConversationsConversationWebhooks(),
		"twilio_conversations_conversation":          dataSourceConversationsConversation(),
		"twilio_conversations_conversations":         dataSourceConversationsConversations(),
		"twilio_conversations_role":                  dataSourceConversationsRole(),
		"twilio_conversations_roles":                 dataSourceConversationsRoles(),
		"twilio_conversations_service_configuration": dataSourceConversationsServiceConfiguration(),
		"twilio_conversations_service_notification":  dataSourceConversationsServiceNotification(),
		"twilio_conversations_service":               dataSourceConversationsService(),
		"twilio_conversations_user":                  dataSourceConversationsUser(),
		"twilio_conversations_users":                 dataSourceConversationsUsers(),
		"twilio_conversations_webhook":               dataSourceConversationsWebhook(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_conversations_configuration":                resourceConversationsConfiguration(),
		"twilio_conversations_conversation_studio_webhook":  resourceConversationsConversationStudioWebhook(),
		"twilio_conversations_conversation_trigger_webhook": resourceConversationsConversationTriggerWebhook(),
		"twilio_conversations_conversation_webhook":         resourceConversationsConversationWebhook(),
		"twilio_conversations_conversation":                 resourceConversationsConversation(),
		"twilio_conversations_push_credential_apn":          resourceConversationsPushCredentialAPN(),
		"twilio_conversations_push_credential_fcm":          resourceConversationsPushCredentialFCM(),
		"twilio_conversations_role":                         resourceConversationsRole(),
		"twilio_conversations_service_configuration":        resourceConversationsServiceConfiguration(),
		"twilio_conversations_service_notification":         resourceConversationsServiceNotification(),
		"twilio_conversations_service":                      resourceConversationsService(),
		"twilio_conversations_user":                         resourceConversationsUser(),
		"twilio_conversations_webhook":                      resourceConversationsWebhook(),
	}
}
