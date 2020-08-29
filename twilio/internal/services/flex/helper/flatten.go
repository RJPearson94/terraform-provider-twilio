package helper

import "github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"

func FlattenIntegration(integration *flex_flow.FetchFlexFlowResponseIntegration) *[]interface{} {
	if integration == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["channel"] = integration.Channel
	result["creation_on_message"] = integration.CreationOnMessage
	result["flow_sid"] = integration.FlowSid
	result["priority"] = integration.Priority
	result["retry_count"] = integration.RetryCount
	result["timeout"] = integration.Timeout
	result["url"] = integration.URL
	result["workspace_sid"] = integration.WorkspaceSid

	results = append(results, result)
	return &results
}
