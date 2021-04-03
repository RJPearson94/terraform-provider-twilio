package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingOriginationURLs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingOriginationURLsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: helper.TrunkSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"origination_urls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceSIPTrunkingOriginationURLsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	paginator := client.Trunk(trunkSid).OriginationURLs.NewOriginationURLsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No origination urls were found for SIP trunk with sid (%s)", trunkSid)
		}
		return diag.Errorf("Failed to list SIP trunk origination urls: %s", err.Error())
	}

	d.SetId(trunkSid)
	d.Set("trunk_sid", trunkSid)

	originationURLs := make([]interface{}, 0)

	for _, originationURL := range paginator.OriginationURLs {
		d.Set("account_sid", originationURL.AccountSid)

		originationURLMap := make(map[string]interface{})

		originationURLMap["sid"] = originationURL.Sid
		originationURLMap["enabled"] = originationURL.Enabled
		originationURLMap["friendly_name"] = originationURL.FriendlyName
		originationURLMap["priority"] = originationURL.Priority
		originationURLMap["sip_url"] = originationURL.SipURL
		originationURLMap["weight"] = originationURL.Weight
		originationURLMap["date_created"] = originationURL.DateCreated.Format(time.RFC3339)

		if originationURL.DateUpdated != nil {
			originationURLMap["date_updated"] = originationURL.DateUpdated.Format(time.RFC3339)
		}

		originationURLMap["url"] = originationURL.URL

		originationURLs = append(originationURLs, originationURLMap)
	}

	d.Set("origination_urls", &originationURLs)

	return nil
}
