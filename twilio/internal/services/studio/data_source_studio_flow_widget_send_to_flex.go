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

func dataSourceStudioFlowWidgetSendToFlex() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetSendToFlexRead,

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
						"call_complete": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"call_failure": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"failed_to_enqueue": {
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
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"channel_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterTaskChannelSidValidation(),
			},
			"priority": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"wait_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"wait_url_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
			},
			"workflow_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkflowSidValidation(),
			},
		},
	}
}

func dataSourceStudioFlowWidgetSendToFlexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.SendToFlexNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.CallComplete = utils.OptionalString(d, "transitions.0.call_complete")
		nextTransitions.CallFailure = utils.OptionalString(d, "transitions.0.call_failure")
		nextTransitions.FailedToEnqueue = utils.OptionalString(d, "transitions.0.failed_to_enqueue")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := widgets.SendToFlex{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.SendToFlexProperties{
			Attributes:    utils.OptionalJSONString(d, "attributes"),
			Channel:       d.Get("channel_sid").(string),
			Offset:        offset,
			Priority:      utils.OptionalString(d, "priority"),
			Timeout:       utils.OptionalString(d, "timeout"),
			WaitURL:       utils.OptionalString(d, "wait_url"),
			WaitURLMethod: utils.OptionalString(d, "wait_url_method"),
			Workflow:      d.Get("workflow_sid").(string),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Send to flex widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create send to flex widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal send to flex widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
