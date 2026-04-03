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

func dataSourceStudioFlowWidgetSayPlay() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSayPlayRead,

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
						"audio_complete": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when audio playback or text-to-speech finishes",
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
			"digits": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"digits", "play", "say"},
				Description:  "DTMF digits to play to the caller. Exactly one of `digits`, `play`, or `say` must be set",
			},
			"language": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"digits", "play"},
				Description:   "The language for text-to-speech (e.g. `en-US`). Conflicts with `digits` and `play`",
			},
			"loop": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "How many times to repeat the audio or speech. Use 0 for infinite loop",
			},
			"play": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				ExactlyOneOf: []string{"digits", "play", "say"},
				Description:  "URL of the audio file to play. Exactly one of `digits`, `play`, or `say` must be set",
			},
			"say": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"digits", "play", "say"},
				Description:  "Text to speak to the caller using text-to-speech. Exactly one of `digits`, `play`, or `say` must be set",
			},
			"voice": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"digits", "play"},
				Description:   "The text-to-speech voice to use (e.g. `alice`). Conflicts with `digits` and `play`",
			},
		},
	}
}

func dataSourceStudioFlowWidgetSayPlayRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SayPlayNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.AudioComplete = utils.OptionalString(d, "transitions.0.audio_complete")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.SayPlay{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SayPlayProperties{
			Digits:   utils.OptionalString(d, "digits"),
			Language: utils.OptionalString(d, "language"),
			Loop:     utils.OptionalInt(d, "loop"),
			Offset:   offset,
			Play:     utils.OptionalString(d, "play"),
			Say:      utils.OptionalString(d, "say"),
			Voice:    utils.OptionalString(d, "voice"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Say play widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create say play widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal say play widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
