package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain"
)

func FlattenEmergency(resp *domain.FetchDomainResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"caller_sid":      resp.EmergencyCallerSid,
			"calling_enabled": resp.EmergencyCallingEnabled,
		},
	}
}

func FlattenVoice(resp *domain.FetchDomainResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"status_callback_url":    resp.VoiceStatusCallbackURL,
			"status_callback_method": resp.VoiceStatusCallbackMethod,
			"fallback_method":        resp.VoiceFallbackMethod,
			"fallback_url":           resp.VoiceFallbackURL,
			"method":                 resp.VoiceMethod,
			"url":                    resp.VoiceURL,
		},
	}
}
