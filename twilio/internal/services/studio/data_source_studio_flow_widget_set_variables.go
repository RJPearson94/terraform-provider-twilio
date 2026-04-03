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

func dataSourceStudioFlowWidgetSetVariables() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSetVariablesRead,

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
						"next": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget after the variables are set",
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
			"variables": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of flow variables to set or update",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The variable name",
						},
						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The variable value. Supports Liquid template expressions",
						},
					},
				},
			},
		},
	}
}

func dataSourceStudioFlowWidgetSetVariablesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SetVariablesNextTransitions{}
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

	var flowVariables *[]widgets.SetVariablesVariable
	if v, ok := d.GetOk("variables"); ok {
		variables := []widgets.SetVariablesVariable{}
		for _, variable := range v.([]interface{}) {
			variableMap := variable.(map[string]interface{})
			variables = append(variables, widgets.SetVariablesVariable{
				Key:   variableMap["key"].(string),
				Value: variableMap["value"].(string),
			})
		}
		flowVariables = &variables
	}

	widget := widgets.SetVariables{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SetVariablesProperties{
			Variables: flowVariables,
			Offset:    offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Set variables widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create set variables widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal set variables widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
