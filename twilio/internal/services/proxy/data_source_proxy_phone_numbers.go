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
				Type:     schema.TypeString,
				Required: true,
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
						"phone_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_reserved": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"capabilities": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax_inbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fax_outbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mms_inbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mms_outbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_fax_domestic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_mms_domestic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_sms_domestic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_voice_domestic": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sip_trunking": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sms_inbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sms_outbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"voice_inbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"voice_outbound": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iso_country": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"in_use": {
							Type:     schema.TypeInt,
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
