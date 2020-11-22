package phone_number

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/phone_number/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/incoming_phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/incoming_phone_numbers"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourcePhoneNumber() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePhoneNumberCreate,
		ReadContext:   resourcePhoneNumberRead,
		UpdateContext: resourcePhoneNumberUpdate,
		DeleteContext: resourcePhoneNumberDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/PhoneNumbers/(.*)"
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
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone_number": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"phone_number", "area_code"},
			},
			"area_code": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"phone_number", "area_code"},
			},
			"address_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address_requirements": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"beta": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capabilities": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax": {
							Deprecated: "Due to Twilio disabling Programmable Fax for some accounts the api no longer return the necessary data so support will be removed in the next version and only voice will be supported",
							Type:       schema.TypeBool,
							Computed:   true,
						},
						"sms": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"mms": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"voice": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"emergency_address_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"emergency_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"messaging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Optional: true,
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
			"trunk_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				MaxItems:      1,
				ConflictsWith: []string{"fax"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"caller_id_lookup": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
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
			"fax": {
				Type:          schema.TypeList,
				Deprecated:    "Due to Twilio disabling Programmable Fax for some accounts the api no longer return the necessary data so support will be removed in the next version and only voice will be supported",
				Optional:      true,
				Computed:      true,
				MaxItems:      1,
				ConflictsWith: []string{"voice"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:     schema.TypeString,
							Optional: true,
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
			"identity_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bundle_sid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"origin": {
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

func resourcePhoneNumberCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &incoming_phone_numbers.CreateIncomingPhoneNumberInput{
		AddressSid:           utils.OptionalString(d, "address_sid"),
		AreaCode:             utils.OptionalString(d, "area_code"),
		BundleSid:            utils.OptionalString(d, "bundle_sid"),
		EmergencyAddressSid:  utils.OptionalString(d, "emergency_address_sid"),
		EmergencyStatus:      utils.OptionalString(d, "emergency_status"),
		FriendlyName:         utils.OptionalString(d, "friendly_name"),
		IdentitySid:          utils.OptionalString(d, "identity_sid"),
		PhoneNumber:          utils.OptionalString(d, "phone_number"),
		StatusCallback:       utils.OptionalString(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		TrunkSid:             utils.OptionalString(d, "trunk_sid"),
	}

	if _, ok := d.GetOk("messaging"); ok {
		createInput.SmsApplicationSid = utils.OptionalString(d, "messaging.0.application_sid")
		createInput.SmsFallbackMethod = utils.OptionalString(d, "messaging.0.fallback_method")
		createInput.SmsFallbackURL = utils.OptionalString(d, "messaging.0.fallback_url")
		createInput.SmsMethod = utils.OptionalString(d, "messaging.0.method")
		createInput.SmsURL = utils.OptionalString(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		createInput.VoiceApplicationSid = utils.OptionalString(d, "voice.0.application_sid")
		createInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalString(d, "voice.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		createInput.VoiceReceiveMode = sdkUtils.String(d.Get("voice").(string))
		createInput.VoiceURL = utils.OptionalString(d, "voice.0.url")
	}

	if _, ok := d.GetOk("fax"); ok {
		createInput.VoiceApplicationSid = utils.OptionalString(d, "fax.0.application_sid")
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "fax.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalString(d, "fax.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "fax.0.method")
		createInput.VoiceReceiveMode = sdkUtils.String(d.Get("fax").(string))
		createInput.VoiceURL = utils.OptionalString(d, "fax.0.url")
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).IncomingPhoneNumbers.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create phone number %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourcePhoneNumberRead(ctx, d, meta)
}

func resourcePhoneNumberRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).IncomingPhoneNumber(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read phone number: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("address_sid", getResponse.AddressSid)
	d.Set("address_requirements", getResponse.AddressRequirements)
	d.Set("beta", getResponse.Beta)
	d.Set("bundle_sid", getResponse.BundleSid)
	d.Set("capabilities", helper.FlattenCapabilities(&getResponse.Capabilities))
	d.Set("emergency_address_sid", getResponse.EmergencyAddressSid)
	d.Set("emergency_status", getResponse.EmergencyStatus)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("identity_sid", getResponse.IdentitySid)
	d.Set("messaging", helper.FlattenMessaging(getResponse))
	d.Set("origin", getResponse.Origin)
	d.Set("phone_number", getResponse.PhoneNumber)
	d.Set("status", getResponse.Status)
	d.Set("status_callback_url", getResponse.StatusCallback)
	d.Set("status_callback_method", getResponse.StatusCallbackMethod)
	d.Set("trunk_sid", getResponse.TrunkSid)

	// Since Programmable fax has been disabled on some accounts voice receive mode is no
	// longer being returned
	if getResponse.VoiceReceiveMode == "" || getResponse.VoiceReceiveMode == "voice" {
		d.Set("voice", helper.FlattenVoice(getResponse))
	}
	if getResponse.VoiceReceiveMode == "fax" {
		d.Set("fax", helper.FlattenFax(getResponse))
	}

	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}

func resourcePhoneNumberUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &incoming_phone_number.UpdateIncomingPhoneNumberInput{
		AddressSid:           utils.OptionalString(d, "address_sid"),
		BundleSid:            utils.OptionalString(d, "bundle_sid"),
		EmergencyAddressSid:  utils.OptionalString(d, "emergency_address_sid"),
		EmergencyStatus:      utils.OptionalString(d, "emergency_status"),
		FriendlyName:         utils.OptionalString(d, "friendly_name"),
		IdentitySid:          utils.OptionalString(d, "identity_sid"),
		StatusCallback:       utils.OptionalString(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		TrunkSid:             utils.OptionalString(d, "trunk_sid"),
	}

	if _, ok := d.GetOk("messaging"); ok {
		updateInput.SmsApplicationSid = utils.OptionalString(d, "messaging.0.application_sid")
		updateInput.SmsFallbackMethod = utils.OptionalString(d, "messaging.0.fallback_method")
		updateInput.SmsFallbackURL = utils.OptionalString(d, "messaging.0.fallback_url")
		updateInput.SmsMethod = utils.OptionalString(d, "messaging.0.method")
		updateInput.SmsURL = utils.OptionalString(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		updateInput.VoiceApplicationSid = utils.OptionalString(d, "voice.0.application_sid")
		updateInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalString(d, "voice.0.fallback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		updateInput.VoiceReceiveMode = sdkUtils.String(d.Get("voice").(string))
		updateInput.VoiceURL = utils.OptionalString(d, "voice.0.url")
	}

	if _, ok := d.GetOk("fax"); ok {
		updateInput.VoiceApplicationSid = utils.OptionalString(d, "fax.0.application_sid")
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "fax.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalString(d, "fax.0.fallback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "fax.0.method")
		updateInput.VoiceReceiveMode = sdkUtils.String(d.Get("fax").(string))
		updateInput.VoiceURL = utils.OptionalString(d, "fax.0.url")
	}

	updateResp, err := client.Account(d.Get("account_sid").(string)).IncomingPhoneNumber(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update phone number: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourcePhoneNumberRead(ctx, d, meta)
}

func resourcePhoneNumberDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).IncomingPhoneNumber(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete phone number: %s", err.Error())
	}

	d.SetId("")
	return nil
}
