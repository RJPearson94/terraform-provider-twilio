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

func dataSourceSIPTrunkingTrunk() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingTrunkRead,

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
			"cnam_lookup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disaster_recovery_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disaster_recovery_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recording": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trim": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"secure": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"transfer_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type_set": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func dataSourceSIPTrunkingTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	sid := d.Get("sid").(string)
	getResponse, err := client.Trunk(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("SIP trunk with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read SIP trunk: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("cnam_lookup_enabled", getResponse.CnamLookupEnabled)
	d.Set("disaster_recovery_method", getResponse.DisasterRecoveryMethod)
	d.Set("disaster_recovery_url", getResponse.DisasterRecoveryURL)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("recording", helper.FlattenRecording(getResponse.Recording))
	d.Set("secure", getResponse.Secure)
	d.Set("transfer_mode", getResponse.TransferMode)
	d.Set("auth_type", getResponse.AuthType)
	d.Set("auth_type_set", getResponse.AuthTypeSet)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
