package utils

import "github.com/RJPearson94/twilio-sdk-go/utils"

func IsNotFoundError(err error) bool {
	if _, ok := err.(*utils.TwilioError); ok {
		return err.(*utils.TwilioError).IsNotFoundError()
	}
	return false
}
