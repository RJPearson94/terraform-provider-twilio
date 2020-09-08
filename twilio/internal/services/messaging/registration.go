package messaging

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Messaging"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_messaging_service":       dataSourceMessagingService(),
		"twilio_messaging_phone_number":  dataSourceMessagingPhoneNumber(),
		"twilio_messaging_phone_numbers": dataSourceMessagingPhoneNumbers(),
		"twilio_messaging_short_code":    dataSourceMessagingShortCode(),
		"twilio_messaging_short_codes":   dataSourceMessagingShortCodes(),
		"twilio_messaging_alpha_sender":  dataSourceMessagingAlphaSender(),
		"twilio_messaging_alpha_senders": dataSourceMessagingAlphaSenders(),
	}
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
