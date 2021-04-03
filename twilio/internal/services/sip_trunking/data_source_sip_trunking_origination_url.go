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
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sip_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
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
