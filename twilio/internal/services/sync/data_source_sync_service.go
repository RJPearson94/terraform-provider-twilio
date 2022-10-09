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
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acl_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reachability_debouncing_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"reachability_debouncing_window": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reachability_webhooks_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"webhook_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"webhooks_from_rest_enabled": {
				Type:     schema.TypeBool,
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
