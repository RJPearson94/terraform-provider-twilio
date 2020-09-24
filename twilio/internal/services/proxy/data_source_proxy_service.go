package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceProxyService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chat_instance_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"chat_service_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"geo_match_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"number_selection_behavior": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"intercept_callback_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"out_of_session_callback_url": {
				Type:     schema.TypeString,
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

func dataSourceProxyServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Proxy service with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read proxy service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_instance_sid", getResponse.ChatInstanceSid)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
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
