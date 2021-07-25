package credentials

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Credentials"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_credentials_aws":        resourceCredentialsAWS(),
		"twilio_credentials_public_key": resourceCredentialsPublicKey(),
	}
}
