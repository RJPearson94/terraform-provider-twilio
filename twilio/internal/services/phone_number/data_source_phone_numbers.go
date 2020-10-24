package phone_number

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePhoneNumbers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourcePhoneNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"phone_numbers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"address_sid": {
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
						"emergency_address_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"emergency_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"messaging": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"trunk_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"voice": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"caller_id_lookup": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fallback_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"fax": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"identity_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bundle_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_callback_method": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourcePhoneNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	accountSid := d.Get("account_sid").(string)
	paginator := client.Account(accountSid).IncomingPhoneNumbers.NewIncomingPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list phone numbers: %s", err.Error())
	}

	d.SetId(accountSid)
	d.Set("account_sid", accountSid)

	phoneNumbers := make([]interface{}, 0)

	for _, phoneNumber := range paginator.PhoneNumbers {
		phoneNumberMap := make(map[string]interface{})

		phoneNumberMap["sid"] = phoneNumber.Sid
		phoneNumberMap["address_sid"] = phoneNumber.AddressSid
		phoneNumberMap["address_requirements"] = phoneNumber.AddressRequirements
		phoneNumberMap["beta"] = phoneNumber.Beta
		phoneNumberMap["bundle_sid"] = phoneNumber.BundleSid
		phoneNumberMap["capabilities"] = []interface{}{
			map[string]interface{}{
				"fax":   phoneNumber.Capabilities.Fax,
				"sms":   phoneNumber.Capabilities.Sms,
				"mms":   phoneNumber.Capabilities.Mms,
				"voice": phoneNumber.Capabilities.Voice,
			},
		}
		phoneNumberMap["emergency_address_sid"] = phoneNumber.EmergencyAddressSid
		phoneNumberMap["emergency_status"] = phoneNumber.EmergencyStatus
		phoneNumberMap["friendly_name"] = phoneNumber.FriendlyName
		phoneNumberMap["identity_sid"] = phoneNumber.IdentitySid
		phoneNumberMap["messaging"] = []interface{}{
			map[string]interface{}{
				"application_sid": phoneNumber.SmsApplicationSid,
				"fallback_method": phoneNumber.SmsFallbackMethod,
				"fallback_url":    phoneNumber.SmsFallbackURL,
				"method":          phoneNumber.SmsMethod,
				"url":             phoneNumber.SmsURL,
			},
		}
		phoneNumberMap["origin"] = phoneNumber.Origin
		phoneNumberMap["phone_number"] = phoneNumber.PhoneNumber
		phoneNumberMap["status"] = phoneNumber.Status
		phoneNumberMap["status_callback_url"] = phoneNumber.StatusCallback
		phoneNumberMap["status_callback_method"] = phoneNumber.StatusCallbackMethod
		phoneNumberMap["trunk_sid"] = phoneNumber.TrunkSid

		if phoneNumber.VoiceReceiveMode == "voice" {
			phoneNumberMap["voice"] = []interface{}{
				map[string]interface{}{
					"application_sid":  phoneNumber.VoiceApplicationSid,
					"caller_id_lookup": phoneNumber.VoiceCallerIDLookup,
					"fallback_method":  phoneNumber.VoiceFallbackMethod,
					"fallback_url":     phoneNumber.VoiceFallbackURL,
					"method":           phoneNumber.VoiceMethod,
					"url":              phoneNumber.VoiceURL,
				},
			}
		}

		if phoneNumber.VoiceReceiveMode == "fax" {
			phoneNumberMap["fax"] = []interface{}{
				map[string]interface{}{
					"application_sid": phoneNumber.VoiceApplicationSid,
					"fallback_method": phoneNumber.VoiceFallbackMethod,
					"fallback_url":    phoneNumber.VoiceFallbackURL,
					"method":          phoneNumber.VoiceMethod,
					"url":             phoneNumber.VoiceURL,
				},
			}
		}
		phoneNumberMap["date_created"] = phoneNumber.DateCreated.Time.Format(time.RFC3339)

		if phoneNumber.DateUpdated != nil {
			phoneNumberMap["date_updated"] = phoneNumber.DateUpdated.Format(time.RFC3339)
		}

		phoneNumbers = append(phoneNumbers, phoneNumberMap)
	}

	d.Set("phone_numbers", &phoneNumbers)

	return nil
}
