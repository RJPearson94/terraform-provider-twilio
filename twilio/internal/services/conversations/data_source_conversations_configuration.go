package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsConfigurationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this configuration",
			},
			"default_service_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default conversations service",
			},
			"default_closed_timer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default ISO 8601 duration after which a conversation will be automatically closed",
			},
			"default_inactive_timer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default ISO 8601 duration after which a conversation will be marked as inactive",
			},
			"default_messaging_service_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default messaging service",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the configuration resource",
			},
		},
	}
}

func dataSourceConversationsConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	getResponse, err := client.Configuration().FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation configuration was not found")
		}
		return diag.Errorf("Failed to read conversation configuration: %s", err.Error())
	}

	d.SetId(getResponse.AccountSid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("default_service_sid", getResponse.DefaultChatServiceSid)
	d.Set("default_closed_timer", getResponse.DefaultClosedTimer)
	d.Set("default_inactive_timer", getResponse.DefaultInactiveTimer)
	d.Set("default_messaging_service_sid", getResponse.DefaultMessagingServiceSid)
	d.Set("url", getResponse.URL)

	return nil
}
