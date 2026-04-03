package verify

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/verify/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/services"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVerifyService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVerifyServiceCreate,
		ReadContext:   resourceVerifyServiceRead,
		UpdateContext: resourceVerifyServiceUpdate,
		DeleteContext: resourceVerifyServiceDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this Verify service by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Verify service",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
				Description:  "A human-readable label for the Verify service. Must be between 1 and 30 characters",
			},
			"code_length": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(4, 10),
				Default:      6,
				Description:  "The length of the verification code to generate. Must be between 4 and 10. Defaults to `6`",
			},
			"custom_code_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to allow sending verifications with a custom code instead of a randomly generated one. Defaults to `false`",
			},
			"do_not_share_warning_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to include a warning not to share the verification code in the SMS body. Defaults to `false`",
			},
			"dtmf_input_required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to require the user to press a key to deliver the verification code via phone call. Defaults to `true`",
			},
			"lookup_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to perform a phone number lookup with each verification. Defaults to `false`",
			},
			"mailer_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.MailerSidValidation(),
				Description:  "The SID of the mailer service to associate with this Verify service for email verifications",
			},
			"psd2_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to pass PSD2 transaction parameters when starting a verification. Defaults to `false`",
			},
			"push": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Push notification configuration for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"apn_credential_sid": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.CredentialSidValidation(),
							Description:  "The SID of the Apple Push Notification Service (APN) credential to use for push notifications",
						},
						"fcm_credential_sid": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.CredentialSidValidation(),
							Description:  "The SID of the Firebase Cloud Messaging (FCM) credential to use for push notifications",
						},
					},
				},
			},
			"totp": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Time-based one-time password (TOTP) configuration for the Verify service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"issuer": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The name that appears in the user's authenticator app as the issuer of the TOTP code",
						},
						"time_step": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      30,
							ValidateFunc: validation.IntBetween(20, 60),
							Description:  "The number of seconds each TOTP code is valid for. Must be between 20 and 60. Defaults to `30`",
						},
						"code_length": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      6,
							ValidateFunc: validation.IntBetween(3, 8),
							Description:  "The number of digits in the generated TOTP code. Must be between 3 and 8. Defaults to `6`",
						},
						"skew": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      1,
							ValidateFunc: validation.IntBetween(0, 2),
							Description:  "The number of past and future time steps to allow during TOTP code validation. Must be between 0 and 2. Defaults to `1`",
						},
					},
				},
			},
			"default_template_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.VerifyTemplateSidValidation(),
				Description:  "The SID of the default verification template to use for this Verify service",
			},
			"skip_sms_to_landlines": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to skip sending SMS verifications to landlines. Defaults to `false`",
			},
			"tts_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the text-to-speech voice to use for phone call verifications",
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

func resourceVerifyServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	createInput := &services.CreateServiceInput{
		FriendlyName:             d.Get("friendly_name").(string),
		CodeLength:               utils.OptionalInt(d, "code_length"),
		CustomCodeEnabled:        utils.OptionalBool(d, "custom_code_enabled"),
		DefaultTemplateSid:       utils.OptionalStringWithEmptyStringOnChange(d, "default_template_sid"),
		DoNotShareWarningEnabled: utils.OptionalBool(d, "do_not_share_warning_enabled"),
		DtmfInputRequired:        utils.OptionalBool(d, "dtmf_input_required"),
		LookupEnabled:            utils.OptionalBool(d, "lookup_enabled"),
		Psd2Enabled:              utils.OptionalBool(d, "psd2_enabled"),
		SkipSmsToLandlines:       utils.OptionalBool(d, "skip_sms_to_landlines"),
		TtsName:                  utils.OptionalStringWithEmptyStringOnChange(d, "tts_name"),
		MailerSid:                utils.OptionalStringWithEmptyStringOnChange(d, "mailer_sid"),
	}

	if _, ok := d.GetOk("push"); ok {
		createInput.Push = &services.CreateServicePushInput{
			ApnCredentialSid: utils.OptionalStringWithEmptyStringOnChange(d, "push.0.apn_credential_sid"),
			FcmCredentialSid: utils.OptionalStringWithEmptyStringOnChange(d, "push.0.fcm_credential_sid"),
		}
	}

	if _, ok := d.GetOk("totp"); ok {
		createInput.Totp = &services.CreateServiceTotpInput{
			Issuer:     utils.OptionalString(d, "totp.0.issuer"),
			TimeStep:   utils.OptionalInt(d, "totp.0.time_step"),
			CodeLength: utils.OptionalInt(d, "totp.0.code_length"),
			Skew:       sdkUtils.Int(d.Get("totp.0.skew").(int)), // This is set as GetOk cannot detect 0, this allow the skew to be set to 0 on creation. This can be set due to the resource default
		}
	}

	createResult, err := client.Services.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create service: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVerifyServiceRead(ctx, d, meta)
}

func resourceVerifyServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	getResponse, err := client.Service(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read service: %s", err.Error())
	}

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

func resourceVerifyServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	updateInput := &service.UpdateServiceInput{
		FriendlyName:             utils.OptionalString(d, "friendly_name"),
		CodeLength:               utils.OptionalInt(d, "code_length"),
		CustomCodeEnabled:        utils.OptionalBool(d, "custom_code_enabled"),
		DefaultTemplateSid:       utils.OptionalStringWithEmptyStringOnChange(d, "default_template_sid"),
		DoNotShareWarningEnabled: utils.OptionalBool(d, "do_not_share_warning_enabled"),
		DtmfInputRequired:        utils.OptionalBool(d, "dtmf_input_required"),
		LookupEnabled:            utils.OptionalBool(d, "lookup_enabled"),
		Psd2Enabled:              utils.OptionalBool(d, "psd2_enabled"),
		SkipSmsToLandlines:       utils.OptionalBool(d, "skip_sms_to_landlines"),
		TtsName:                  utils.OptionalStringWithEmptyStringOnChange(d, "tts_name"),
		MailerSid:                utils.OptionalStringWithEmptyStringOnChange(d, "mailer_sid"),
	}

	if _, ok := d.GetOk("push"); ok {
		updateInput.Push = &service.UpdateServicePushInput{
			ApnCredentialSid: utils.OptionalStringWithEmptyStringOnChange(d, "push.0.apn_credential_sid"),
			FcmCredentialSid: utils.OptionalStringWithEmptyStringOnChange(d, "push.0.fcm_credential_sid"),
		}
	}

	if _, ok := d.GetOk("totp"); ok {
		updateInput.Totp = &service.UpdateServiceTotpInput{
			Issuer:     utils.OptionalString(d, "totp.0.issuer"),
			TimeStep:   utils.OptionalInt(d, "totp.0.time_step"),
			CodeLength: utils.OptionalInt(d, "totp.0.code_length"),
			Skew:       utils.OptionalIntWith0OnChange(d, "totp.0.skew"),
		}
	}

	updateResp, err := client.Service(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update service: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVerifyServiceRead(ctx, d, meta)
}

func resourceVerifyServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Verify

	if err := client.Service(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete service: %s", err.Error())
	}

	d.SetId("")
	return nil
}
