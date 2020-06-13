package common

import (
	serverless "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	taskrouter "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	"github.com/kevinburke/twilio-go"
)

type TwilioClient struct {
	AccountSid       string
	TerraformVersion string
	Twilio           *twilio.Client
	Serverless       *serverless.Serverless
	Studio           *studio.Studio
	TaskRouter       *taskrouter.TaskRouter
}
