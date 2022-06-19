package verify

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/verify/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVerifyService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVerifyServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.VerifyServiceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"code_length": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custom_code_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"do_not_share_warning_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dtmf_input_required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lookup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"mailer_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"psd2_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"push": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn_credential_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fcm_credential_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"totp": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_step": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"code_length": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"skew": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"default_template_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skip_sms_to_landlines": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tts_name": {
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

func dataSourceVerifyServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Verify service with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read Verify service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("code_length", getResponse.CodeLength)
	d.Set("custom_code_enabled", getResponse.CustomCodeEnabled)
	d.Set("do_not_share_warning_enabled", getResponse.DoNotShareWarningEnabled)
	d.Set("dtmf_input_required", getResponse.DtmfInputRequired)
	d.Set("lookup_enabled", getResponse.LookupEnabled)
	d.Set("mailer_sid", getResponse.MailerSid)
	d.Set("psd2_enabled", getResponse.Psd2Enabled)
	d.Set("push", helper.FlattenPush(getResponse.Push))
	d.Set("totp", helper.FlattenTotp(getResponse.Totp))
	d.Set("skip_sms_to_landlines", getResponse.SkipSmsToLandlines)
	d.Set("tts_name", getResponse.TtsName)
	d.Set("default_template_sid", getResponse.DefaultTemplateSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)

	return nil
}
