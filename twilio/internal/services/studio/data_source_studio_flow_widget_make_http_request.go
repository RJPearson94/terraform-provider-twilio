package studio

import (
	"context"
	"fmt"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetMakeHttpRequest() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetMakeHttpRequestRead,

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
						"failed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the HTTP request fails",
						},
						"success": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the HTTP request succeeds",
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
			"body": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The request body content for POST requests",
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"application/x-www-form-urlencoded",
					"application/json",
				}, false),
				Description: "The Content-Type header for the request. Valid values: `application/x-www-form-urlencoded`, `application/json`",
			},
			"charset": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "utf-8",
				Description: "The character encoding for the request body. Defaults to `utf-8`",
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
				Description: "The HTTP method for the request. Valid values: `GET`, `POST`",
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Key/value parameters to include in the request as query parameters (GET) or form data (POST)",
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
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The HTTP/HTTPS URL to send the request to. Supports Liquid template expressions",
			},
		},
	}
}

func dataSourceStudioFlowWidgetMakeHttpRequestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.MakeHTTPRequestNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Failed = utils.OptionalString(d, "transitions.0.failed")
		nextTransitions.Success = utils.OptionalString(d, "transitions.0.success")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	var requestParameters *[]widgets.MakeHTTPRequestParameter
	if v, ok := d.GetOk("parameters"); ok {
		parameters := []widgets.MakeHTTPRequestParameter{}
		for _, parameter := range v.([]interface{}) {
			parameterMap := parameter.(map[string]interface{})
			parameters = append(parameters, widgets.MakeHTTPRequestParameter{
				Key:   parameterMap["key"].(string),
				Value: parameterMap["value"].(string),
			})
		}
		requestParameters = &parameters
	}

	widget := &widgets.MakeHTTPRequest{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.MakeHTTPRequestProperties{
			Body:        utils.OptionalString(d, "body"),
			ContentType: fmt.Sprintf("%s;charset=%s", d.Get("content_type").(string), d.Get("charset").(string)),
			Method:      d.Get("method").(string),
			Offset:      offset,
			Parameters:  requestParameters,
			URL:         d.Get("url").(string),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Make HTTP request widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create make HTTP request widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal make HTTP request to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
