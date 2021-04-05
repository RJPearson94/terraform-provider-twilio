package conversations

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/webhook"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceConversationsWebhook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsWebhookCreate,
		ReadContext:   resourceConversationsWebhookRead,
		UpdateContext: resourceConversationsWebhookUpdate,
		DeleteContext: resourceConversationsWebhookDelete,

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"webhook",
					"flex",
				}, false),
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
			"pre_webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"post_webhook_url": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"filters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"onConversationAdd",
						"onConversationAdded",
						"onConversationRemove",
						"onConversationRemoved",
						"onConversationStateUpdated",
						"onConversationUpdate",
						"onConversationUpdated",
						"onDeliveryUpdated",
						"onMessageAdd",
						"onMessageAdded",
						"onMessageRemove",
						"onMessageRemoved",
						"onMessageUpdate",
						"onMessageUpdated",
						"onParticipantAdd",
						"onParticipantAdded",
						"onParticipantRemove",
						"onParticipantRemoved",
						"onParticipantUpdate",
						"onParticipantUpdated",
						"onUserAdded",
						"onUserUpdate",
						"onUserUpdated",
					}, false),
				},
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConversationsWebhookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Conversations webhook already exists so updating the configuration
	return resourceConversationsWebhookUpdate(ctx, d, meta)
}

func resourceConversationsWebhookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Webhook().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversation webhook: %s", err.Error())
	}

	d.Set("account_sid", getResponse.AccountSid)
	d.Set("target", getResponse.Target)
	d.Set("pre_webhook_url", getResponse.PreWebhookURL)
	d.Set("post_webhook_url", getResponse.PostWebhookURL)
	d.Set("filters", getResponse.Filters)
	d.Set("method", getResponse.Method)
	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsWebhookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &webhook.UpdateWebhookInput{
		PreWebhookURL:  utils.OptionalString(d, "pre_webhook_url"),
		PostWebhookURL: utils.OptionalString(d, "post_webhook_url"),
		Method:         utils.OptionalString(d, "method"),
		Filters:        utils.OptionalStringSlice(d, "filters"),
		Target:         utils.OptionalString(d, "target"),
	}

	updateResp, err := client.Webhook().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations webhook: %s", err.Error())
	}

	d.SetId(updateResp.AccountSid)
	return resourceConversationsWebhookRead(ctx, d, meta)
}

func resourceConversationsWebhookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Conversations webhook cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
