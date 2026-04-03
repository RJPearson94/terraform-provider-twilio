package proxy

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_codes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProxyShortCodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProxyShortCodesRead,

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
			"short_codes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of short codes associated with the Proxy service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this Proxy short code by Twilio",
						},
						"is_reserved": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the short code is reserved and not assigned to a Proxy session",
						},
						"capabilities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The capabilities of the Proxy short code",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fax_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can receive faxes",
									},
									"fax_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can send faxes",
									},
									"mms_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can receive MMS messages",
									},
									"mms_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can send MMS messages",
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
										Description: "Whether the short code supports SIP trunking",
									},
									"sms_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can receive SMS messages",
									},
									"sms_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can send SMS messages",
									},
									"voice_inbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can receive voice calls",
									},
									"voice_outbound": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the short code can make voice calls",
									},
								},
							},
						},
						"short_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The short code value",
						},
						"iso_country": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ISO country code of the short code",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the Proxy short code was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the Proxy short code was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the Proxy short code resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceProxyShortCodesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Proxy

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).ShortCodes.NewShortCodesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No short codes were found for proxy service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to list proxy short codes resource: %s", err.Error())
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

func flattenPageShortCodeCapabilities(capabilities *short_codes.PageShortCodeCapabilitiesResponse) *[]interface{} {
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
