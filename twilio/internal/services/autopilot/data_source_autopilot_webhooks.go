package autopilot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotWebhooks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhooks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"events": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"webhook_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"webhook_method": &schema.Schema{
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

func dataSourceAutopilotWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).Webhooks.NewWebhooksPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No webhooks were found for assistant with sid (%s)", assistantSid)
		}
		return fmt.Errorf("[ERROR] Failed to list autopilot webhooks: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	webhooks := make([]interface{}, 0)

	for _, webhook := range paginator.Webhooks {
		d.Set("account_sid", webhook.AccountSid)

		webhookMap := make(map[string]interface{})

		webhookMap["sid"] = webhook.Sid
		webhookMap["unique_name"] = webhook.UniqueName
		webhookMap["webhook_url"] = webhook.WebhookURL
		webhookMap["webhook_method"] = webhook.WebhookMethod
		webhookMap["events"] = strings.Split(webhook.Events, " ")
		webhookMap["date_created"] = webhook.DateCreated.Format(time.RFC3339)

		if webhook.DateUpdated != nil {
			webhookMap["date_updated"] = webhook.DateUpdated.Format(time.RFC3339)
		}

		webhookMap["url"] = webhook.URL

		webhooks = append(webhooks, webhookMap)
	}

	d.Set("webhooks", &webhooks)

	return nil
}
