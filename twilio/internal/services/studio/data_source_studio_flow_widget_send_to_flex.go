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
						"call_complete": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the Flex interaction completes",
						},
						"call_failure": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the Flex interaction fails",
						},
						"failed_to_enqueue": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the task cannot be enqueued to Flex",
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
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
				Description:      "A JSON string of custom attributes to attach to the TaskRouter task",
			},
			"channel_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterTaskChannelSidValidation(),
				Description:  "The SID of the TaskRouter task channel (e.g. voice, chat) to use for routing in Flex",
			},
			"priority": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The priority of the TaskRouter task, as a string integer",
			},
			"timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The number of seconds before the TaskRouter task times out, as a string integer",
			},
			"wait_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The HTTP/HTTPS URL of the TwiML document to execute while the caller waits for a Flex agent",
			},
			"wait_url_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
				Description: "The HTTP method to use when fetching the `wait_url` document. Valid values: `GET`, `POST`",
			},
			"workflow_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkflowSidValidation(),
				Description:  "The SID of the TaskRouter workflow to route the task through in Flex",
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
