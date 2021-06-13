package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/application"
)

func FlattenMessaging(resp *application.FetchApplicationResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"status_callback_url": resp.MessageStatusCallback,
			"fallback_method":     resp.SmsFallbackMethod,
			"fallback_url":        resp.SmsFallbackURL,
			"method":              resp.SmsMethod,
			"url":                 resp.SmsURL,
		},
	}
}

func FlattenVoice(resp *application.FetchApplicationResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"caller_id_lookup":       resp.VoiceCallerIDLookup,
			"fallback_method":        resp.VoiceFallbackMethod,
			"fallback_url":           resp.VoiceFallbackURL,
			"method":                 resp.VoiceMethod,
			"url":                    resp.VoiceURL,
			"status_callback_method": resp.StatusCallbackMethod,
			"status_callback_url":    resp.StatusCallback,
		},
	}
}
