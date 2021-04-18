package twilio

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
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
				Description: "The Account SID which should be used.",
			},
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_AUTH_TOKEN", nil),
				Description: "The Auth Token which should be used.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_API_KEY", nil),
				Description: "The API Key which should be used.",
			},
			"api_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("TWILIO_API_SECRET", nil),
				Description: "The API Key secret which should be used.",
			},
		},

		DataSourcesMap: dataSources,
		ResourcesMap:   resources,
	}

	provider.ConfigureContextFunc = providerConfigure(provider)

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

func providerConfigure(p *schema.Provider) schema.ConfigureContextFunc {
	return func(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		terraformVersion := p.TerraformVersion

		config := Config{
			AccountSid:       d.Get("account_sid").(string),
			AuthToken:        d.Get("auth_token").(string),
			APIKey:           d.Get("api_key").(string),
			APISecret:        d.Get("api_secret").(string),
			terraformVersion: terraformVersion,
		}

		return config.Client()
	}
}
