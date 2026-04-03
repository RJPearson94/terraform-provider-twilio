package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannelWebhooks() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatChannelWebhooksRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the Programmable Chat service",
			},
			"channel_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatChannelSidValidation(),
				Description:  "The SID of the chat channel",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the channel webhooks",
			},
			"webhooks": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of channel webhooks",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to the channel webhook by Twilio",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the channel webhook",
						},
						"configuration": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration of the channel webhook",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used for webhook requests",
									},
									"webhook_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to send webhook requests to",
									},
									"filters": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of events that trigger the webhook",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"flow_sid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SID of the Studio Flow to trigger",
									},
									"triggers": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of keywords that trigger the webhook",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"retry_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of retry attempts for failed webhook requests",
									},
								},
							},
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel webhook was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the channel webhook was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the channel webhook resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceChatChannelWebhooksRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

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
				return diag.Errorf("No channel webhooks were found for chat service with sid (%s) and channel with sid (%s)", serviceSid, channelSid)
			}
		}
		return diag.Errorf("Failed to read chat channel webhook: %s", err.Error())
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

func flattenPageConfiguration(input webhooks.PageChannelWebhookConfigurationResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"webhook_url": input.URL,
			"method":      input.Method,
			"retry_count": input.RetryCount,
			"triggers":    input.Triggers,
			"flow_sid":    input.FlowSid,
			"filters":     input.Filters,
		},
	}
}
