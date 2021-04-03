package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterTaskQueues() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterTaskQueuesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"workspace_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.TaskRouterWorkspaceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_queues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
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
				},
			},
		},
	}
}

func dataSourceTaskRouterTaskQueuesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).TaskRouter

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).TaskQueues.NewTaskQueuesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No task queues were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return diag.Errorf("Failed to read task queue: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)

	taskQueues := make([]interface{}, 0)

	for _, taskQueue := range paginator.TaskQueues {
		d.Set("account_sid", taskQueue.AccountSid)

		taskQueuesMap := make(map[string]interface{})

		taskQueuesMap["sid"] = taskQueue.Sid
		taskQueuesMap["friendly_name"] = taskQueue.FriendlyName
		taskQueuesMap["event_callback_url"] = taskQueue.EventCallbackURL
		taskQueuesMap["task_order"] = taskQueue.TaskOrder
		taskQueuesMap["assignment_activity_name"] = taskQueue.AssignmentActivityName
		taskQueuesMap["assignment_activity_sid"] = taskQueue.AssignmentActivitySid
		taskQueuesMap["reservation_activity_name"] = taskQueue.ReservationActivityName
		taskQueuesMap["reservation_activity_sid"] = taskQueue.ReservationActivitySid
		taskQueuesMap["target_workers"] = taskQueue.TargetWorkers
		taskQueuesMap["max_reserved_workers"] = taskQueue.MaxReservedWorkers
		taskQueuesMap["date_created"] = taskQueue.DateCreated.Format(time.RFC3339)

		if taskQueue.DateUpdated != nil {
			taskQueuesMap["date_updated"] = taskQueue.DateUpdated.Format(time.RFC3339)
		}

		taskQueuesMap["url"] = taskQueue.URL

		taskQueues = append(taskQueues, taskQueuesMap)
	}

	d.Set("task_queues", &taskQueues)

	return nil
}
