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
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"next": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"offset": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"y": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"stream_action": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"start",
					"stop",
				}, false),
			},
			"stream_connector": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stream_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
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
			},
			"stream_transport_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"siprec",
					"websocket",
				}, false),
			},
			"stream_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringMatch(regexp.MustCompile(`^wss://.+$`), ""),
				),
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
