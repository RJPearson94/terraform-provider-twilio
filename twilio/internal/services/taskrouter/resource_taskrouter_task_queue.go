package taskrouter

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queue"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/task_queues"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceTaskRouterTaskQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskRouterTaskQueueCreate,
		Read:   resourceTaskRouterTaskQueueRead,
		Update: resourceTaskRouterTaskQueueUpdate,
		Delete: resourceTaskRouterTaskQueueDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"event_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assignment_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_activity_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"reservation_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_activity_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"max_reserved_workers": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"target_workers": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_order": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FIFO",
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

func resourceTaskRouterTaskQueueCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	createInput := &task_queues.CreateTaskQueueInput{
		FriendlyName:           d.Get("friendly_name").(string),
		AssignmentActivitySid:  utils.OptionalString(d, "assignment_activity_sid"),
		MaxReservedWorkers:     utils.OptionalInt(d, "max_reserved_workers"),
		TargetWorkers:          utils.OptionalString(d, "target_workers"),
		TaskOrder:              utils.OptionalString(d, "task_order"),
		ReservationActivitySid: utils.OptionalString(d, "reservation_activity_sid"),
	}

	createResult, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueues.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create task queue: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceTaskRouterTaskQueueRead(d, meta)
}

func resourceTaskRouterTaskQueueRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	getResponse, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read task queue: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("workspace_sid", getResponse.WorkspaceSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("event_callback_url", getResponse.EventCallbackURL)
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

func resourceTaskRouterTaskQueueUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	updateInput := &task_queue.UpdateTaskQueueInput{
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		AssignmentActivitySid:  utils.OptionalString(d, "assignment_activity_sid"),
		MaxReservedWorkers:     utils.OptionalInt(d, "max_reserved_workers"),
		TargetWorkers:          utils.OptionalString(d, "target_workers"),
		TaskOrder:              utils.OptionalString(d, "task_order"),
		ReservationActivitySid: utils.OptionalString(d, "reservation_activity_sid"),
	}

	updateResp, err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update task queue: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceTaskRouterTaskQueueRead(d, meta)
}

func resourceTaskRouterTaskQueueDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).TaskRouter

	if err := client.Workspace(d.Get("workspace_sid").(string)).TaskQueue(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete task queue: %s", err.Error())
	}
	d.SetId("")
	return nil
}
