package messaging

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Messaging"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_messaging_service":      resourceMessagingService(),
		"twilio_messaging_phone_number": resourceMessagingPhoneNumber(),
		"twilio_messaging_short_code":   resourceMessagingShortCode(),
		"twilio_messaging_alpha_sender": resourceMessagingAlphaSender(),
	}
}
