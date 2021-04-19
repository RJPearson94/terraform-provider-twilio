package phone_number

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/available_phone_number/toll_free"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourcePhoneNumberAvailableTollFreeNumbers() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "As data sources are read at the plan phase and retrive a new list of available phone numbers, this data source cannot be used to purchase a phone number. Please use the `search_criteria` block on the `twilio_phone_number` resource instead. The data source will be removed in a future release",

		ReadContext: dataSourcePhoneNumberAvailableTollFreeNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"iso_country": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"area_code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"allow_beta_numbers": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"contains_number_pattern": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude_address_requirements": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"all": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"local": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"foreign": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"location": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"in_postal_code": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_lata": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_locality": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"in_rate_center": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"near_number": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"near_lat_long": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"distance": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"capabilities": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fax_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"sms_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"mms_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"voice_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
			"available_phone_numbers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
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
						"lata": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"rate_center": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"latitude": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"longitude": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"locality": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"postal_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePhoneNumberAvailableTollFreeNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	options := &toll_free.AvailablePhoneNumbersPageOptions{
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
	pageResponse, err := client.Account(accountSid).AvailablePhoneNumber(countryCode).TollFree.PageWithContext(ctx, options)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No toll free phone numbers were found for country (%s) in account (%s)", countryCode, accountSid)
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list available toll free phone numbers: %s", err.Error())
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
