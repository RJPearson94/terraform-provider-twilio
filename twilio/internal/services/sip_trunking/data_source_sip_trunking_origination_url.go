package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingOriginationURL() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingOriginationURLRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPOriginationURLValidation(),
				Description:  "The SID of the SIP trunk origination URL to look up",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this SIP trunk origination URL",
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk the origination URL belongs to",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the origination URL is enabled and available for use",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP trunk origination URL",
			},
			"priority": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the origination URL. Lower values have higher priority",
			},
			"sip_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SIP address to route origination calls to",
			},
			"weight": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The weight of the origination URL, used for load balancing among URLs with the same priority",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk origination URL was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk origination URL was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the SIP trunk origination URL resource",
			},
		},
	}
}

func dataSourceSIPTrunkingOriginationURLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Trunk(trunkSid).OriginationURL(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP trunk origination url with sid (%s) was not found for SIP trunk with sid (%s)", sid, trunkSid)
		}
		return diag.Errorf("Failed to read SIP trunk origination url: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("enabled", getResponse.Enabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("priority", getResponse.Priority)
	d.Set("sip_url", getResponse.SipURL)
	d.Set("trunk_sid", getResponse.TrunkSid)
	d.Set("weight", getResponse.Weight)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
