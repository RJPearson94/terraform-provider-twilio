package proxy

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProxyService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProxyServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ProxyServiceSidValidation(),
				Description:  "The SID of the Proxy service to read",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Proxy service",
			},
			"chat_instance_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the Chat service instance associated with the Proxy service",
			},
			"unique_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique, developer-assigned name for the Proxy service",
			},
			"default_ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The default time-to-live (TTL) for sessions in the Proxy service, in seconds",
			},
			"callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to receive callback events for the Proxy service",
			},
			"geo_match_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The geographic area for matching proxy numbers",
			},
			"number_selection_behavior": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The behavior for selecting proxy numbers",
			},
			"intercept_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to receive intercept callback events for the Proxy service",
			},
			"out_of_session_callback_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL to receive out-of-session callback events for the Proxy service",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Proxy service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Proxy service resource",
			},
		},
	}
}

func dataSourceProxyServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Proxy service with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read proxy service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_instance_sid", getResponse.ChatInstanceSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("default_ttl", getResponse.DefaultTtl)
	d.Set("callback_url", getResponse.CallbackURL)
	d.Set("geo_match_level", getResponse.GeoMatchLevel)
	d.Set("number_selection_behavior", getResponse.NumberSelectionBehavior)
	d.Set("intercept_callback_url", getResponse.InterceptCallbackURL)
	d.Set("out_of_session_callback_url", getResponse.OutOfSessionCallbackURL)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
