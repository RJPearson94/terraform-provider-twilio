package conversations

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConversationsAddressConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConversationsAddressConfigurationRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ConversationAddressConfigurationSidValidation(),
				Description:  "The SID of the address configuration to retrieve",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this address configuration",
			},
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The address (e.g. phone number) for the configuration",
			},
			"auto_creation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The auto-creation settings for the address configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the conversations service",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether auto-creation is enabled",
						},
						"flow_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the Studio flow to trigger",
						},
						"retry_count": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of times to retry the Studio flow",
						},
						"integration_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of auto-creation integration",
						},
						"webhook_filters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of webhook event filters",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"webhook_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method for the webhook",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for the webhook",
						},
					},
				},
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the address configuration",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of address",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the address configuration was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the address configuration was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the address configuration resource",
			},
		},
	}
}

func dataSourceConversationsAddressConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Conversations

	sid := d.Get("sid").(string)
	getResponse, err := client.Configuration().Address(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Conversation address configuration with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read conversation address configuration: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("address", getResponse.Address)
	d.Set("auto_creation", &[]interface{}{
		map[string]interface{}{
			"service_sid":     getResponse.AutoCreation.ConversationServiceSid,
			"enabled":         getResponse.AutoCreation.Enabled,
			"flow_sid":        getResponse.AutoCreation.StudioFlowSid,
			"retry_count":     getResponse.AutoCreation.StudioRetryCount,
			"type":            getResponse.AutoCreation.Type,
			"webhook_filters": getResponse.AutoCreation.WebhookFilters,
			"webhook_method":  getResponse.AutoCreation.WebhookMethod,
			"webhook_url":     getResponse.AutoCreation.WebhookUrl,
		},
	})
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("type", getResponse.Type)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
