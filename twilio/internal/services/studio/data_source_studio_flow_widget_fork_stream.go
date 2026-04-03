package studio

import (
	"context"
	"regexp"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetForkStream() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetForkStreamRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source",
			},
			"transitions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The next widget(s) to transition to after this widget",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget after the stream is started or stopped",
						},
					},
				},
			},
			"offset": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The position of this widget in the Studio visual editor",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The x-axis position. Defaults to 0",
						},
						"y": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The y-axis position. Defaults to 0",
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The unique name of this widget within the flow, used to reference it in transitions",
			},
			"stream_action": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"start",
					"stop",
				}, false),
				Description: "Whether to start or stop the audio stream. Valid values: `start`, `stop`",
			},
			"stream_connector": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the SIPREC connector when using `siprec` transport",
			},
			"stream_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A friendly name for the stream, used to reference it when stopping",
			},
			"stream_parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Additional key/value parameters to send to the remote stream service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The parameter name",
						},
						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The parameter value",
						},
					},
				},
			},
			"stream_track": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"both_tracks",
					"inbound_track",
					"outbound_track",
				}, false),
				Description: "Which audio track(s) to stream. Valid values: `both_tracks`, `inbound_track`, `outbound_track`",
			},
			"stream_transport_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"siprec",
					"websocket",
				}, false),
				Description: "The transport protocol for the stream. Valid values: `siprec`, `websocket`",
			},
			"stream_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringMatch(regexp.MustCompile(`^wss://.+$`), ""),
				),
				Description: "The WebSocket URL (`wss://`) to stream audio to when using `websocket` transport",
			},
		},
	}
}

func dataSourceStudioFlowWidgetForkStreamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.ForkStreamNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Next = utils.OptionalString(d, "transitions.0.next")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	var streamParameters *[]widgets.ForkStreamStreamParameter
	if v, ok := d.GetOk("stream_parameters"); ok {
		parameters := []widgets.ForkStreamStreamParameter{}
		for _, parameter := range v.([]interface{}) {
			parameterMap := parameter.(map[string]interface{})
			parameters = append(parameters, widgets.ForkStreamStreamParameter{
				Key:   parameterMap["key"].(string),
				Value: parameterMap["value"].(string),
			})
		}
		streamParameters = &parameters
	}

	widget := &widgets.ForkStream{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.ForkStreamProperties{
			Offset:              offset,
			StreamAction:        d.Get("stream_action").(string),
			StreamConnector:     utils.OptionalString(d, "stream_connector"),
			StreamName:          utils.OptionalString(d, "stream_name"),
			StreamParameters:    streamParameters,
			StreamTrack:         utils.OptionalString(d, "stream_track"),
			StreamTransportType: utils.OptionalString(d, "stream_transport_type"),
			StreamURL:           utils.OptionalString(d, "stream_url"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Fork stream widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create fork stream widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal fork stream widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
