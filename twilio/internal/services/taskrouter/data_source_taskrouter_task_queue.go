package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterTaskQueue() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterTaskQueueRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterTaskQueueSidValidation(),
			},
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"event_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_activity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_activity_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reservation_activity_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_reserved_workers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"target_workers": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_order": {
				Type:     schema.TypeString,
				Computed: true,
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

func dataSourceTaskRouterTaskQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Workspace(workspaceSid).TaskQueue(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Task queue with sid (%s) was not found for taskrouter workspace with sid (%s)", sid, workspaceSid)
		}
		return diag.Errorf("Failed to read task queue: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
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
