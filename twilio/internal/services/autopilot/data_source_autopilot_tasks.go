package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

func dataSourceAutopilotTasks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotTasksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tasks": {
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
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actions_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actions": {
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

func dataSourceAutopilotTasksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).Tasks.NewTasksPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No tasks were found for assistant with sid (%s)", assistantSid)
		}
		return diag.Errorf("Failed to list autopilot tasks: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	tasks := make([]interface{}, 0)

	for _, task := range paginator.Tasks {
		d.Set("account_sid", task.AccountSid)

		taskMap := make(map[string]interface{})

		taskMap["sid"] = task.Sid
		taskMap["unique_name"] = task.UniqueName
		taskMap["friendly_name"] = task.FriendlyName
		taskMap["actions_url"] = task.ActionsURL
		taskMap["date_created"] = task.DateCreated.Format(time.RFC3339)

		if task.DateUpdated != nil {
			taskMap["date_updated"] = task.DateUpdated.Format(time.RFC3339)
		}

		taskMap["url"] = task.URL

		getActionsResponse, err := client.Assistant(task.AssistantSid).Task(task.Sid).Actions().FetchWithContext(ctx)
		if err != nil {
			return diag.Errorf("Failed to read autopilot task actions: %s", err.Error())
		}
		actionsJSONString, err := structure.FlattenJsonToString(getActionsResponse.Data)
		if err != nil {
			return diag.Errorf("Unable to flatten actions json to string: %s", err.Error())
		}

		taskMap["actions"] = actionsJSONString

		tasks = append(tasks, taskMap)
	}

	d.Set("tasks", &tasks)

	return nil
}
