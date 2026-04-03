package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
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
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to retrieve origination URLs for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the SIP trunk origination URLs",
			},
			"origination_urls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of origination URLs associated with the SIP trunk",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this SIP trunk origination URL by Twilio",
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
