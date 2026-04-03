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

func dataSourceStudioFlowWidgetGatherInputOnCall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetGatherInputOnCallRead,

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
						"keypress": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the caller presses a key (DTMF input)",
						},
						"speech": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when speech input is detected",
						},
						"timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the gather times out without input",
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
			"finish_on_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The DTMF key that signals the end of digit input (e.g. `#`)",
			},
			"gather_language": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language for automatic speech recognition (e.g. `en-US`)",
			},
			"hints": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of words or phrases to improve speech recognition accuracy",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"language": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"play"},
				Description:   "The language for text-to-speech when using `say`. Conflicts with `play`",
			},
			"loop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "How many times to repeat the prompt before timing out. Use 0 for infinite loop",
			},
			"number_of_digits": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The exact number of DTMF digits to collect before automatically submitting",
			},
			"play": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				ExactlyOneOf: []string{"say", "play"},
				Description:  "URL of the audio file to play as a prompt. Exactly one of `say` or `play` must be set",
			},
			"profanity_filter": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"true",
						"false",
					}, false),
				),
				Description: "Whether to filter profanity from speech recognition results. Valid values: `true`, `false`",
			},
			"say": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"say", "play"},
				Description:  "Text to speak to the caller as a prompt using text-to-speech. Exactly one of `say` or `play` must be set",
			},
			"speech_model": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"default",
						"numbers_and_commands",
						"phone_call",
					}, false),
				),
				Description: "The speech recognition model to use. Valid values: `default`, `numbers_and_commands`, `phone_call`",
			},
			"speech_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The duration of silence (in seconds) before speech recognition ends, or `auto` for automatic detection",
			},
			"stop_gather": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to allow a keypress to stop the current audio and submit the gathered input",
			},
			"timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 30),
				Description:  "The number of seconds to wait for a DTMF keypress before timing out (1–30)",
			},
			"voice": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"play"},
				Description:   "The text-to-speech voice to use when using `say` (e.g. `alice`). Conflicts with `play`",
			},
		},
	}
}

func dataSourceStudioFlowWidgetGatherInputOnCallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.GatherInputOnCallNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Keypress = utils.OptionalString(d, "transitions.0.keypress")
		nextTransitions.Speech = utils.OptionalString(d, "transitions.0.speech")
		nextTransitions.Timeout = utils.OptionalString(d, "transitions.0.timeout")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.GatherInputOnCall{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.GatherInputOnCallProperties{
			FinishOnKey:     utils.OptionalString(d, "finish_on_key"),
			GatherLanguage:  utils.OptionalString(d, "gather_language"),
			Hints:           utils.OptionalSeperatedString(d, "hints", ","),
			Language:        utils.OptionalString(d, "language"),
			Loop:            utils.OptionalInt(d, "loop"),
			NumberOfDigits:  utils.OptionalInt(d, "number_of_digits"),
			Offset:          offset,
			Play:            utils.OptionalString(d, "play"),
			ProfanityFilter: utils.OptionalString(d, "profanity_filter"),
			Say:             utils.OptionalString(d, "say"),
			SpeechModel:     utils.OptionalString(d, "speech_model"),
			SpeechTimeout:   utils.OptionalString(d, "speech_timeout"),
			StopGather:      utils.OptionalBool(d, "stop_gather"),
			Timeout:         utils.OptionalInt(d, "timeout"),
			Voice:           utils.OptionalString(d, "voice"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Gather input on call widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create gather input on call widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal gather input on call widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
