package autopilot

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Autopilot"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_autopilot_assistant":    dataSourceAutopilotAssistant(),
		"twilio_autopilot_field_type":   dataSourceAutopilotFieldType(),
		"twilio_autopilot_field_types":  dataSourceAutopilotFieldTypes(),
		"twilio_autopilot_field_value":  dataSourceAutopilotFieldValue(),
		"twilio_autopilot_field_values": dataSourceAutopilotFieldValues(),
		"twilio_autopilot_model_build":  dataSourceAutopilotModelBuild(),
		"twilio_autopilot_model_builds": dataSourceAutopilotModelBuilds(),
		"twilio_autopilot_task":         dataSourceAutopilotTask(),
		"twilio_autopilot_tasks":        dataSourceAutopilotTasks(),
		"twilio_autopilot_task_field":   dataSourceAutopilotTaskField(),
		"twilio_autopilot_task_fields":  dataSourceAutopilotTaskFields(),
		"twilio_autopilot_task_sample":  dataSourceAutopilotTaskSample(),
		"twilio_autopilot_task_samples": dataSourceAutopilotTaskSamples(),
		"twilio_autopilot_webhook":      dataSourceAutopilotWebhook(),
		"twilio_autopilot_webhooks":     dataSourceAutopilotWebhooks(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_autopilot_assistant":   resourceAutopilotAssistant(),
		"twilio_autopilot_field_type":  resourceAutopilotFieldType(),
		"twilio_autopilot_field_value": resourceAutopilotFieldValue(),
		"twilio_autopilot_model_build": resourceAutopilotModelBuild(),
		"twilio_autopilot_task":        resourceAutopilotTask(),
		"twilio_autopilot_task_field":  resourceAutopilotTaskField(),
		"twilio_autopilot_task_sample": resourceAutopilotTaskSample(),
		"twilio_autopilot_webhook":     resourceAutopilotWebhook(),
	}
}
