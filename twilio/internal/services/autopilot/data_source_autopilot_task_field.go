package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotTaskField() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotTaskFieldRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotTaskFieldSidValidation(),
			},
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"task_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotTaskSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_type": {
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

func dataSourceAutopilotTaskFieldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Assistant(assistantSid).Task(taskSid).Field(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Task field with sid (%s) was not found for assistant with sid (%s) and task with sid (%s)", sid, assistantSid, taskSid)
		}
		return diag.Errorf("Failed to read autopilot task field: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("task_sid", getResponse.TaskSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("field_type", getResponse.FieldType)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	return nil
}
