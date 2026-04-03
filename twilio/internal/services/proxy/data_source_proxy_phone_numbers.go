package proxy

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProxyPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProxyPhoneNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ProxyServiceSidValidation(),
				Description:  "The SID of the Proxy service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this Proxy service",
			},
			"phone_numbers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of phone numbers associated with the Proxy service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this Proxy phone number by Twilio",
						},
						"phone_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The phone number in E.164 format",
						},
						"is_reserved": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the phone number is reserved and not assigned to a Proxy session",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The capabilities of the Proxy phone number",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can receive faxes",
									},
									"fax_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can send faxes",
									},
									"mms_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can receive MMS messages",
									},
									"mms_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can send MMS messages",
									},
									"restriction_fax_domestic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether fax is restricted to domestic only",
									},
									"restriction_mms_domestic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether MMS is restricted to domestic only",
									},
									"restriction_sms_domestic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether SMS is restricted to domestic only",
									},
									"restriction_voice_domestic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether voice is restricted to domestic only",
									},
									"sip_trunking": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number supports SIP trunking",
									},
									"sms_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can receive SMS messages",
									},
									"sms_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can send SMS messages",
									},
									"voice_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can receive voice calls",
									},
									"voice_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the phone number can make voice calls",
									},
								},
							},
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A human-readable label for the Proxy phone number",
						},
						"iso_country": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ISO country code of the phone number",
						},
						"in_use": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of active Proxy sessions assigned to this phone number",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the Proxy phone number was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the Proxy phone number was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the Proxy phone number resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceProxyPhoneNumbersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).PhoneNumbers.NewPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No phone numbers were found for proxy service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list proxy phone numbers resource: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	phoneNumbers := make([]interface{}, 0)

	for _, phoneNumber := range paginator.PhoneNumbers {
		d.Set("account_sid", phoneNumber.AccountSid)

		phoneNumberMap := make(map[string]interface{})

		phoneNumberMap["sid"] = phoneNumber.Sid
		phoneNumberMap["phone_number"] = phoneNumber.PhoneNumber
		phoneNumberMap["friendly_name"] = phoneNumber.FriendlyName
		phoneNumberMap["iso_country"] = phoneNumber.IsoCountry
		phoneNumberMap["is_reserved"] = phoneNumber.IsReserved
		phoneNumberMap["capabilities"] = flattenPagePhoneNumberCapabilities(phoneNumber.Capabilities)
		phoneNumberMap["in_use"] = phoneNumber.InUse
		phoneNumberMap["date_created"] = phoneNumber.DateCreated.Format(time.RFC3339)

		if phoneNumber.DateUpdated != nil {
			phoneNumberMap["date_updated"] = phoneNumber.DateUpdated.Format(time.RFC3339)
		}

		phoneNumberMap["url"] = phoneNumber.URL

		phoneNumbers = append(phoneNumbers, phoneNumberMap)
	}

	d.Set("phone_numbers", &phoneNumbers)

	return nil
}

func flattenPagePhoneNumberCapabilities(capabilities *phone_numbers.PagePhoneNumberCapabilitiesResponse) *[]interface{} {
	if capabilities == nil {
		return nil
	}

	return &[]interface{}{
		map[string]interface{}{
			"fax_inbound":                capabilities.FaxInbound,
			"fax_outbound":               capabilities.FaxOutbound,
			"mms_inbound":                capabilities.MmsInbound,
			"mms_outbound":               capabilities.MmsOutbound,
			"restriction_fax_domestic":   capabilities.RestrictionFaxDomestic,
			"restriction_mms_domestic":   capabilities.RestrictionMmsDomestic,
			"restriction_sms_domestic":   capabilities.RestrictionSmsDomestic,
			"restriction_voice_domestic": capabilities.RestrictionVoiceDomestic,
			"sip_trunking":               capabilities.SipTrunking,
			"sms_inbound":                capabilities.SmsInbound,
			"sms_outbound":               capabilities.SmsOutbound,
			"voice_inbound":              capabilities.VoiceInbound,
			"voice_outbound":             capabilities.VoiceOutbound,
		},
	}
}
