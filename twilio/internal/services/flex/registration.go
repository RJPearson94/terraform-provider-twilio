package flex

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Flex"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_flex_flow":                 dataSourceFlexFlow(),
		"twilio_flex_plugin":               dataSourceFlexPlugin(),
		"twilio_flex_plugin_configuration": dataSourceFlexPluginConfiguration(),
		"twilio_flex_plugin_release":       dataSourceFlexPluginRelease(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_flex_flow":                 resourceFlexFlow(),
		"twilio_flex_plugin":               resourceFlexPlugin(),
		"twilio_flex_plugin_configuration": resourceFlexPluginConfiguration(),
		"twilio_flex_plugin_release":       resourceFlexPluginRelease(),
	}
}
