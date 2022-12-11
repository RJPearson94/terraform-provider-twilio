package helper

import sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"

func StringToBool(val *string) *bool {
	if val == nil {
		return nil
	}
	if *val == "True" {
		return sdkUtils.Bool(true)
	}
	return sdkUtils.Bool(false)
}
