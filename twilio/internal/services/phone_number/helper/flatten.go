package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/incoming_phone_number"
)

func FlattenCapabilities(resp *incoming_phone_number.FetchIncomingPhoneNumberResponseCapabilities) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"fax":   resp.Fax,
			"sms":   resp.Sms,
			"mms":   resp.Mms,
			"voice": resp.Voice,
		},
	}
}

func FlattenMessaging(resp *incoming_phone_number.FetchIncomingPhoneNumberResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"application_sid": resp.SmsApplicationSid,
			"fallback_method": resp.SmsFallbackMethod,
			"fallback_url":    resp.SmsFallbackURL,
			"method":          resp.SmsMethod,
			"url":             resp.SmsURL,
		},
	}
}

func FlattenVoice(resp *incoming_phone_number.FetchIncomingPhoneNumberResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"application_sid":  resp.VoiceApplicationSid,
			"caller_id_lookup": resp.VoiceCallerIDLookup,
			"fallback_method":  resp.VoiceFallbackMethod,
			"fallback_url":     resp.VoiceFallbackURL,
			"method":           resp.VoiceMethod,
			"receive_mode":     resp.VoiceReceiveMode,
			"url":              resp.VoiceURL,
		},
	}
}
