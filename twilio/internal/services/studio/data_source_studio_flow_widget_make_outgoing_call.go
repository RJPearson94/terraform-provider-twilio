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

func dataSourceStudioFlowWidgetMakeOutgoingCall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetMakeOutgoingCallRead,

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
						"answered": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the outgoing call is answered",
						},
						"busy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the called party is busy",
						},
						"failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the outgoing call fails",
						},
						"no_answer": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the outgoing call is not answered",
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
			"detect_answering_machine": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable answering machine detection on the outgoing call",
			},
			"from": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{{flow.channel.address}}",
				Description: "The caller ID phone number for the outgoing call. Defaults to `{{flow.channel.address}}`",
			},
			"machine_detection": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"Enable",
						"DetectMessageEnd",
					}, false),
				),
				Description: "The answering machine detection mode. Valid values: `Enable`, `DetectMessageEnd`",
			},
			"machine_detection_silence_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.StringDigitsBetween(2000, 10000),
				),
				Description: "The duration of silence in milliseconds after a greeting that indicates a machine (2000–10000)",
			},
			"machine_detection_speech_end_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.StringDigitsBetween(500, 5000),
				),
				Description: "The duration of no speech in milliseconds that marks the end of a greeting (500–5000)",
			},
			"machine_detection_speech_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.StringDigitsBetween(1000, 6000),
				),
				Description: "The minimum duration of speech in milliseconds to classify as a human (1000–6000)",
			},
			"machine_detection_timeout": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.StringDigitsBetween(3, 120),
				),
				Description: "The total time in seconds to wait for machine detection to complete (3–120)",
			},
			"record": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to record the outgoing call",
			},
			"recording_channels": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"mono",
					"dual",
				}, false),
				Description: "The number of recording channels. Valid values: `mono`, `dual`",
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
			"send_digits": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DTMF digits to send after the call is connected (e.g. `ww1234#` where `w` is a 0.5s pause)",
			},
			"sip_auth_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password for SIP authentication. Sensitive — will not be shown in logs or plans",
			},
			"sip_auth_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username for SIP authentication",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds to wait for the call to be answered before timing out",
			},
			"to": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{{contact.channel.address}}",
				Description: "The phone number or SIP URI to call. Defaults to `{{contact.channel.address}}`",
			},
			"trim": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"trim-silence",
					"do-not-trim",
				}, false),
				Description: "Whether to trim silence from the beginning and end of recordings. Valid values: `trim-silence`, `do-not-trim`",
			},
		},
	}
}

func dataSourceStudioFlowWidgetMakeOutgoingCallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.MakeOutgoingCallNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Answered = utils.OptionalString(d, "transitions.0.answered")
		nextTransitions.Busy = utils.OptionalString(d, "transitions.0.busy")
		nextTransitions.Failed = utils.OptionalString(d, "transitions.0.failed")
		nextTransitions.NoAnswer = utils.OptionalString(d, "transitions.0.no_answer")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.MakeOutgoingCall{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.MakeOutgoingCallProperties{
			DetectAnsweringMachine:             utils.OptionalBool(d, "detect_answering_machine"),
			From:                               d.Get("from").(string),
			MachineDetection:                   utils.OptionalString(d, "machine_detection"),
			MachineDetectionSilenceTimeout:     utils.OptionalString(d, "machine_detection_silence_timeout"),
			MachineDetectionSpeechEndThreshold: utils.OptionalString(d, "machine_detection_speech_end_threshold"),
			MachineDetectionSpeechThreshold:    utils.OptionalString(d, "machine_detection_speech_threshold"),
			MachineDetectionTimeout:            utils.OptionalString(d, "machine_detection_timeout"),
			Offset:                             offset,
			Record:                             utils.OptionalBool(d, "record"),
			RecordingChannels:                  utils.OptionalString(d, "recording_channels"),
			RecordingStatusCallbackURL:         utils.OptionalString(d, "recording_status_callback_url"),
			SendDigits:                         utils.OptionalString(d, "send_digits"),
			SipAuthPassword:                    utils.OptionalString(d, "sip_auth_password"),
			SipAuthUsername:                    utils.OptionalString(d, "sip_auth_username"),
			Timeout:                            utils.OptionalInt(d, "timeout"),
			To:                                 d.Get("to").(string),
			Trim:                               utils.OptionalString(d, "trim"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Make outgoing call widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create make outgoing call widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal make outgoing call to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
