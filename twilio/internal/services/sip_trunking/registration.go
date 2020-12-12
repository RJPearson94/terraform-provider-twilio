package sip_trunking

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "SIP Trunking"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_sip_trunking_origination_url":  dataSourceSIPTrunkingOriginationURL(),
		"twilio_sip_trunking_origination_urls": dataSourceSIPTrunkingOriginationURLs(),
		"twilio_sip_trunking_phone_number":     dataSourceSIPTrunkingPhoneNumber(),
		"twilio_sip_trunking_phone_numbers":    dataSourceSIPTrunkingPhoneNumbers(),
		"twilio_sip_trunking_trunk":            dataSourceSIPTrunkingTrunk(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_sip_trunking_origination_url": resourceSIPTrunkingOriginationURL(),
		"twilio_sip_trunking_phone_number":    resourceSIPTrunkingPhoneNumber(),
		"twilio_sip_trunking_trunk":           resourceSIPTrunkingTrunk(),
	}
}
