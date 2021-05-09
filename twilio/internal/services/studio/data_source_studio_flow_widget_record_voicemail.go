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
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hangup": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"no_audio": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"recording_complete": {
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
			},
			"transcribe": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"transcription_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
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
			},
			"recording_status_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
			},
			"finish_on_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_length": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 14400),
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
