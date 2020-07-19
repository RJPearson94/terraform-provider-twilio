package autopilot

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/sample"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/samples"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAutopilotTaskSample() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotTaskSampleCreate,
		Read:   resourceAutopilotTaskSampleRead,
		Update: resourceAutopilotTaskSampleUpdate,
		Delete: resourceAutopilotTaskSampleDelete,
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
			"assistant_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"task_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tagged_text": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_channel": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceAutopilotTaskSampleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &samples.CreateSampleInput{
		Language:      d.Get("language").(string),
		TaggedText:    d.Get("tagged_text").(string),
		SourceChannel: utils.OptionalString(d, "source_channel"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Samples.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot task sample: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotTaskSampleRead(d, meta)
}

func resourceAutopilotTaskSampleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot task sample: %s", err.Error())
	}

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

func resourceAutopilotTaskSampleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	updateInput := &sample.UpdateSampleInput{
		Language:      utils.OptionalString(d, "language"),
		TaggedText:    utils.OptionalString(d, "tagged_text"),
		SourceChannel: utils.OptionalString(d, "source_channel"),
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot task sample: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotTaskSampleRead(d, meta)
}

func resourceAutopilotTaskSampleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete autopilot task sample: %s", err.Error())
	}
	d.SetId("")
	return nil
}