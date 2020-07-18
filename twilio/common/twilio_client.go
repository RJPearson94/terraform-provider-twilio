package common

import (
	api "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	autopilot "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	chat "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	proxy "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	serverless "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	taskrouter "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
)

type TwilioClient struct {
	AccountSid       string
	TerraformVersion string
	API              *api.V2010
	Autopilot        *autopilot.Autopilot
	Chat             *chat.Chat
	Proxy            *proxy.Proxy
	Serverless       *serverless.Serverless
	Studio           *studio.Studio
	TaskRouter       *taskrouter.TaskRouter
}
