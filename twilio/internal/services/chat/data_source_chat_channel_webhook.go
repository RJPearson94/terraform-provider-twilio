package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannelWebhook() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatChannelWebhookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatChannelWebhookSidValidation(),
				Description:  "The SID of the channel webhook",
			},
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
				Description: "The SID of the account that owns this channel webhook",
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
	}
}

func dataSourceChatChannelWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Channel(channelSid).Webhook(sid).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("Channel webhook with sid (%s) was not found for chat service with sid (%s) and channel with sid (%s)", sid, serviceSid, channelSid)
			}
		}
		return diag.Errorf("Failed to read chat channel webhook: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("channel_sid", getResponse.ChannelSid)
	d.Set("type", getResponse.Type)
	d.Set("configuration", flattenFetchConfiguration(getResponse.Configuration))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func flattenFetchConfiguration(input webhook.FetchChannelWebhookConfigurationResponse) *[]interface{} {
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
