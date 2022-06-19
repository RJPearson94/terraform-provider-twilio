package verify

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Verify"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_verify_messaging_configuration":    dataSourceVerifyMessagingConfiguration(),
		"twilio_verify_messaging_configurations":   dataSourceVerifyMessagingConfigurations(),
		"twilio_verify_service":                    dataSourceVerifyService(),
		"twilio_verify_service_rate_limit":         dataSourceVerifyServiceRateLimit(),
		"twilio_verify_service_rate_limits":        dataSourceVerifyServiceRateLimits(),
		"twilio_verify_service_rate_limit_bucket":  dataSourceVerifyServiceRateLimitBucket(),
		"twilio_verify_service_rate_limit_buckets": dataSourceVerifyServiceRateLimitBuckets(),
		"twilio_verify_webhook":                    dataSourceVerifyWebhook(),
		"twilio_verify_webhooks":                   dataSourceVerifyWebhooks(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_verify_messaging_configuration":   resourceVerifyMessagingConfiguration(),
		"twilio_verify_service":                   resourceVerifyService(),
		"twilio_verify_service_rate_limit":        resourceVerifyServiceRateLimit(),
		"twilio_verify_service_rate_limit_bucket": resourceVerifyServiceRateLimitBucket(),
		"twilio_verify_webhook":                   resourceVerifyWebhook(),
	}
}
