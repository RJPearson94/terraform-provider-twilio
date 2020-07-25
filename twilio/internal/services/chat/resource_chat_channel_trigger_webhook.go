package chat

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhook"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/webhooks"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceChatChannelTriggerWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatChannelTriggerWebhookCreate,
		Read:   resourceChatChannelTriggerWebhookRead,
		Update: resourceChatChannelTriggerWebhookUpdate,
		Delete: resourceChatChannelTriggerWebhookDelete,
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
			"method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
				Computed: true,
			},
			"webhook_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"triggers": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceChatChannelTriggerWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	createInput := &webhooks.CreateChannelWebhookInput{
		Type:                    "trigger",
		ConfigurationURL:        utils.OptionalString(d, "webhook_url"),
		ConfigurationMethod:     utils.OptionalString(d, "method"),
		ConfigurationRetryCount: utils.OptionalInt(d, "retry_count"),
		ConfigurationTriggers:   utils.OptionalStringSlice(d, "triggers"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhooks.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat channel webhook: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceChatChannelTriggerWebhookRead(d, meta)
}

func resourceChatChannelTriggerWebhookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	getResponse, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).Get()
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
	d.Set("webhook_url", getResponse.Configuration.URL)
	d.Set("method", getResponse.Configuration.Method)
	d.Set("retry_count", getResponse.Configuration.RetryCount)
	d.Set("triggers", getResponse.Configuration.Triggers)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceChatChannelTriggerWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	updateInput := &webhook.UpdateChannelWebhookInput{
		ConfigurationURL:        utils.OptionalString(d, "webhook_url"),
		ConfigurationMethod:     utils.OptionalString(d, "method"),
		ConfigurationRetryCount: utils.OptionalInt(d, "retry_count"),
		ConfigurationTriggers:   utils.OptionalStringSlice(d, "triggers"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat channel webhook: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatChannelTriggerWebhookRead(d, meta)
}

func resourceChatChannelTriggerWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	if err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Webhook(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete chat channel webhook: %s", err.Error())
	}
	d.SetId("")
	return nil
}
