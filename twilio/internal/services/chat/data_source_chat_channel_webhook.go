package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannelWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChatChannelWebhookRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
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
