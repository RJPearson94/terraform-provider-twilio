package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotTaskSample() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotTaskSampleRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"assistant_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"task_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tagged_text": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_channel": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceAutopilotTaskSampleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Assistant(assistantSid).Task(taskSid).Sample(sid).FetchWithContext(ctx)

	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Task sample with sid (%s) was not found for assistant with sid (%s) and task with sid (%s)", sid, assistantSid, taskSid)
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot task sample: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("task_sid", getResponse.TaskSid)
	d.Set("language", getResponse.Language)
	d.Set("tagged_text", getResponse.TaggedText)
	d.Set("source_channel", getResponse.SourceChannel)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	return nil
}
