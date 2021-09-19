package taskrouter

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "TaskRouter"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_taskrouter_activities":    dataSourceTaskRouterActivities(),
		"twilio_taskrouter_activity":      dataSourceTaskRouterActivity(),
		"twilio_taskrouter_task_channel":  dataSourceTaskRouterTaskChannel(),
		"twilio_taskrouter_task_channels": dataSourceTaskRouterTaskChannels(),
		"twilio_taskrouter_task_queue":    dataSourceTaskRouterTaskQueue(),
		"twilio_taskrouter_task_queues":   dataSourceTaskRouterTaskQueues(),
		"twilio_taskrouter_worker":        dataSourceTaskRouterWorker(),
		"twilio_taskrouter_workers":       dataSourceTaskRouterWorkers(),
		"twilio_taskrouter_workflow":      dataSourceTaskRouterWorkflow(),
		"twilio_taskrouter_workflows":     dataSourceTaskRouterWorkflows(),
		"twilio_taskrouter_workspace":     dataSourceTaskRouterWorkspace(),
		"twilio_taskrouter_workspaces":    dataSourceTaskRouterWorkspaces(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_taskrouter_activity":                resourceTaskRouterActivity(),
		"twilio_taskrouter_task_channel":            resourceTaskRouterTaskChannel(),
		"twilio_taskrouter_task_queue":              resourceTaskRouterTaskQueue(),
		"twilio_taskrouter_worker":                  resourceTaskRouterWorker(),
		"twilio_taskrouter_workflow":                resourceTaskRouterWorkflow(),
		"twilio_taskrouter_workspace":               resourceTaskRouterWorkspace(),
		"twilio_taskrouter_workspace_configuration": resourceTaskRouterWorkspaceConfiguration(),
	}
}
