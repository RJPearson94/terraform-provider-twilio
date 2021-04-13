package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotTaskSamples() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotTaskSamplesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"samples": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"language": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"tagged_text": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_channel": {
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

func dataSourceAutopilotTaskSamplesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	paginator := client.Assistant(assistantSid).Task(taskSid).Samples.NewSamplesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No samples were found for assistant with sid (%s) and task with sid (%s)", assistantSid, taskSid)
		}
		return diag.Errorf("Failed to list autopilot task samples: %s", err.Error())
	}

	d.SetId(assistantSid + "/" + taskSid)
	d.Set("assistant_sid", assistantSid)
	d.Set("task_sid", taskSid)

	samples := make([]interface{}, 0)

	for _, sample := range paginator.Samples {
		d.Set("account_sid", sample.AccountSid)

		sampleMap := make(map[string]interface{})

		sampleMap["sid"] = sample.Sid
		sampleMap["language"] = sample.Language
		sampleMap["tagged_text"] = sample.TaggedText
		sampleMap["source_channel"] = sample.SourceChannel
		sampleMap["date_created"] = sample.DateCreated.Format(time.RFC3339)

		if sample.DateUpdated != nil {
			sampleMap["date_updated"] = sample.DateUpdated.Format(time.RFC3339)
		}

		sampleMap["url"] = sample.URL

		samples = append(samples, sampleMap)
	}

	d.Set("samples", &samples)

	return nil
}
