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
				ValidateFunc: utils.SIPTrunkSidValidation(),
				Description:  "The SID of the SIP trunk to retrieve phone numbers for",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns the SIP trunk phone numbers",
			},
			"phone_numbers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of phone numbers associated with the SIP trunk",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this SIP trunk phone number by Twilio",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the SIP trunk phone number",
						},
						"phone_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The phone number in E.164 format",
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
						"messaging": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The messaging settings for the phone number",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SID of the TwiML application that handles incoming messages",
									},
									"fallback_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the messaging fallback URL",
									},
									"fallback_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming messages",
									},
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when the phone number receives an incoming message",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the messaging URL",
									},
								},
							},
						},
						"voice": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The voice settings for the phone number",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SID of the TwiML application that handles incoming voice calls",
									},
									"caller_id_lookup": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether caller ID lookup is enabled for incoming voice calls",
									},
									"fallback_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the voice fallback URL",
									},
									"fallback_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming voice calls",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the voice URL",
									},
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when the phone number receives an incoming voice call",
									},
								},
							},
						},
						"fax": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The fax settings for the phone number",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"application_sid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The SID of the TwiML application that handles incoming faxes",
									},
									"fallback_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the fax fallback URL",
									},
									"fallback_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming faxes",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The HTTP method used to call the fax URL",
									},
									"url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL to call when the phone number receives an incoming fax",
									},
								},
							},
						},
						"status_callback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL called for status callback events on the phone number",
						},
						"status_callback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the status callback URL",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the SIP trunk phone number was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the SIP trunk phone number was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the SIP trunk phone number resource",
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
