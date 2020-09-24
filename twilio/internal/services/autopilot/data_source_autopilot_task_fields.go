package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotTaskFields() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotTaskFieldsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
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
				},
			},
		},
	}
}

func dataSourceAutopilotTaskFieldsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	paginator := client.Assistant(assistantSid).Task(taskSid).Fields.NewFieldsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No task fields were found for assistant with sid (%s) and task with sid (%s)", assistantSid, taskSid)
		}
		return fmt.Errorf("[ERROR] Failed to list autopilot task fields: %s", err.Error())
	}

	d.SetId(assistantSid + "/" + taskSid)
	d.Set("assistant_sid", assistantSid)
	d.Set("task_sid", taskSid)

	fields := make([]interface{}, 0)

	for _, field := range paginator.Fields {
		d.Set("account_sid", field.AccountSid)

		fieldMap := make(map[string]interface{})

		fieldMap["sid"] = field.Sid
		fieldMap["unique_name"] = field.UniqueName
		fieldMap["field_type"] = field.FieldType
		fieldMap["date_created"] = field.DateCreated.Format(time.RFC3339)

		if field.DateUpdated != nil {
			fieldMap["date_updated"] = field.DateUpdated.Format(time.RFC3339)
		}

		fieldMap["url"] = field.URL

		fields = append(fields, fieldMap)
	}

	d.Set("fields", &fields)

	return nil
}
