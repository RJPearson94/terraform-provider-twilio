package sip_trunking

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSIPTrunkingPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSIPTrunkingPhoneNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"trunk_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: helper.TrunkSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
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
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
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
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"status_callback_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status_callback_method": {
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
				},
			},
		},
	}
}

func dataSourceSIPTrunkingPhoneNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	trunkSid := d.Get("trunk_sid").(string)
	paginator := client.Trunk(trunkSid).PhoneNumbers.NewPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No phone numbers were found for SIP trunk with sid (%s)", trunkSid)
		}
		return diag.Errorf("Failed to list SIP trunk phone numbers: %s", err.Error())
	}

	d.SetId(trunkSid)
	d.Set("trunk_sid", trunkSid)

	phoneNumbers := make([]interface{}, 0)

	for _, phoneNumber := range paginator.PhoneNumbers {
		d.Set("account_sid", phoneNumber.AccountSid)

		phoneNumbersMap := make(map[string]interface{})

		phoneNumbersMap["sid"] = phoneNumber.Sid
		phoneNumbersMap["address_requirements"] = phoneNumber.AddressRequirements
		phoneNumbersMap["beta"] = phoneNumber.Beta
		phoneNumbersMap["capabilities"] = &[]interface{}{
			map[string]interface{}{
				"fax":   phoneNumber.Capabilities.Fax,
				"sms":   phoneNumber.Capabilities.Sms,
				"mms":   phoneNumber.Capabilities.Mms,
				"voice": phoneNumber.Capabilities.Voice,
			},
		}
		phoneNumbersMap["friendly_name"] = phoneNumber.FriendlyName
		phoneNumbersMap["messaging"] = &[]interface{}{
			map[string]interface{}{
				"application_sid": phoneNumber.SmsApplicationSid,
				"fallback_method": phoneNumber.SmsFallbackMethod,
				"fallback_url":    phoneNumber.SmsFallbackURL,
				"method":          phoneNumber.SmsMethod,
				"url":             phoneNumber.SmsURL,
			},
		}
		phoneNumbersMap["phone_number"] = phoneNumber.PhoneNumber
		phoneNumbersMap["status_callback_url"] = phoneNumber.StatusCallback
		phoneNumbersMap["status_callback_method"] = phoneNumber.StatusCallbackMethod

		if helper.IsVoiceReceiveMode(phoneNumber.VoiceReceiveMode) {
			phoneNumbersMap["voice"] = &[]interface{}{
				map[string]interface{}{
					"application_sid":  phoneNumber.VoiceApplicationSid,
					"caller_id_lookup": phoneNumber.VoiceCallerIDLookup,
					"fallback_method":  phoneNumber.VoiceFallbackMethod,
					"fallback_url":     phoneNumber.VoiceFallbackURL,
					"method":           phoneNumber.VoiceMethod,
					"url":              phoneNumber.VoiceURL,
				},
			}
			phoneNumbersMap["fax"] = &[]interface{}{}
		} else {
			phoneNumbersMap["fax"] = &[]interface{}{
				map[string]interface{}{
					"application_sid": phoneNumber.VoiceApplicationSid,
					"fallback_method": phoneNumber.VoiceFallbackMethod,
					"fallback_url":    phoneNumber.VoiceFallbackURL,
					"method":          phoneNumber.VoiceMethod,
					"url":             phoneNumber.VoiceURL,
				},
			}
			phoneNumbersMap["voice"] = &[]interface{}{}
		}

		phoneNumbersMap["date_created"] = phoneNumber.DateCreated.Format(time.RFC3339)

		if phoneNumber.DateUpdated != nil {
			phoneNumbersMap["date_updated"] = phoneNumber.DateUpdated.Format(time.RFC3339)
		}
		phoneNumbersMap["url"] = phoneNumber.URL

		phoneNumbers = append(phoneNumbers, phoneNumbersMap)
	}

	d.Set("phone_numbers", &phoneNumbers)

	return nil
}
