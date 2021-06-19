package twiml

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "TwiML"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_twiml_app":  dataSourceTwimlApp(),
		"twilio_twiml_apps": dataSourceTwimlApps(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_twiml_app": resourceTwimlApp(),
	}
}
