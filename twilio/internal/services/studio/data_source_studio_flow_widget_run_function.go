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
						"fail": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the function execution fails",
						},
						"success": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the function executes successfully",
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
			"environment_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
				Description:  "The SID of the Serverless environment to execute the function in",
			},
			"function_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ServerlessFunctionSidValidation(),
				Description:  "The SID of the Serverless function to execute",
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
				Description: "The SID of the Serverless service containing the function, or `default` for the default service",
			},
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsURLWithHTTPS,
				Description:  "The HTTPS URL of the Serverless function to execute",
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Key/value parameters to pass to the function as arguments",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The parameter name",
						},
						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "The parameter value",
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
