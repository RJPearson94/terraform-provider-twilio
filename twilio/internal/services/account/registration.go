package account

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Account"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_account_balance": dataSourceAccountBalance(),
		"twilio_account_details": dataSourceAccountDetails(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_account_sub_account": resourceAccountSubAccount(),
	}
}
