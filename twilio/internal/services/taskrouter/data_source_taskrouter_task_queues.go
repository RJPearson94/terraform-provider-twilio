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
				Description:  "The SID of the TaskRouter workspace",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the task queues",
			},
			"task_queues": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of task queues in the workspace",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this task queue by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the task queue",
						},
						"assignment_activity_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the activity to assign workers when a task is assigned",
						},
						"assignment_activity_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the activity to assign workers when a task is assigned",
						},
						"reservation_activity_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the activity to assign workers when a task is reserved",
						},
						"reservation_activity_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the activity to assign workers when a task is reserved",
						},
						"max_reserved_workers": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of workers to reserve for a task in this queue",
						},
						"target_workers": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A worker selection expression for this task queue",
						},
						"task_order": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The order in which tasks are assigned to workers",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the task queue was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the task queue was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the task queue resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceTaskRouterTaskQueuesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.TaskRouter

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
	d.Set("account_sid", twilioClient.AccountSid)

	taskQueues := make([]interface{}, 0)

	for _, taskQueue := range paginator.TaskQueues {
		taskQueuesMap := make(map[string]interface{})

		taskQueuesMap["sid"] = taskQueue.Sid
		taskQueuesMap["friendly_name"] = taskQueue.FriendlyName
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
