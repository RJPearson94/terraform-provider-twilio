package phone_number

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/phone_number/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/local"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/mobile"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/toll_free"
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
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
				ExactlyOneOf: []string{"phone_number", "area_code", "search_criteria"},
				ValidateFunc: utils.PhoneNumberValidation(),
			},
			"area_code": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"phone_number", "area_code", "search_criteria"},
			},
			"search_criteria": {
				Type:         schema.TypeList,
				Optional:     true,
				ForceNew:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"phone_number", "area_code", "search_criteria"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							ValidateFunc: validation.StringInSlice([]string{
								"local",
								"mobile",
								"toll_free",
							}, false),
						},
						"iso_country": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"area_code": {
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
						},
						"allow_beta_numbers": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"contains_number_pattern": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"exclude_address_requirements": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"all": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"local": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"foreign": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"location": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"in_postal_code": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"in_region": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"in_lata": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"in_locality": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"in_rate_center": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"near_number": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"near_lat_long": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
									},
									"distance": {
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
						"capabilities": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"sms_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"mms_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
									"voice_enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
									},
								},
							},
						},
					},
				},
			},
			"address_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.AddressSidValidation(),
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
							Type:     schema.TypeBool,
							Computed: true,
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
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.AddressSidValidation(),
			},
			"emergency_status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Active",
					"Inactive",
				}, false),
			},
			"messaging": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.ApplicationSidValidation(),
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
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
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
					},
				},
			},
			"trunk_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.SIPTrunkSidValidation(),
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
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.ApplicationSidValidation(),
						},
						"caller_id_lookup": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
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
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
					},
				},
			},
			"fax": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				MaxItems:      1,
				ConflictsWith: []string{"voice"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"application_sid": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: utils.ApplicationSidValidation(),
						},
						"fallback_method": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "POST",
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
							Default:  "POST",
							ValidateFunc: validation.StringInSlice([]string{
								"GET",
								"POST",
							}, false),
						},
						"url": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.IsURLWithHTTPorHTTPS,
						},
					},
				},
			},
			"identity_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.IdentitySidValidation(),
			},
			"bundle_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: utils.BundleSidValidation(),
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
				Default:  "POST",
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
		AddressSid:           utils.OptionalStringWithEmptyStringOnChange(d, "address_sid"),
		BundleSid:            utils.OptionalStringWithEmptyStringOnChange(d, "bundle_sid"),
		EmergencyAddressSid:  utils.OptionalStringWithEmptyStringOnChange(d, "emergency_address_sid"),
		EmergencyStatus:      utils.OptionalString(d, "emergency_status"),
		FriendlyName:         utils.OptionalString(d, "friendly_name"),
		IdentitySid:          utils.OptionalStringWithEmptyStringOnChange(d, "identity_sid"),
		StatusCallback:       utils.OptionalStringWithEmptyStringOnChange(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		TrunkSid:             utils.OptionalStringWithEmptyStringOnChange(d, "trunk_sid"),
	}

	if _, ok := d.GetOk("area_code"); ok {
		createInput.AreaCode = utils.OptionalString(d, "area_code")
	}

	if _, ok := d.GetOk("phone_number"); ok {
		createInput.PhoneNumber = utils.OptionalString(d, "phone_number")
	}

	if _, ok := d.GetOk("search_criteria"); ok {
		phoneNumber, err := searchForPhoneNumber(ctx, d, meta)
		if err != nil {
			return err
		}
		createInput.PhoneNumber = phoneNumber
	}

	if _, ok := d.GetOk("messaging"); ok {
		createInput.SmsApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.application_sid")
		createInput.SmsFallbackMethod = utils.OptionalString(d, "messaging.0.fallback_method")
		createInput.SmsFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_url")
		createInput.SmsMethod = utils.OptionalString(d, "messaging.0.method")
		createInput.SmsURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		createInput.VoiceReceiveMode = sdkUtils.String("voice")
		createInput.VoiceApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.application_sid")
		createInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		createInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
	}

	if _, ok := d.GetOk("fax"); ok {
		createInput.VoiceReceiveMode = sdkUtils.String("fax")
		createInput.VoiceApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.application_sid")
		createInput.VoiceFallbackMethod = utils.OptionalString(d, "fax.0.fallback_method")
		createInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.fallback_url")
		createInput.VoiceMethod = utils.OptionalString(d, "fax.0.method")
		createInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.url")
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

	if helper.IsVoiceReceiveMode(getResponse.VoiceReceiveMode) {
		d.Set("voice", helper.FlattenVoice(getResponse))
		d.Set("fax", &[]interface{}{})
	} else {
		d.Set("fax", helper.FlattenFax(getResponse))
		d.Set("voice", &[]interface{}{})
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
		AddressSid:           utils.OptionalStringWithEmptyStringOnChange(d, "address_sid"),
		BundleSid:            utils.OptionalStringWithEmptyStringOnChange(d, "bundle_sid"),
		EmergencyAddressSid:  utils.OptionalStringWithEmptyStringOnChange(d, "emergency_address_sid"),
		EmergencyStatus:      utils.OptionalString(d, "emergency_status"),
		FriendlyName:         utils.OptionalString(d, "friendly_name"),
		IdentitySid:          utils.OptionalStringWithEmptyStringOnChange(d, "identity_sid"),
		StatusCallback:       utils.OptionalStringWithEmptyStringOnChange(d, "status_callback_url"),
		StatusCallbackMethod: utils.OptionalString(d, "status_callback_method"),
		TrunkSid:             utils.OptionalStringWithEmptyStringOnChange(d, "trunk_sid"),
	}

	if _, ok := d.GetOk("messaging"); ok {
		updateInput.SmsApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.application_sid")
		updateInput.SmsFallbackMethod = utils.OptionalString(d, "messaging.0.fallback_method")
		updateInput.SmsFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.fallback_url")
		updateInput.SmsMethod = utils.OptionalString(d, "messaging.0.method")
		updateInput.SmsURL = utils.OptionalStringWithEmptyStringOnChange(d, "messaging.0.url")
	}

	if _, ok := d.GetOk("voice"); ok {
		updateInput.VoiceReceiveMode = sdkUtils.String("voice")
		updateInput.VoiceApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.application_sid")
		updateInput.VoiceCallerIDLookup = utils.OptionalBool(d, "voice.0.caller_id_lookup")
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "voice.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.fallback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "voice.0.method")
		updateInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "voice.0.url")
	}

	if _, ok := d.GetOk("fax"); ok {
		updateInput.VoiceReceiveMode = sdkUtils.String("fax")
		updateInput.VoiceApplicationSid = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.application_sid")
		updateInput.VoiceFallbackMethod = utils.OptionalString(d, "fax.0.fallback_method")
		updateInput.VoiceFallbackURL = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.fallback_url")
		updateInput.VoiceMethod = utils.OptionalString(d, "fax.0.method")
		updateInput.VoiceURL = utils.OptionalStringWithEmptyStringOnChange(d, "fax.0.url")
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

func searchForPhoneNumber(ctx context.Context, d *schema.ResourceData, meta interface{}) (*string, diag.Diagnostics) {
	typeOfPhoneNumber := d.Get("search_criteria.0.type")

	if typeOfPhoneNumber == "local" {
		return searchForLocalNumber(ctx, d, meta)
	}
	if typeOfPhoneNumber == "mobile" {
		return searchForMobileNumber(ctx, d, meta)
	}
	return searchForTollFreeNumber(ctx, d, meta)
}

func searchForLocalNumber(ctx context.Context, d *schema.ResourceData, meta interface{}) (*string, diag.Diagnostics) {
	client := meta.(*common.TwilioClient).API

	pageOptions := populateAvailablePhoneNumberPageOptions(d)
	options := &local.AvailablePhoneNumbersPageOptions{
		AreaCode:                      pageOptions.AreaCode,
		Beta:                          pageOptions.Beta,
		Contains:                      pageOptions.Contains,
		PageSize:                      pageOptions.PageSize,
		ExcludeAllAddressRequired:     pageOptions.ExcludeAllAddressRequired,
		ExcludeLocalAddressRequired:   pageOptions.ExcludeLocalAddressRequired,
		ExcludeForeignAddressRequired: pageOptions.ExcludeForeignAddressRequired,
		FaxEnabled:                    pageOptions.FaxEnabled,
		SmsEnabled:                    pageOptions.SmsEnabled,
		MmsEnabled:                    pageOptions.MmsEnabled,
		VoiceEnabled:                  pageOptions.VoiceEnabled,
		NearNumber:                    pageOptions.NearNumber,
		NearLatLong:                   pageOptions.NearLatLong,
		Distance:                      pageOptions.Distance,
		InPostalCode:                  pageOptions.InPostalCode,
		InRegion:                      pageOptions.InRegion,
		InRateCenter:                  pageOptions.InRateCenter,
		InLata:                        pageOptions.InLata,
		InLocality:                    pageOptions.InLocality,
	}

	accountSid := d.Get("account_sid").(string)
	countryCode := d.Get("search_criteria.0.iso_country").(string)
	pageResponse, err := client.Account(accountSid).AvailablePhoneNumber(countryCode).Local.PageWithContext(ctx, options)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return nil, diag.Errorf("No local phone numbers were found for country (%s) in account (%s)", countryCode, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return nil, diag.Errorf("Failed to list available local phone numbers: %s", err.Error())
	}

	availableNumbers := pageResponse.AvailablePhoneNumbers
	if len(availableNumbers) == 0 {
		return nil, diag.Errorf("No local phone numbers have been found")
	}
	return sdkUtils.String(availableNumbers[0].PhoneNumber), nil
}

func searchForMobileNumber(ctx context.Context, d *schema.ResourceData, meta interface{}) (*string, diag.Diagnostics) {
	client := meta.(*common.TwilioClient).API

	pageOptions := populateAvailablePhoneNumberPageOptions(d)
	options := &mobile.AvailablePhoneNumbersPageOptions{
		AreaCode:                      pageOptions.AreaCode,
		Beta:                          pageOptions.Beta,
		Contains:                      pageOptions.Contains,
		PageSize:                      pageOptions.PageSize,
		ExcludeAllAddressRequired:     pageOptions.ExcludeAllAddressRequired,
		ExcludeLocalAddressRequired:   pageOptions.ExcludeLocalAddressRequired,
		ExcludeForeignAddressRequired: pageOptions.ExcludeForeignAddressRequired,
		FaxEnabled:                    pageOptions.FaxEnabled,
		SmsEnabled:                    pageOptions.SmsEnabled,
		MmsEnabled:                    pageOptions.MmsEnabled,
		VoiceEnabled:                  pageOptions.VoiceEnabled,
		NearNumber:                    pageOptions.NearNumber,
		NearLatLong:                   pageOptions.NearLatLong,
		Distance:                      pageOptions.Distance,
		InPostalCode:                  pageOptions.InPostalCode,
		InRegion:                      pageOptions.InRegion,
		InRateCenter:                  pageOptions.InRateCenter,
		InLata:                        pageOptions.InLata,
		InLocality:                    pageOptions.InLocality,
	}

	accountSid := d.Get("account_sid").(string)
	countryCode := d.Get("search_criteria.0.iso_country").(string)
	pageResponse, err := client.Account(accountSid).AvailablePhoneNumber(countryCode).Mobile.PageWithContext(ctx, options)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return nil, diag.Errorf("No mobile phone numbers were found for country (%s) in account (%s)", countryCode, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return nil, diag.Errorf("Failed to list available mobile phone numbers: %s", err.Error())
	}

	availableNumbers := pageResponse.AvailablePhoneNumbers
	if len(availableNumbers) == 0 {
		return nil, diag.Errorf("No mobile phone numbers have been found")
	}
	return sdkUtils.String(availableNumbers[0].PhoneNumber), nil
}

func searchForTollFreeNumber(ctx context.Context, d *schema.ResourceData, meta interface{}) (*string, diag.Diagnostics) {
	client := meta.(*common.TwilioClient).API

	pageOptions := populateAvailablePhoneNumberPageOptions(d)
	options := &toll_free.AvailablePhoneNumbersPageOptions{
		AreaCode:                      pageOptions.AreaCode,
		Beta:                          pageOptions.Beta,
		Contains:                      pageOptions.Contains,
		PageSize:                      pageOptions.PageSize,
		ExcludeAllAddressRequired:     pageOptions.ExcludeAllAddressRequired,
		ExcludeLocalAddressRequired:   pageOptions.ExcludeLocalAddressRequired,
		ExcludeForeignAddressRequired: pageOptions.ExcludeForeignAddressRequired,
		FaxEnabled:                    pageOptions.FaxEnabled,
		SmsEnabled:                    pageOptions.SmsEnabled,
		MmsEnabled:                    pageOptions.MmsEnabled,
		VoiceEnabled:                  pageOptions.VoiceEnabled,
		NearNumber:                    pageOptions.NearNumber,
		NearLatLong:                   pageOptions.NearLatLong,
		Distance:                      pageOptions.Distance,
		InPostalCode:                  pageOptions.InPostalCode,
		InRegion:                      pageOptions.InRegion,
		InRateCenter:                  pageOptions.InRateCenter,
		InLata:                        pageOptions.InLata,
		InLocality:                    pageOptions.InLocality,
	}

	accountSid := d.Get("account_sid").(string)
	countryCode := d.Get("search_criteria.0.iso_country").(string)
	pageResponse, err := client.Account(accountSid).AvailablePhoneNumber(countryCode).TollFree.PageWithContext(ctx, options)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return nil, diag.Errorf("No toll free phone numbers were found for country (%s) in account (%s)", countryCode, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return nil, diag.Errorf("Failed to list available toll free phone numbers: %s", err.Error())
	}

	availableNumbers := pageResponse.AvailablePhoneNumbers
	if len(availableNumbers) == 0 {
		return nil, diag.Errorf("No toll free phone numbers have been found")
	}
	return sdkUtils.String(availableNumbers[0].PhoneNumber), nil
}

type AvailablePhoneNumbersPageOptions struct {
	PageSize                      *int
	AreaCode                      *int
	Contains                      *string
	SmsEnabled                    *bool
	MmsEnabled                    *bool
	VoiceEnabled                  *bool
	ExcludeAllAddressRequired     *bool
	ExcludeLocalAddressRequired   *bool
	ExcludeForeignAddressRequired *bool
	Beta                          *bool
	NearNumber                    *string
	NearLatLong                   *string
	Distance                      *int
	InPostalCode                  *string
	InRegion                      *string
	InRateCenter                  *string
	InLata                        *string
	InLocality                    *string
	FaxEnabled                    *bool
}

func populateAvailablePhoneNumberPageOptions(d *schema.ResourceData) *AvailablePhoneNumbersPageOptions {
	options := &AvailablePhoneNumbersPageOptions{
		AreaCode: utils.OptionalInt(d, "search_criteria.0.area_code"),
		Beta:     utils.OptionalBool(d, "search_criteria.0.allow_beta_numbers"),
		Contains: utils.OptionalString(d, "search_criteria.0.contains_number_pattern"),
		PageSize: sdkUtils.Int(1), // We only need 1 phone number to purchase
	}

	if _, ok := d.GetOk("search_criteria.0.exclude_address_requirements"); ok {
		options.ExcludeAllAddressRequired = utils.OptionalBool(d, "search_criteria.0.exclude_address_requirements.0.all")
		options.ExcludeLocalAddressRequired = utils.OptionalBool(d, "search_criteria.0.exclude_address_requirements.0.local")
		options.ExcludeForeignAddressRequired = utils.OptionalBool(d, "search_criteria.0.exclude_address_requirements.0.foreign")
	}

	if _, ok := d.GetOk("search_criteria.0.capabilities"); ok {
		options.FaxEnabled = utils.OptionalBool(d, "search_criteria.0.capabilities.0.fax_enabled")
		options.SmsEnabled = utils.OptionalBool(d, "search_criteria.0.capabilities.0.sms_enabled")
		options.MmsEnabled = utils.OptionalBool(d, "search_criteria.0.capabilities.0.mms_enabled")
		options.VoiceEnabled = utils.OptionalBool(d, "search_criteria.0.capabilities.0.voice_enabled")
	}

	if _, ok := d.GetOk("search_criteria.0.location"); ok {
		options.NearNumber = utils.OptionalString(d, "search_criteria.0.location.0.near_number")
		options.NearLatLong = utils.OptionalString(d, "search_criteria.0.location.0.near_lat_long")
		options.Distance = utils.OptionalInt(d, "search_criteria.0.location.0.distance")
		options.InPostalCode = utils.OptionalString(d, "search_criteria.0.location.0.in_postal_code")
		options.InRegion = utils.OptionalString(d, "search_criteria.0.location.0.in_region")
		options.InRateCenter = utils.OptionalString(d, "search_criteria.0.location.0.in_rate_center")
		options.InLata = utils.OptionalString(d, "search_criteria.0.location.0.in_lata")
		options.InLocality = utils.OptionalString(d, "search_criteria.0.location.0.in_locality")
	}
	return options
}
