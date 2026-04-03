package taskrouter

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1/workspace/workers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTaskRouterWorkers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTaskRouterWorkersRead,

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
			"activity_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter workers by activity name",
			},
			"activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
				Description:  "Filter workers by activity SID",
			},
			"available": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Filter workers by availability",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter workers by friendly name",
			},
			"target_workers_expression": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter workers by a target workers expression",
			},
			"task_queue_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter workers by task queue name",
			},
			"task_queue_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterTaskQueueSidValidation(),
				Description:  "Filter workers by task queue SID",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the workers",
			},
			"workers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of workers in the workspace",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this worker by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the worker",
						},
						"activity_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the worker's current activity",
						},
						"attributes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A JSON string of attributes for the worker",
						},
						"activity_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The friendly name of the worker's current activity",
						},
						"available": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the worker is available to receive tasks",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the worker was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the worker was last updated, in RFC 3339 format",
						},
						"date_status_changed": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the worker's activity status last changed, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the worker resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceTaskRouterWorkersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	twilioClient := meta.(*common.TwilioClient)
	client := twilioClient.TaskRouter

	options := &workers.WorkersPageOptions{
		ActivityName:            utils.OptionalString(d, "activity_name"),
		ActivitySid:             utils.OptionalString(d, "activity_sid"),
		Available:               utils.OptionalBool(d, "available"),
		FriendlyName:            utils.OptionalString(d, "friendly_name"),
		TargetWorkersExpression: utils.OptionalString(d, "target_workers_expression"),
		TaskQueueName:           utils.OptionalString(d, "task_queue_name"),
		TaskQueueSid:            utils.OptionalString(d, "task_queue_sid"),
	}

	workspaceSid := d.Get("workspace_sid").(string)
	paginator := client.Workspace(workspaceSid).Workers.NewWorkersPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No workers were found for taskrouter workspace with sid (%s)", workspaceSid)
		}
		return diag.Errorf("Failed to read taskrouter worker: %s", err.Error())
	}

	d.SetId(workspaceSid)
	d.Set("workspace_sid", workspaceSid)
	d.Set("account_sid", twilioClient.AccountSid)

	workers := make([]interface{}, 0)

	for _, worker := range paginator.Workers {
		workersMap := make(map[string]interface{})

		workersMap["sid"] = worker.Sid
		workersMap["friendly_name"] = worker.FriendlyName
		workersMap["activity_sid"] = worker.ActivitySid
		workersMap["attributes"] = worker.Attributes
		workersMap["activity_name"] = worker.ActivityName
		workersMap["available"] = worker.Available
		workersMap["date_created"] = worker.DateCreated.Format(time.RFC3339)

		if worker.DateUpdated != nil {
			workersMap["date_updated"] = worker.DateUpdated.Format(time.RFC3339)
		}

		if worker.DateStatusChanged != nil {
			workersMap["date_status_changed"] = worker.DateStatusChanged.Format(time.RFC3339)
		}

		workersMap["url"] = worker.URL

		workers = append(workers, workersMap)
	}

	d.Set("workers", &workers)

	return nil
}
