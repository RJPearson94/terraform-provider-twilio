package taskrouter

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTaskRouterTaskQueue() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTaskRouterTaskQueueCreate,
		ReadContext:   resourceTaskRouterTaskQueueRead,
		UpdateContext: resourceTaskRouterTaskQueueUpdate,
		DeleteContext: resourceTaskRouterTaskQueueDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Workspaces/(.*)/TaskQueues/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("workspace_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"assignment_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
			},
			"reservation_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
			},
			"max_reserved_workers": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 50),
			},
			"target_workers": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1==1",
			},
			"task_order": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FIFO",
				ValidateFunc: validation.StringInSlice([]string{
					"LIFO",
					"FIFO",
				}, false),
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceTaskRouterTaskQueueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &task_queues.CreateTaskQueueInput{
		FriendlyName:           d.Get("friendly_name").(string),
		AssignmentActivitySid:  utils.OptionalStringWithEmptyStringDefault(d, "assignment_activity_sid"),
		MaxReservedWorkers:     utils.OptionalInt(d, "max_reserved_workers"),
		TargetWorkers:          utils.OptionalString(d, "target_workers"),
		TaskOrder:              utils.OptionalString(d, "task_order"),
		ReservationActivitySid: utils.OptionalStringWithEmptyStringDefault(d, "reservation_activity_sid"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueues.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create task queue: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterTaskQueueRead(ctx, d, meta)
}

func resourceTaskRouterTaskQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read task queue: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("task_order", getResponse.TaskOrder)
	d.Set("assignment_activity_name", getResponse.AssignmentActivityName)
	d.Set("assignment_activity_sid", getResponse.AssignmentActivitySid)
	d.Set("reservation_activity_name", getResponse.ReservationActivityName)
	d.Set("reservation_activity_sid", getResponse.ReservationActivitySid)
	d.Set("target_workers", getResponse.TargetWorkers)
	d.Set("max_reserved_workers", getResponse.MaxReservedWorkers)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceTaskRouterTaskQueueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &task_queue.UpdateTaskQueueInput{
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		AssignmentActivitySid:  utils.OptionalStringWithEmptyStringDefault(d, "assignment_activity_sid"),
		MaxReservedWorkers:     utils.OptionalInt(d, "max_reserved_workers"),
		TargetWorkers:          utils.OptionalString(d, "target_workers"),
		TaskOrder:              utils.OptionalString(d, "task_order"),
		ReservationActivitySid: utils.OptionalStringWithEmptyStringDefault(d, "reservation_activity_sid"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update task queue: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterTaskQueueRead(ctx, d, meta)
}

func resourceTaskRouterTaskQueueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete task queue: %s", err.Error())
	}
	d.SetId("")
	return nil
}
