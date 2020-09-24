package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatChannelWebhooks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatChannelWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhooks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"webhook_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"filters": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"flow_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"triggers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"retry_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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

func dataSourceChatChannelWebhooksRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	paginator := client.Service(serviceSid).Channel(channelSid).Webhooks.NewChannelWebhooksPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] No channel webhooks were found for chat service with sid (%s) and channel with sid (%s)", serviceSid, channelSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel webhook: %s", err)
	}

	d.SetId(serviceSid + "/" + channelSid)
	d.Set("service_sid", serviceSid)
	d.Set("channel_sid", channelSid)

	webhooks := make([]interface{}, 0)

	for _, webhook := range paginator.Webhooks {
		d.Set("account_sid", webhook.AccountSid)

		webhookMap := make(map[string]interface{})

		webhookMap["sid"] = webhook.Sid
		webhookMap["type"] = webhook.Type
		webhookMap["configuration"] = flattenPageConfiguration(webhook.Configuration)
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

func flattenPageConfiguration(input webhooks.PageChannelWebhookResponseConfiguration) *[]interface{} {
	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["webhook_url"] = input.URL
	result["method"] = input.Method
	result["retry_count"] = input.RetryCount
	result["triggers"] = input.Triggers
	result["flow_sid"] = input.FlowSid
	result["filters"] = input.Filters
	results = append(results, result)

	return &results
}
