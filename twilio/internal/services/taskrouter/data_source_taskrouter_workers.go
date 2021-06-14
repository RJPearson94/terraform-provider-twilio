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
			},
			"activity_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"activity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterActivitySidValidation(),
			},
			"available": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"target_workers_expression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_queue_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"task_queue_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.TaskRouterTaskQueueSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"workers": {
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
						"activity_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"activity_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"available": {
							Type:     schema.TypeBool,
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
						"date_status_changed": {
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
