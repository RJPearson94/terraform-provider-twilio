package taskrouter

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "TaskRouter"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_taskrouter_workspace":  resourceTaskRouterWorkspace(),
		"twilio_taskrouter_activity":   resourceTaskRouterActivity(),
		"twilio_taskrouter_task_queue": resourceTaskRouterTaskQueue(),
		"twilio_taskrouter_worker":     resourceTaskRouterWorker(),
		"twilio_taskrouter_workflow":   resourceTaskRouterWorkflow(),
	}
}
