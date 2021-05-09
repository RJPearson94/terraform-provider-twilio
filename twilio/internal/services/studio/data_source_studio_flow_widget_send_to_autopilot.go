package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetSendToAutopilot() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSendToAutopilotRead,

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
						"failure": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"session_ended": {
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
			"autopilot_assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"body": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{{trigger.Message.Body}}",
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatChannelSidValidation(),
				),
			},
			"service_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.ChatServiceSidValidation(),
				),
			},
			"from": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{{flow.channel.address}}",
			},
			"memory_parameters": {
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
			"target_task": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  14400,
			},
		},
	}
}

func dataSourceStudioFlowWidgetSendToAutopilotRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SendToAutopilotNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Failure = utils.OptionalString(d, "transitions.0.failure")
		nextTransitions.SessionEnded = utils.OptionalString(d, "transitions.0.session_ended")
		nextTransitions.Timeout = utils.OptionalString(d, "transitions.0.timeout")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	var memoryParameters *[]widgets.SendToAutopilotMemoryParameter
	if v, ok := d.GetOk("memory_parameters"); ok {
		parameters := []widgets.SendToAutopilotMemoryParameter{}
		for _, parameter := range v.([]interface{}) {
			parameterMap := parameter.(map[string]interface{})
			parameters = append(parameters, widgets.SendToAutopilotMemoryParameter{
				Key:   parameterMap["key"].(string),
				Value: parameterMap["value"].(string),
			})
		}
		memoryParameters = &parameters
	}

	widget := widgets.SendToAutopilot{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SendToAutopilotProperties{
			AutopilotAssistantSid: d.Get("autopilot_assistant_sid").(string),
			Body:                  d.Get("body").(string),
			ChatAttributes:        utils.OptionalJSONString(d, "attributes"),
			ChatChannel:           utils.OptionalString(d, "channel_sid"),
			ChatService:           utils.OptionalString(d, "service_sid"),
			From:                  d.Get("from").(string),
			MemoryParameters:      memoryParameters,
			Offset:                offset,
			TargetTask:            utils.OptionalString(d, "target_task"),
			Timeout:               d.Get("timeout").(int),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Send to autopilot failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create send to autopilot widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal send to autopilot widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
