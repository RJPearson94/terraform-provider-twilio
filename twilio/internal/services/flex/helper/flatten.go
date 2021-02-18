package helper

import (
	"time"

	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/flex_flow"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configuration/plugins"
)

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

func FlattenPlugins(plugins []plugins.PagePluginResponse) *[]interface{} {
	results := make([]interface{}, 0)
	for _, prop := range plugins {
		results = append(results, map[string]interface{}{
			"plugin_version_sid": prop.PluginVersionSid,
			"plugin_sid":         prop.PluginSid,
			"plugin_url":         prop.PluginURL,
			"phase":              prop.Phase,
			"private":            prop.Private,
			"unique_name":        prop.UniqueName,
			"version":            prop.Version,
			"date_created":       prop.DateCreated.Format(time.RFC3339),
			"url":                prop.URL,
		})
	}
	return &results
}
