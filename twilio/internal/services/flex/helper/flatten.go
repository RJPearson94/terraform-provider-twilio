package helper

import "github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"

func FlattenIntegration(integration *flex_flow.FetchFlexFlowIntegrationResponse) *[]interface{} {
	if integration == nil {
		return nil
	}

	return &[]interface{}{
		map[string]interface{}{
			"channel":             integration.Channel,
			"creation_on_message": integration.CreationOnMessage,
			"flow_sid":            integration.FlowSid,
			"priority":            integration.Priority,
			"retry_count":         integration.RetryCount,
			"timeout":             integration.Timeout,
			"url":                 integration.URL,
			"workspace_sid":       integration.WorkspaceSid,
		},
	}
}
