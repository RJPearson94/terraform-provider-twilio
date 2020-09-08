package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotTaskSamples() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotTaskSamplesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"samples": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
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
				},
			},
		},
	}
}

func dataSourceAutopilotTaskSamplesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	taskSid := d.Get("task_sid").(string)
	paginator := client.Assistant(assistantSid).Task(taskSid).Samples.NewSamplesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No samples were found for assistant with sid (%s) and task with sid (%s)", assistantSid, taskSid)
		}
		return fmt.Errorf("[ERROR] Failed to list autopilot task samples: %s", err.Error())
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
