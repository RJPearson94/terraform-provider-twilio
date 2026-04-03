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

func dataSourceStudioFlowWidgetAddTwiMLRedirect() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetAddTwiMLRedirectRead,

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
							Description: "The name of the next widget when the TwiML redirect request fails",
						},
						"return": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the TwiML redirect completes and returns",
						},
						"timeout": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the TwiML redirect request times out",
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
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.IsURLWithHTTPorHTTPS,
				),
				Description: "The absolute URL of the TwiML document to redirect to. Must be an HTTP/HTTPS URL or a Liquid template expression",
			},
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					validation.StringInSlice([]string{
						"GET",
						"POST",
					}, false),
				),
				Description: "The HTTP method to use when fetching the TwiML document. Valid values: `GET`, `POST`",
			},
			"timeout": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.Any(
					utils.StudioFlowWidgetLiquidTemplateValidation(),
					utils.StringDigitsBetween(0, 14400),
				),
				Description: "The timeout in seconds for the TwiML redirect request (0–14400)",
			},
		},
	}
}

func dataSourceStudioFlowWidgetAddTwiMLRedirectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.AddTwimlRedirectNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Fail = utils.OptionalString(d, "transitions.0.fail")
		nextTransitions.Return = utils.OptionalString(d, "transitions.0.return")
		nextTransitions.Timeout = utils.OptionalString(d, "transitions.0.timeout")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.AddTwimlRedirect{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.AddTwimlRedirectProperties{
			URL:     d.Get("url").(string),
			Method:  utils.OptionalString(d, "method"),
			Timeout: utils.OptionalString(d, "timeout"),
			Offset:  offset,
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Add TwiML redirect widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create add TwiML redirect widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal add TwiML redirect to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
