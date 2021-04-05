package conversations

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConversationsServiceConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConversationsServiceConfigurationCreate,
		ReadContext:   resourceConversationsServiceConfigurationRead,
		UpdateContext: resourceConversationsServiceConfigurationUpdate,
		DeleteContext: resourceConversationsServiceConfigurationDelete,

		Timeouts: &schema.ResourceTimeout{
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
			},
			"default_chat_service_role_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ConversationRoleSidValidation(),
			},
			"default_conversation_creator_role_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ConversationRoleSidValidation(),
			},
			"default_conversation_role_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ConversationRoleSidValidation(),
			},
			"reachability_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceConversationsServiceConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Conversations service configuration already exists so updating the configuration
	return resourceConversationsServiceConfigurationUpdate(ctx, d, meta)
}

func resourceConversationsServiceConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Service(d.Get("service_sid").(string)).Configuration().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read conversations service configuration: %s", err.Error())
	}

	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("default_chat_service_role_sid", getResponse.DefaultChatServiceRoleSid)
	d.Set("default_conversation_creator_role_sid", getResponse.DefaultConversationCreatorRoleSid)
	d.Set("default_conversation_role_sid", getResponse.DefaultConversationRoleSid)
	d.Set("reachability_enabled", getResponse.ReachabilityEnabled)
	d.Set("url", getResponse.URL)

	return nil
}

func resourceConversationsServiceConfigurationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	updateInput := &configuration.UpdateConfigurationInput{
		DefaultChatServiceRoleSid:         utils.OptionalString(d, "default_chat_service_role_sid"),
		DefaultConversationCreatorRoleSid: utils.OptionalString(d, "default_conversation_creator_role_sid"),
		DefaultConversationRoleSid:        utils.OptionalString(d, "default_conversation_role_sid"),
		ReachabilityEnabled:               utils.OptionalBool(d, "reachability_enabled"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Configuration().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update conversations service configuration: %s", err.Error())
	}

	d.SetId(updateResp.ChatServiceSid)
	return resourceConversationsServiceConfigurationRead(ctx, d, meta)
}

func resourceConversationsServiceConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Conversations service configuration cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
