package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
)

func FlattenPhoneNumberCapabilities(capabilities *phone_number.FetchPhoneNumberCapabilitiesResponse) *[]interface{} {
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

func FlattenShortCodeCapabilities(capabilities *short_code.FetchShortCodeCapabilitiesResponse) *[]interface{} {
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
