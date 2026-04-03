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

func dataSourceStudioFlowWidgetConnectVirtualAgent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetConnectVirtualAgentRead,

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
							Description: "The name of the next widget when the caller hangs up during the virtual agent session",
						},
						"return": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the virtual agent session completes and returns control",
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
			"connector": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The unique name of the virtual agent connector to use",
			},
			"language": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The BCP-47 language tag for the virtual agent session (e.g. `en-US`)",
			},
			"sentiment_analysis": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"true",
						"false",
					}, false),
				),
				Description: "Whether to enable sentiment analysis during the virtual agent session. Valid values: `true`, `false`",
			},
			"status_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The HTTP/HTTPS URL to receive status callback events for the virtual agent session",
			},
		},
	}
}

func dataSourceStudioFlowWidgetConnectVirtualAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.ConnectVirtualAgentNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Hangup = utils.OptionalString(d, "transitions.0.hangup")
		nextTransitions.Return = utils.OptionalString(d, "transitions.0.return")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.ConnectVirtualAgent{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.ConnectVirtualAgentProperties{
			Connector:         d.Get("connector").(string),
			Language:          utils.OptionalString(d, "language"),
			Offset:            offset,
			SentimentAnalysis: utils.OptionalString(d, "sentiment_analysis"),
			StatusCallbackURL: utils.OptionalString(d, "status_callback_url"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Connect virtual agent widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create connect virtual agent widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal connect virtual agent widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
