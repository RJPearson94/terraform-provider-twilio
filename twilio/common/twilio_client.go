package common

import "github.com/kevinburke/twilio-go"

type TwilioClient struct {
	AccountSid       string
	TerraformVersion string
	Twilio           *twilio.Client
}
