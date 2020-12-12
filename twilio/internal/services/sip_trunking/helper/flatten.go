package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/phone_number"
)

func FlattenRecording(recording trunk.FetchTrunkRecordingResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"mode": recording.Mode,
			"trim": recording.Trim,
		},
	}
}

func FlattenCapabilities(resp *phone_number.FetchPhoneNumberCapabilitiesResponse) *[]interface{} {
	capabilities := make(map[string]interface{})
	if resp.Fax != nil {
		capabilities["fax"] = resp.Fax
	}
	capabilities["sms"] = resp.Sms
	capabilities["mms"] = resp.Mms
	capabilities["voice"] = resp.Voice

	return &[]interface{}{capabilities}
}

func FlattenMessaging(resp *phone_number.FetchPhoneNumberResponse) *[]interface{} {
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

func FlattenVoice(resp *phone_number.FetchPhoneNumberResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"application_sid":  resp.VoiceApplicationSid,
			"caller_id_lookup": resp.VoiceCallerIDLookup,
			"fallback_method":  resp.VoiceFallbackMethod,
			"fallback_url":     resp.VoiceFallbackURL,
			"method":           resp.VoiceMethod,
			"url":              resp.VoiceURL,
		},
	}
}

func FlattenFax(resp *phone_number.FetchPhoneNumberResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"application_sid": resp.VoiceApplicationSid,
			"fallback_method": resp.VoiceFallbackMethod,
			"fallback_url":    resp.VoiceFallbackURL,
			"method":          resp.VoiceMethod,
			"url":             resp.VoiceURL,
		},
	}
}
