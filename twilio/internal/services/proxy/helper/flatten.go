package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service/short_code"
)

func FlattenPhoneNumberCapabilities(capabilities *phone_number.FetchPhoneNumberResponseCapabilities) *[]interface{} {
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

func FlattenShortCodeCapabilities(capabilities *short_code.FetchShortCodeResponseCapabilities) *[]interface{} {
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
