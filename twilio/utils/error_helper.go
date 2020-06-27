package utils

import "github.com/RJPearson94/twilio-sdk-go/utils"

func IsNotFoundError(err error) bool {
	if twilioError, ok := err.(*utils.TwilioError); ok {
		return twilioError.IsNotFoundError()
	}
	return false
}
