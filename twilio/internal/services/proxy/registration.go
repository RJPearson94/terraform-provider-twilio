package proxy

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Proxy"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_proxy_service":       dataSourceProxyService(),
		"twilio_proxy_phone_number":  dataSourceProxyPhoneNumber(),
		"twilio_proxy_phone_numbers": dataSourceProxyPhoneNumbers(),
		"twilio_proxy_short_code":    dataSourceProxyShortCode(),
		"twilio_proxy_short_codes":   dataSourceProxyShortCodes(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_proxy_service":      resourceProxyService(),
		"twilio_proxy_phone_number": resourceProxyPhoneNumber(),
		"twilio_proxy_short_code":   resourceProxyShortCode(),
	}
}
