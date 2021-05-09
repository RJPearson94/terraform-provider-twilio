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
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keypress": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"speech": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"timeout": {
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
			"finish_on_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gather_language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hints": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"language": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"play"},
			},
			"loop": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"number_of_digits": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"play": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				ExactlyOneOf: []string{"say", "play"},
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
			},
			"say": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"say", "play"},
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
			},
			"speech_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stop_gather": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(1, 30),
			},
			"voice": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"play"},
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
