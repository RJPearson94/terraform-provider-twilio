package twilio

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	dataSources := make(map[string]*schema.Resource)
	resources := make(map[string]*schema.Resource)

	for _, service := range SupportedServices() {
		serviceName := service.Name()
		validateAndRegisterSupportedResources(dataSources, service.SupportedDataSources(), serviceName, "Data Sources")
		validateAndRegisterSupportedResources(resources, service.SupportedResources(), serviceName, "Resources")
	}

	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_ACCOUNT_SID", nil),
				Description: "The Account Sid which should be used.",
			},
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "The Auth Token which should be used.",
			},
		},

		DataSourcesMap: dataSources,
		ResourcesMap:   resources,
	}

	provider.ConfigureFunc = providerConfigure(provider)

	return provider
}

func validateAndRegisterSupportedResources(registeredResources map[string]*schema.Resource, resourcesToAdd map[string]*schema.Resource, serviceName string, resourceType string) {
	log.Printf("[DEBUG] Registering %s for %q..", resourceType, serviceName)
	for key, value := range resourcesToAdd {
		if existing := registeredResources[key]; existing != nil {
			//lintignore:R009
			panic(fmt.Sprintf("An existing %s exists for %q", resourceType, key))
		}

		registeredResources[key] = value
	}
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}

		config := Config{
			AccountSid:       d.Get("account_sid").(string),
			AuthToken:        d.Get("auth_token").(string),
			terraformVersion: terraformVersion,
		}

		return config.Client()
	}
}
