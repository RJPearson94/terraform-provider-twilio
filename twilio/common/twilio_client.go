package common

import (
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	taskrouter "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/kevinburke/twilio-go"
)

type TwilioClient struct {
	AccountSid       string
	TerraformVersion string
	Twilio           *twilio.Client
	Studio           *studio.Studio
	TaskRouter       *taskrouter.TaskRouter
}
