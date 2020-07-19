package autopilot

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Autopilot"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_autopilot_assistant":   resourceAutopilotAssistant(),
		"twilio_autopilot_webhook":     resourceAutopilotWebhook(),
		"twilio_autopilot_task":        resourceAutopilotTask(),
		"twilio_autopilot_task_sample": resourceAutopilotTaskSample(),
		"twilio_autopilot_task_field":  resourceAutopilotTaskField(),
		"twilio_autopilot_field_type":  resourceAutopilotFieldType(),
	}
}
