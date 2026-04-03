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

func dataSourceStudioFlowWidgetEnqueueCall() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetEnqueueCallRead,

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
							Description: "The name of the next widget when the enqueued call completes",
						},
						"call_failure": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the enqueued call fails",
						},
						"failed_to_enqueue": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the call cannot be added to the queue or workflow",
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
			"priority": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"queue_name"},
				Description:   "The priority of the TaskRouter task created for this call. Conflicts with `queue_name`",
			},
			"queue_name": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"workflow_sid", "queue_name"},
				Description:  "The name of the Twilio queue to enqueue the call into. At least one of `workflow_sid` or `queue_name` must be set",
			},
			"task_attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
				ConflictsWith:    []string{"queue_name"},
				Description:      "A JSON string of attributes to attach to the TaskRouter task. Conflicts with `queue_name`",
			},
			"timeout": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"queue_name"},
				Description:   "The number of seconds to wait for a worker to accept the task before timing out. Conflicts with `queue_name`",
			},
			"wait_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Description:  "The HTTP/HTTPS URL of the TwiML document to execute while the caller waits in the queue",
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
				Optional:     true,
				ValidateFunc: utils.TaskRouterWorkflowSidValidation(),
				AtLeastOneOf: []string{"workflow_sid", "queue_name"},
				Description:  "The SID of the TaskRouter workflow to route the call through. At least one of `workflow_sid` or `queue_name` must be set",
			},
		},
	}
}

func dataSourceStudioFlowWidgetEnqueueCallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.EnqueueCallNextTransitions{}
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

	widget := &widgets.EnqueueCall{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.EnqueueCallProperties{
			Offset:         offset,
			Priority:       utils.OptionalInt(d, "priority"),
			QueueName:      utils.OptionalString(d, "queue_name"),
			TaskAttributes: utils.OptionalString(d, "task_attributes"),
			Timeout:        utils.OptionalInt(d, "timeout"),
			WaitURL:        utils.OptionalString(d, "wait_url"),
			WaitURLMethod:  utils.OptionalString(d, "wait_url_method"),
			WorkflowSid:    utils.OptionalString(d, "workflow_sid"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Enqueue call widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create enqueue call widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal enqueue call widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
