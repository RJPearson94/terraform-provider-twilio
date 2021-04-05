package conversations

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/configuration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConversationsConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsConfigurationCreate,
		ReadContext:   resourceConversationsConfigurationRead,
		UpdateContext: resourceConversationsConfigurationUpdate,
		DeleteContext: resourceConversationsConfigurationDelete,

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_service_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
			},
			"default_closed_timer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_inactive_timer": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_messaging_service_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConversationsConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Conversations configuration already exists so updating the configuration
	return resourceConversationsConfigurationUpdate(ctx, d, meta)
}

func resourceConversationsConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Configuration().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversations configuration: %s", err.Error())
	}

	d.Set("account_sid", getResponse.AccountSid)
	d.Set("default_service_sid", getResponse.DefaultChatServiceSid)
	d.Set("default_closed_timer", getResponse.DefaultClosedTimer)
	d.Set("default_inactive_timer", getResponse.DefaultInactiveTimer)
	d.Set("default_messaging_service_sid", getResponse.DefaultMessagingServiceSid)
	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &configuration.UpdateConfigurationInput{
		DefaultChatServiceSid:      utils.OptionalString(d, "default_service_sid"),
		DefaultClosedTimer:         utils.OptionalString(d, "default_closed_timer"),
		DefaultInactiveTimer:       utils.OptionalString(d, "default_inactive_timer"),
		DefaultMessagingServiceSid: utils.OptionalString(d, "default_messaging_service_sid"),
	}

	updateResp, err := client.Configuration().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations configuration: %s", err.Error())
	}

	d.SetId(updateResp.AccountSid)
	return resourceConversationsConfigurationRead(ctx, d, meta)
}

func resourceConversationsConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Conversations configuration cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
