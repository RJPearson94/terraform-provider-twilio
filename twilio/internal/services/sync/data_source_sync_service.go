package sync

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSyncService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSyncServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SyncServiceSidValidation(),
				Description:  "The SID of the Sync service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Sync service",
			},
			"acl_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether token identities in the Sync service must be granted access to Sync objects via the Permissions API",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the Sync service",
			},
			"reachability_debouncing_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether every endpoint_disconnected event fires after a configurable delay",
			},
			"reachability_debouncing_window": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The reachability event delay, in milliseconds",
			},
			"reachability_webhooks_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the service instance calls the webhook_url when client endpoints connect or disconnect from Sync",
			},
			"webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to which Sync sends webhooks",
			},
			"webhooks_from_rest_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the service instance calls the webhook_url when the REST API is used to update Sync objects",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Sync service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Sync service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Sync service resource",
			},
		},
	}
}

func dataSourceSyncServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Sync

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Sync service with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read Sync service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("acl_enabled", getResponse.AclEnabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("reachability_debouncing_enabled", getResponse.ReachabilityDebouncingEnabled)
	d.Set("reachability_debouncing_window", getResponse.ReachabilityDebouncingWindow)
	d.Set("reachability_webhooks_enabled", getResponse.ReachabilityWebhooksEnabled)
	d.Set("webhook_url", getResponse.WebhookURL)
	d.Set("webhooks_from_rest_enabled", getResponse.WebhooksFromRestEnabled)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
