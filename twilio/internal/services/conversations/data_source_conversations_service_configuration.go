package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsServiceConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsServiceConfigurationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationServiceSidValidation(),
				Description:  "The SID of the conversations service",
			},
			"default_chat_service_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to users when they join the service",
			},
			"default_conversation_creator_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to the creator of a conversation",
			},
			"default_conversation_role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default role assigned to users when they are added to a conversation",
			},
			"reachability_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the reachability indicator is enabled for the service",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the service configuration resource",
			},
		},
	}
}

func dataSourceConversationsServiceConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	serviceSid := d.Get("service_sid").(string)
	getResponse, err := client.Service(serviceSid).Configuration().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation configuration was not found for service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read conversations service configuration: %s", err.Error())
	}

	d.SetId(getResponse.ChatServiceSid)
	d.Set("service_sid", getResponse.ChatServiceSid)
	d.Set("default_chat_service_role_sid", getResponse.DefaultChatServiceRoleSid)
	d.Set("default_conversation_creator_role_sid", getResponse.DefaultConversationCreatorRoleSid)
	d.Set("default_conversation_role_sid", getResponse.DefaultConversationRoleSid)
	d.Set("reachability_enabled", getResponse.ReachabilityEnabled)
	d.Set("url", getResponse.URL)

	return nil
}
