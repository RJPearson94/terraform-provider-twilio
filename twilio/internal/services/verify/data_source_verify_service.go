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
				Description:  "The SID of the Verify service to fetch",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Verify service",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the Verify service",
			},
			"code_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The length of the verification code",
			},
			"custom_code_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether sending verifications with a custom code is enabled",
			},
			"do_not_share_warning_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether a warning not to share the verification code is included in the SMS body",
			},
			"dtmf_input_required": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the user must press a key to deliver the verification code via phone call",
			},
			"lookup_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether a phone number lookup is performed with each verification",
			},
			"mailer_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the mailer service associated with this Verify service",
			},
			"psd2_enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether PSD2 transaction parameters are passed when starting a verification",
			},
			"push": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Push notification configuration for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn_credential_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the Apple Push Notification Service (APN) credential",
						},
						"fcm_credential_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the Firebase Cloud Messaging (FCM) credential",
						},
					},
				},
			},
			"totp": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Time-based one-time password (TOTP) configuration for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name that appears in the user's authenticator app as the issuer of the TOTP code",
						},
						"time_step": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of seconds each TOTP code is valid for",
						},
						"code_length": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of digits in the generated TOTP code",
						},
						"skew": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of past and future time steps to allow during TOTP code validation",
						},
					},
				},
			},
			"default_template_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the default verification template used by this Verify service",
			},
			"skip_sms_to_landlines": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether SMS verifications to landlines are skipped",
			},
			"tts_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the text-to-speech voice used for phone call verifications",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Verify service was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the Verify service was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the Verify service resource",
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
