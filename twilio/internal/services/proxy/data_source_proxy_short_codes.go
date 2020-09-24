package proxy

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceProxyShortCodes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProxyShortCodesRead,

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
			"short_codes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
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
						"short_code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"iso_country": {
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

func dataSourceProxyShortCodesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Proxy
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).ShortCodes.NewShortCodesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No short codes were found for proxy service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to list proxy short codes resource: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	shortCodes := make([]interface{}, 0)

	for _, shortCode := range paginator.ShortCodes {
		d.Set("account_sid", shortCode.AccountSid)

		shortCodeMap := make(map[string]interface{})

		shortCodeMap["sid"] = shortCode.Sid
		shortCodeMap["short_code"] = shortCode.ShortCode
		shortCodeMap["iso_country"] = shortCode.IsoCountry
		shortCodeMap["is_reserved"] = shortCode.IsReserved
		shortCodeMap["capabilities"] = flattenPageShortCodeCapabilities(shortCode.Capabilities)
		shortCodeMap["date_created"] = shortCode.DateCreated.Format(time.RFC3339)

		if shortCode.DateUpdated != nil {
			shortCodeMap["date_updated"] = shortCode.DateUpdated.Format(time.RFC3339)
		}

		shortCodeMap["url"] = shortCode.URL

		shortCodes = append(shortCodes, shortCodeMap)
	}

	d.Set("short_codes", &shortCodes)

	return nil
}

func flattenPageShortCodeCapabilities(capabilities *short_codes.PageShortCodeResponseCapabilities) *[]interface{} {
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
