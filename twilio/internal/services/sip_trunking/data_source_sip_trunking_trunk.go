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
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to look up",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this SIP trunk",
			},
			"cnam_lookup_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether CNAM (Caller Name) lookup is enabled for the trunk",
			},
			"disaster_recovery_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The HTTP method used to call the disaster recovery URL",
			},
			"disaster_recovery_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL called in the event of a disaster recovery failover",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique domain name for the SIP trunk",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the SIP trunk",
			},
			"recording": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The recording settings for the SIP trunk",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recording mode for the SIP trunk",
						},
						"trim": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether silence is trimmed from recordings",
						},
					},
				},
			},
			"secure": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether secure SIP (SIPS) is required for the trunk",
			},
			"transfer_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The call transfer mode for the SIP trunk",
			},
			"auth_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication type configured for the SIP trunk",
			},
			"auth_type_set": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The set of authentication types configured for the SIP trunk",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP trunk was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the SIP trunk resource",
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
