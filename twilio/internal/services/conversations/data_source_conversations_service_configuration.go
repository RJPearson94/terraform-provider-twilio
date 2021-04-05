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
			},
			"default_chat_service_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_conversation_creator_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_conversation_role_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reachability_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
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
