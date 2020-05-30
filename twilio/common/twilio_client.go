package common

import (
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/kevinburke/twilio-go"
)

type TwilioClient struct {
	AccountSid       string
	TerraformVersion string
	Twilio           *twilio.Client
	Studio           *studio.Studio
}
