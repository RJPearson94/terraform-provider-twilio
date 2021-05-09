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

func dataSourceStudioFlowWidgetRunFunction() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetRunFunctionRead,

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
						"fail": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"success": {
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
			"environment_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
			},
			"function_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ServerlessFunctionSidValidation(),
			},
			"service_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					validation.StringInSlice([]string{
						"default",
					}, false),
					utils.ServerlessServiceSidValidation(),
				),
			},
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsURLWithHTTPS,
			},
			"parameters": {
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
		},
	}
}

func dataSourceStudioFlowWidgetRunFunctionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.RunFunctionNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Fail = utils.OptionalString(d, "transitions.0.fail")
		nextTransitions.Success = utils.OptionalString(d, "transitions.0.success")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	var functionParameters *[]widgets.RunFunctionParameter
	if v, ok := d.GetOk("parameters"); ok {
		parameters := []widgets.RunFunctionParameter{}
		for _, parameter := range v.([]interface{}) {
			parameterMap := parameter.(map[string]interface{})
			parameters = append(parameters, widgets.RunFunctionParameter{
				Key:   parameterMap["key"].(string),
				Value: parameterMap["value"].(string),
			})
		}
		functionParameters = &parameters
	}

	widget := &widgets.RunFunction{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.RunFunctionProperties{
			EnvironmentSid: utils.OptionalString(d, "environment_sid"),
			FunctionSid:    utils.OptionalString(d, "function_sid"),
			ServiceSid:     utils.OptionalString(d, "service_sid"),
			URL:            d.Get("url").(string),
			Parameters:     functionParameters,
			Offset:         offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Run function widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create run function widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal run function widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
