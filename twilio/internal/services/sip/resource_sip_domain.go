package sip

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domains"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSIPDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPDomainCreate,
		ReadContext:   resourceSIPDomainRead,
		UpdateContext: resourceSIPDomainUpdate,
		DeleteContext: resourceSIPDomainDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/SIP/Domains/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
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
				Description: "The unique SID assigned to this SIP domain by Twilio",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this SIP domain. Changing this forces a new resource",
			},
			"domain_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile(`^[a-zA-Z0-9-.]+\.sip\.twilio\.com$`), ""),
				Description:  "The fully qualified domain name for the SIP domain, ending with `.sip.twilio.com`",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A human-readable label for the SIP domain",
			},
			"voice": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "A block to configure voice settings for the SIP domain",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
							Description:  "The URL to call for voice status callback events",
						},
						"status_callback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
							Description: "The HTTP method used to call the voice status callback URL. Valid values are `GET` or `POST`. Defaults to `POST`",
						},
						"fallback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
							Description:  "The URL to call when an error occurs retrieving or executing the TwiML for voice calls",
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
							Description: "The HTTP method used to call the voice fallback URL. Valid values are `GET` or `POST`. Defaults to `POST`",
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
							Description:  "The URL to call when the SIP domain receives an incoming voice call",
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
							Description: "The HTTP method used to call the voice URL. Valid values are `GET` or `POST`. Defaults to `POST`",
						},
					},
				},
			},
			"emergency": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "A block to configure emergency calling settings for the SIP domain",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calling_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether emergency calling is enabled for this SIP domain. Defaults to `false`",
						},
						"caller_sid": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.PhoneNumberSidValidation(),
							Description:  "The SID of the phone number to use as the emergency caller ID",
						},
					},
				},
			},
			"byoc_trunk_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.ByocSidValidation(),
				Description:  "The SID of the BYOC trunk to associate with this SIP domain",
			},
			"secure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether secure SIP (SIPS) is enabled for the domain. Defaults to `false`",
			},
			"sip_registration": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether SIP registration is allowed for the domain. Defaults to `false`",
			},
			"auth_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication type configured for the SIP domain",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the SIP domain was last updated, in RFC 3339 format",
			},
		},
	}
}

func resourceSIPDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &domains.CreateDomainInput{
		DomainName:      d.Get("domain_name").(string),
		ByocTrunkSid:    utils.OptionalStringWithEmptyStringOnChange(d, "byoc_trunk_sid"),
		FriendlyName:    utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		Secure:          utils.OptionalBool(d, "secure"),
		SipRegistration: utils.OptionalBool(d, "sip_registration"),
	}

	if _, ok := d.GetOk("voice"); ok {
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		createInput.VoiceStatusCallbackMethod = utils.OptionalString(d, "voice.0.status_callback_method")
		createInput.VoiceStatusCallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		createInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
	}

	if _, ok := d.GetOk("emergency"); ok {
		createInput.EmergencyCallerSid = utils.OptionalStringWithEmptyStringOnChange(d, "emergency.0.caller_sid")
		createInput.EmergencyCallingEnabled = utils.OptionalBool(d, "emergency.0.calling_enabled")
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Sip.Domains.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP domain: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceSIPDomainRead(ctx, d, meta)
}

func resourceSIPDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP domain: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("auth_type", getResponse.AuthType)
	d.Set("byoc_trunk_sid", getResponse.ByocTrunkSid)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("emergency", helper.FlattenEmergency(getResponse))
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("secure", getResponse.Secure)
	d.Set("sip_registration", getResponse.SipRegistration)
	d.Set("voice", helper.FlattenVoice(getResponse))
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}

func resourceSIPDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &domain.UpdateDomainInput{
		DomainName:      utils.OptionalString(d, "domain_name"),
		ByocTrunkSid:    utils.OptionalStringWithEmptyStringOnChange(d, "byoc_trunk_sid"),
		FriendlyName:    utils.OptionalStringWithEmptyStringOnChange(d, "friendly_name"),
		Secure:          utils.OptionalBool(d, "secure"),
		SipRegistration: utils.OptionalBool(d, "sip_registration"),
	}

	if _, ok := d.GetOk("voice"); ok {
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		updateInput.VoiceStatusCallbackMethod = utils.OptionalString(d, "voice.0.status_callback_method")
		updateInput.VoiceStatusCallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.status_callback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		updateInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
	}

	if _, ok := d.GetOk("emergency"); ok {
		updateInput.EmergencyCallerSid = utils.OptionalStringWithEmptyStringOnChange(d, "emergency.0.caller_sid")
		updateInput.EmergencyCallingEnabled = utils.OptionalBool(d, "emergency.0.calling_enabled")
	}

	updateResult, err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP domain: %s", err.Error())
	}

	d.SetId(updateResult.Sid)
	return resourceSIPDomainRead(ctx, d, meta)
}

func resourceSIPDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Sip.Domain(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP domain: %s", err.Error())
	}
	d.SetId("")
	return nil
}
