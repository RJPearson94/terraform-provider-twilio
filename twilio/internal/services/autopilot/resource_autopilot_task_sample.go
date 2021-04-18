package autopilot

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/sample"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/task/samples"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAutopilotTaskSample() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutopilotTaskSampleCreate,
		ReadContext:   resourceAutopilotTaskSampleRead,
		UpdateContext: resourceAutopilotTaskSampleUpdate,
		DeleteContext: resourceAutopilotTaskSampleDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Assistants/(.*)/Tasks/(.*)/Samples/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("assistant_sid", match[1])
				d.Set("task_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"task_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AutopilotTaskSidValidation(),
			},
			"language": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"tagged_text": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"source_channel": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "voice",
				ValidateFunc: validation.StringInSlice([]string{
					"voice",
					"sms",
					"chat",
					"alexa",
					"google-assistant",
					"slack",
				}, false),
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

func resourceAutopilotTaskSampleCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	createInput := &samples.CreateSampleInput{
		Language:      d.Get("language").(string),
		TaggedText:    d.Get("tagged_text").(string),
		SourceChannel: utils.OptionalString(d, "source_channel"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Samples.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create autopilot task sample: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceAutopilotTaskSampleRead(ctx, d, meta)
}

func resourceAutopilotTaskSampleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read autopilot task sample: %s", err.Error())
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

func resourceAutopilotTaskSampleUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	updateInput := &sample.UpdateSampleInput{
		Language:      utils.OptionalString(d, "language"),
		TaggedText:    utils.OptionalString(d, "tagged_text"),
		SourceChannel: utils.OptionalString(d, "source_channel"),
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update autopilot task sample: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceAutopilotTaskSampleRead(ctx, d, meta)
}

func resourceAutopilotTaskSampleDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Get("assistant_sid").(string)).Task(d.Get("task_sid").(string)).Sample(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete autopilot task sample: %s", err.Error())
	}
	d.SetId("")
	return nil
}
