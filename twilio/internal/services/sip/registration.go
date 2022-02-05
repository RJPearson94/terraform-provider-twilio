package sip

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "SIP"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_sip_credential":                                   dataSourceSIPCredential(),
		"twilio_sip_credentials":                                  dataSourceSIPCredentials(),
		"twilio_sip_credential_list":                              dataSourceSIPCredentialList(),
		"twilio_sip_domain":                                       dataSourceSIPDomain(),
		"twilio_sip_domain_credential_list_mapping":               dataSourceSIPDomainCredentialListMapping(),
		"twilio_sip_domain_credential_list_mappings":              dataSourceSIPDomainCredentialListMappings(),
		"twilio_sip_domain_ip_access_control_list_mapping":        dataSourceSIPDomainIPAccessControlListMapping(),
		"twilio_sip_domain_ip_access_control_list_mappings":       dataSourceSIPDomainIPAccessControlListMappings(),
		"twilio_sip_domain_registration_credential_list_mapping":  dataSourceSIPDomainRegistrationCredentialListMapping(),
		"twilio_sip_domain_registration_credential_list_mappings": dataSourceSIPDomainRegistrationCredentialListMappings(),
		"twilio_sip_ip_access_control_list":                       dataSourceSIPIPAccessControlList(),
		"twilio_sip_ip_address":                                   dataSourceSIPIPAddress(),
		"twilio_sip_ip_addresses":                                 dataSourceSIPIPAddresses(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_sip_credential":                                  resourceSIPCredential(),
		"twilio_sip_credential_list":                             resourceSIPCredentialList(),
		"twilio_sip_domain":                                      resourceSIPDomain(),
		"twilio_sip_domain_credential_list_mapping":              resourceSIPDomainCredentialListMapping(),
		"twilio_sip_domain_ip_access_control_list_mapping":       resourceSIPDomainIPAccessControlListMapping(),
		"twilio_sip_domain_registration_credential_list_mapping": resourceSIPDomainRegistrationCredentialListMapping(),
		"twilio_sip_ip_access_control_list":                      resourceSIPIPAccessControlList(),
		"twilio_sip_ip_address":                                  resourceSIPIPAddress(),
	}
}
