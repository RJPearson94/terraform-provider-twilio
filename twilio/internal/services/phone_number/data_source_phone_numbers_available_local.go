package phone_number

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/local"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePhoneNumberAvailableLocalNumbers() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "As data sources are read at the plan phase and retrive a new list of available phone numbers, this data source cannot be used to purchase a phone number. Please use the `search_criteria` block on the `twilio_phone_number` resource instead. The data source will be removed in a future release",

		ReadContext: dataSourcePhoneNumberAvailableLocalNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account to search for available local phone numbers",
			},
			"iso_country": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The ISO 3166-1 alpha-2 country code to search for available local phone numbers",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum number of results to return",
			},
			"area_code": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The area code to filter available phone numbers by",
			},
			"allow_beta_numbers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to include beta phone numbers in the search results",
			},
			"contains_number_pattern": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A pattern to match phone numbers against, using '*' as a wildcard",
			},
			"exclude_address_requirements": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A block to exclude phone numbers that require an address",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to exclude phone numbers that require any address",
						},
						"local": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to exclude phone numbers that require a local address",
						},
						"foreign": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to exclude phone numbers that require a foreign address",
						},
					},
				},
			},
			"location": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A block to filter available phone numbers by location",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"in_postal_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The postal code to filter available phone numbers by",
						},
						"in_region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The region (state or province) to filter available phone numbers by",
						},
						"in_lata": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The LATA to filter available phone numbers by",
						},
						"in_locality": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The locality (city) to filter available phone numbers by",
						},
						"in_rate_center": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The rate center to filter available phone numbers by",
						},
						"near_number": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A phone number to search near",
						},
						"near_lat_long": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A latitude/longitude coordinate pair to search near, specified as `latitude,longitude`",
						},
						"distance": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The distance in miles from the `near_number` or `near_lat_long` to search within",
						},
					},
				},
			},
			"capabilities": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A block to filter available phone numbers by their capabilities",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to filter for fax-capable phone numbers",
						},
						"sms_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to filter for SMS-capable phone numbers",
						},
						"mms_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to filter for MMS-capable phone numbers",
						},
						"voice_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to filter for voice-capable phone numbers",
						},
					},
				},
			},
			"available_phone_numbers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of available local phone numbers matching the search criteria",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the available phone number",
						},
						"phone_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The available phone number in E.164 format",
						},
						"address_requirements": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of address required for this phone number",
						},
						"beta": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number is a beta number new to the Twilio platform",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The set of boolean capabilities of the phone number",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number supports fax",
									},
									"sms": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number supports SMS",
									},
									"mms": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number supports MMS",
									},
									"voice": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number supports voice",
									},
								},
							},
						},
						"lata": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The LATA of the phone number",
						},
						"rate_center": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rate center of the phone number",
						},
						"latitude": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latitude of the phone number's location",
						},
						"longitude": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The longitude of the phone number's location",
						},
						"locality": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The locality (city) of the phone number",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region (state or province) of the phone number",
						},
						"postal_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The postal code of the phone number",
						},
					},
				},
			},
		},
	}
}

func dataSourcePhoneNumberAvailableLocalNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	options := &local.AvailablePhoneNumbersPageOptions{
		AreaCode: utils.OptionalInt(d, "area_code"),
		Beta:     utils.OptionalBool(d, "allow_beta_numbers"),
		Contains: utils.OptionalString(d, "contains_number_pattern"),
		PageSize: utils.OptionalInt(d, "limit"),
	}

	if _, ok := d.GetOk("exclude_address_requirements"); ok {
		options.ExcludeAllAddressRequired = utils.OptionalBool(d, "exclude_address_requirements.0.all")
		options.ExcludeLocalAddressRequired = utils.OptionalBool(d, "exclude_address_requirements.0.local")
		options.ExcludeForeignAddressRequired = utils.OptionalBool(d, "exclude_address_requirements.0.foreign")
	}

	if _, ok := d.GetOk("capabilities"); ok {
		options.FaxEnabled = utils.OptionalBool(d, "capabilities.0.fax_enabled")
		options.SmsEnabled = utils.OptionalBool(d, "capabilities.0.sms_enabled")
		options.MmsEnabled = utils.OptionalBool(d, "capabilities.0.mms_enabled")
		options.VoiceEnabled = utils.OptionalBool(d, "capabilities.0.voice_enabled")
	}

	if _, ok := d.GetOk("location"); ok {
		options.NearNumber = utils.OptionalString(d, "location.0.near_number")
		options.NearLatLong = utils.OptionalString(d, "location.0.near_lat_long")
		options.Distance = utils.OptionalInt(d, "location.0.distance")
		options.InPostalCode = utils.OptionalString(d, "location.0.in_postal_code")
		options.InRegion = utils.OptionalString(d, "location.0.in_region")
		options.InRateCenter = utils.OptionalString(d, "location.0.in_rate_center")
		options.InLata = utils.OptionalString(d, "location.0.in_lata")
		options.InLocality = utils.OptionalString(d, "location.0.in_locality")
	}

	accountSid := d.Get("account_sid").(string)
	countryCode := d.Get("iso_country").(string)
	pageResponse, err := client.Account(accountSid).AvailablePhoneNumber(countryCode).Local.PageWithContext(ctx, options)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No local phone numbers were found for country (%s) in account (%s)", countryCode, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list available local phone numbers: %s", err.Error())
	}

	d.SetId(accountSid + "/" + countryCode)
	d.Set("account_sid", accountSid)
	d.Set("iso_country", countryCode)

	phoneNumbers := make([]interface{}, 0)

	for _, phoneNumber := range pageResponse.AvailablePhoneNumbers {
		phoneNumbers = append(phoneNumbers, map[string]interface{}{
			"phone_number":         phoneNumber.PhoneNumber,
			"friendly_name":        phoneNumber.FriendlyName,
			"address_requirements": phoneNumber.AddressRequirements,
			"beta":                 phoneNumber.Beta,
			"capabilities": []interface{}{
				map[string]interface{}{
					"fax":   phoneNumber.Capabilities.Fax,
					"sms":   phoneNumber.Capabilities.Sms,
					"mms":   phoneNumber.Capabilities.Mms,
					"voice": phoneNumber.Capabilities.Voice,
				},
			},
			"lata":        phoneNumber.Lata,
			"rate_center": phoneNumber.RateCenter,
			"latitude":    phoneNumber.Latitude,
			"longitude":   phoneNumber.Longitude,
			"locality":    phoneNumber.Locality,
			"region":      phoneNumber.Region,
			"postal_code": phoneNumber.PostalCode,
		})
	}

	d.Set("available_phone_numbers", &phoneNumbers)

	return nil
}
