package serverless

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Serverless"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_serverless_asset":        dataSourceServerlessAsset(),
		"twilio_serverless_assets":       dataSourceServerlessAssets(),
		"twilio_serverless_build":        dataSourceServerlessBuild(),
		"twilio_serverless_builds":       dataSourceServerlessBuilds(),
		"twilio_serverless_deployment":   dataSourceServerlessDeployment(),
		"twilio_serverless_deployments":  dataSourceServerlessDeployments(),
		"twilio_serverless_environment":  dataSourceServerlessEnvironment(),
		"twilio_serverless_environments": dataSourceServerlessEnvironments(),
		"twilio_serverless_function":     dataSourceServerlessFunction(),
		"twilio_serverless_functions":    dataSourceServerlessFunctions(),
		"twilio_serverless_service":      dataSourceServerlessService(),
		"twilio_serverless_variable":     dataSourceServerlessVariable(),
		"twilio_serverless_variables":    dataSourceServerlessVariables(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_serverless_asset":       resourceServerlessAsset(),
		"twilio_serverless_build":       resourceServerlessBuild(),
		"twilio_serverless_deployment":  resourceServerlessDeployment(),
		"twilio_serverless_environment": resourceServerlessEnvironment(),
		"twilio_serverless_function":    resourceServerlessFunction(),
		"twilio_serverless_service":     resourceServerlessService(),
		"twilio_serverless_variable":    resourceServerlessVariable(),
	}
}
