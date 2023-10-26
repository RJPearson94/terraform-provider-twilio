package studio

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Studio"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_studio_flow":                                dataSourceStudioFlow(),
		"twilio_studio_flow_definition":                     dataSourceStudioFlowDefinition(),
		"twilio_studio_flow_widget_add_twiml_redirect":      dataSourceStudioFlowWidgetAddTwiMLRedirect(),
		"twilio_studio_flow_widget_capture_payments":        dataSourceStudioFlowWidgetCapturePayments(),
		"twilio_studio_flow_widget_connect_call_to":         dataSourceStudioFlowWidgetConnectCallTo(),
		"twilio_studio_flow_widget_connect_virtual_agent":   dataSourceStudioFlowWidgetConnectVirtualAgent(),
		"twilio_studio_flow_widget_enqueue_call":            dataSourceStudioFlowWidgetEnqueueCall(),
		"twilio_studio_flow_widget_fork_stream":             dataSourceStudioFlowWidgetForkStream(),
		"twilio_studio_flow_widget_gather_input_on_call":    dataSourceStudioFlowWidgetGatherInputOnCall(),
		"twilio_studio_flow_widget_make_http_request":       dataSourceStudioFlowWidgetMakeHttpRequest(),
		"twilio_studio_flow_widget_make_outgoing_call":      dataSourceStudioFlowWidgetMakeOutgoingCall(),
		"twilio_studio_flow_widget_record_call":             dataSourceStudioFlowWidgetRecordCall(),
		"twilio_studio_flow_widget_record_voicemail":        dataSourceStudioFlowWidgetRecordVoicemail(),
		"twilio_studio_flow_widget_run_function":            dataSourceStudioFlowWidgetRunFunction(),
		"twilio_studio_flow_widget_say_play":                dataSourceStudioFlowWidgetSayPlay(),
		"twilio_studio_flow_widget_send_and_wait_for_reply": dataSourceStudioFlowWidgetSendAndWaitForReply(),
		"twilio_studio_flow_widget_send_message":            dataSourceStudioFlowWidgetSendMessage(),
		"twilio_studio_flow_widget_send_to_flex":            dataSourceStudioFlowWidgetSendToFlex(),
		"twilio_studio_flow_widget_set_variables":           dataSourceStudioFlowWidgetSetVariables(),
		"twilio_studio_flow_widget_split_based_on":          dataSourceStudioFlowWidgetSplitBasedOn(),
		"twilio_studio_flow_widget_state":                   dataSourceStudioFlowWidgetState(),
		"twilio_studio_flow_widget_trigger":                 dataSourceStudioFlowWidgetTrigger(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_studio_flow": resourceStudioFlow(),
	}
}
