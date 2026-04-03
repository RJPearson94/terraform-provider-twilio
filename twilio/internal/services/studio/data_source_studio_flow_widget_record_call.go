package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetRecordCall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetRecordCallRead,

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
						"failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the recording operation fails",
						},
						"success": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the recording operation succeeds",
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
			"record_call": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to enable call recording. Defaults to false",
			},
			"trim": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"trim-silence",
						"do-not-trim",
					}, false),
				),
				Description: "Whether to trim silence from the recording. Valid values: `trim-silence`, `do-not-trim`",
			},
			"recording_status_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The HTTP/HTTPS URL to receive recording status callback events",
			},
			"recording_status_callback_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"GET",
						"POST",
					}, false),
				),
				Description: "The HTTP method for the recording status callback. Valid values: `GET`, `POST`",
			},
			"recording_channels": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"dual",
						"mono",
					}, false),
				),
				Description: "The number of recording channels. Valid values: `dual`, `mono`",
			},
			"recording_status_callback_events": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The recording events that trigger a callback. Valid values: `absent`, `completed`, `in-progress`",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"absent",
						"completed",
						"in-progress",
					}, false),
				},
			},
		},
	}
}

func dataSourceStudioFlowWidgetRecordCallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.RecordCallNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Failed = utils.OptionalString(d, "transitions.0.failed")
		nextTransitions.Success = utils.OptionalString(d, "transitions.0.success")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.RecordCall{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.RecordCallProperties{
			RecordCall:                    d.Get("record_call").(bool),
			Trim:                          utils.OptionalString(d, "trim"),
			RecordingStatusCallbackURL:    utils.OptionalString(d, "recording_status_callback_url"),
			RecordingStatusCallbackMethod: utils.OptionalString(d, "recording_status_callback_method"),
			RecordingChannels:             utils.OptionalString(d, "recording_channels"),
			RecordingStatusCallbackEvents: utils.OptionalSeperatedString(d, "recording_status_callback_events", " "),
			Offset:                        offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Record call widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create record call widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal record call widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
