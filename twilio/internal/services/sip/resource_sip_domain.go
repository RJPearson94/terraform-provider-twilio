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
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"status_callback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"fallback_url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
					},
				},
			},
			"emergency": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"calling_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"caller_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"byoc_trunk_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secure": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"sip_registration": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"auth_type": {
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
		},
	}
}

func resourceSIPDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &domains.CreateDomainInput{
		DomainName:      d.Get("domain_name").(string),
		ByocTrunkSid:    utils.OptionalString(d, "byoc_trunk_sid"),
		FriendlyName:    utils.OptionalString(d, "friendly_name"),
		Secure:          utils.OptionalBool(d, "secure"),
		SipRegistration: utils.OptionalBool(d, "sip_registration"),
	}

	if _, ok := d.GetOk("voice"); ok {
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalString(d, "voice.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		createInput.VoiceStatusCallbackMethod = utils.OptionalString(d, "voice.0.status_callback_method")
		createInput.VoiceStatusCallbackURL = utils.OptionalString(d, "voice.0.status_callback_url")
		createInput.VoiceURL = utils.OptionalString(d, "voice.0.url")
	}

	if _, ok := d.GetOk("emergency"); ok {
		createInput.EmergencyCallerSid = utils.OptionalString(d, "emergency.0.caller_sid")
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
		ByocTrunkSid:    utils.OptionalString(d, "byoc_trunk_sid"),
		FriendlyName:    utils.OptionalString(d, "friendly_name"),
		Secure:          utils.OptionalBool(d, "secure"),
		SipRegistration: utils.OptionalBool(d, "sip_registration"),
	}

	if _, ok := d.GetOk("voice"); ok {
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.voice_fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalString(d, "voice.0.voice_fallback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "voice.0.voice_method")
		updateInput.VoiceStatusCallbackMethod = utils.OptionalString(d, "voice.0.voice_status_callback_method")
		updateInput.VoiceStatusCallbackURL = utils.OptionalString(d, "voice.0.voice_status_callback_url")
		updateInput.VoiceURL = utils.OptionalString(d, "voice.0.voice_url")
	}

	if _, ok := d.GetOk("emergency"); ok {
		updateInput.EmergencyCallerSid = utils.OptionalString(d, "emergency.0.caller_sid")
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
