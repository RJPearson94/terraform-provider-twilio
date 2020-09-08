package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_numbers"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceProxyPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyPhoneNumbersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_numbers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"phone_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_reserved": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"capabilities": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax_inbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fax_outbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mms_inbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mms_outbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_fax_domestic": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_mms_domestic": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_sms_domestic": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"restriction_voice_domestic": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sip_trunking": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sms_inbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"sms_outbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"voice_inbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
									"voice_outbound": &schema.Schema{
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"iso_country": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"in_use": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"date_created": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"date_updated": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProxyPhoneNumbersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).PhoneNumbers.NewPhoneNumbersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No phone numbers were found for proxy service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to list proxy phone numbers resource: %s", err.Error())
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

func flattenPagePhoneNumberCapabilities(capabilities *phone_numbers.PagePhoneNumberResponseCapabilities) *[]interface{} {
	if capabilities == nil {
		return nil
	}

	results := make([]interface{}, 0)

	result := make(map[string]interface{})
	result["fax_inbound"] = capabilities.FaxInbound
	result["fax_outbound"] = capabilities.FaxOutbound
	result["mms_inbound"] = capabilities.MmsInbound
	result["mms_outbound"] = capabilities.MmsOutbound
	result["restriction_fax_domestic"] = capabilities.RestrictionFaxDomestic
	result["restriction_mms_domestic"] = capabilities.RestrictionMmsDomestic
	result["restriction_sms_domestic"] = capabilities.RestrictionSmsDomestic
	result["restriction_voice_domestic"] = capabilities.RestrictionVoiceDomestic
	result["sip_trunking"] = capabilities.SipTrunking
	result["sms_inbound"] = capabilities.SmsInbound
	result["sms_outbound"] = capabilities.SmsOutbound
	result["voice_inbound"] = capabilities.VoiceInbound
	result["voice_outbound"] = capabilities.VoiceOutbound

	results = append(results, result)
	return &results
}
