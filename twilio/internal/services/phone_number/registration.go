package phone_number

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Phone Number"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_phone_number":                             dataSourcePhoneNumber(),
		"twilio_phone_number_available_local_numbers":     dataSourcePhoneNumberAvailableLocalNumbers(),
		"twilio_phone_number_available_mobile_numbers":    dataSourcePhoneNumberAvailableMobileNumbers(),
		"twilio_phone_number_available_toll_free_numbers": dataSourcePhoneNumberAvailableTollFreeNumbers(),
		"twilio_phone_numbers":                            dataSourcePhoneNumbers(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_phone_number": resourcePhoneNumber(),
	}
}
