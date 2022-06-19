package helper

import (
	"github.com/RJPearson94/twilio-sdk-go/service/verify/v2/service"
)

func FlattenPush(input service.FetchServicePushResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"apn_credential_sid": input.ApnCredentialSid,
			"fcm_credential_sid": input.FcmCredentialSid,
		},
	}
}

func FlattenTotp(input service.FetchServiceTotpResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"issuer":      input.Issuer,
			"time_step":   input.TimeStep,
			"code_length": input.CodeLength,
			"skew":        input.Skew,
		},
	}
}
