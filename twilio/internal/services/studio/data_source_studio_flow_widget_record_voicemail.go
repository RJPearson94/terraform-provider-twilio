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

func dataSourceStudioFlowWidgetRecordVoicemail() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetRecordVoicemailRead,

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
						"hangup": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the caller hangs up during recording",
						},
						"no_audio": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when no audio is detected during recording",
						},
						"recording_complete": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the voicemail recording completes successfully",
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
			"transcribe": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to transcribe the voicemail recording",
			},
			"transcription_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The HTTP/HTTPS URL to receive the transcription result",
			},
			"play_beep": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"true",
						"false",
					}, false),
				),
				Description: "Whether to play a beep before starting the recording. Valid values: `true`, `false`",
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
			"finish_on_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DTMF key that ends the recording (e.g. `#`)",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds of silence before the recording automatically stops",
			},
			"max_length": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 14400),
				Description:  "The maximum recording length in seconds (1–14400)",
			},
		},
	}
}

func dataSourceStudioFlowWidgetRecordVoicemailRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.RecordVoicemailNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Hangup = utils.OptionalString(d, "transitions.0.hangup")
		nextTransitions.NoAudio = utils.OptionalString(d, "transitions.0.no_audio")
		nextTransitions.RecordingComplete = utils.OptionalString(d, "transitions.0.recording_complete")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.RecordVoicemail{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.RecordVoicemailProperties{
			Trim:                       utils.OptionalString(d, "trim"),
			Transcribe:                 utils.OptionalBool(d, "transcribe"),
			TranscriptionCallbackURL:   utils.OptionalString(d, "transcription_callback_url"),
			PlayBeep:                   utils.OptionalString(d, "play_beep"),
			FinishOnKey:                utils.OptionalString(d, "finish_on_key"),
			RecordingStatusCallbackURL: utils.OptionalString(d, "recording_status_callback_url"),
			Timeout:                    utils.OptionalInt(d, "timeout"),
			MaxLength:                  utils.OptionalInt(d, "max_length"),
			Offset:                     offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Record voicemail widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create record voicemail widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal record voicemail widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
