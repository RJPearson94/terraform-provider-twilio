package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceChatChannelStudioWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatChannelStudioWebhookCreate,
		Read:   resourceChatChannelStudioWebhookRead,
		Update: resourceChatChannelStudioWebhookUpdate,
		Delete: resourceChatChannelStudioWebhookDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
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
	}
}

func resourceChatChannelStudioWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &webhooks.CreateChannelWebhookInput{
		Type:                    "studio",
		ConfigurationFlowSid:    utils.OptionalString(d, "flow_sid"),
		ConfigurationRetryCount: utils.OptionalInt(d, "retry_count"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhooks.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat channel webhook: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceChatChannelStudioWebhookRead(d, meta)
}

func resourceChatChannelStudioWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
				d.SetId("")
				return nil
			}
			if twilioError.IsNotFoundError() {
				d.SetId("")
				return nil
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel webhook: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("channel_sid", getResponse.ChannelSid)
	d.Set("type", getResponse.Type)
	d.Set("flow_sid", getResponse.Configuration.FlowSid)
	d.Set("retry_count", getResponse.Configuration.RetryCount)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceChatChannelStudioWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &webhook.UpdateChannelWebhookInput{
		ConfigurationFlowSid:    utils.OptionalString(d, "flow_sid"),
		ConfigurationRetryCount: utils.OptionalInt(d, "retry_count"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat channel webhook: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatChannelStudioWebhookRead(d, meta)
}

func resourceChatChannelStudioWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete chat channel webhook: %s", err.Error())
	}
	d.SetId("")
	return nil
}
