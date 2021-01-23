package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/twilio-sdk-go/client"
	api "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	autopilot "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	chat "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	conversations "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	flex "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	messaging "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	proxy "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	serverless "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	taskrouter "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	trunking "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Config struct {
	AccountSid       string
	AuthToken        string
	APIKey           string
	APISecret        string
	terraformVersion string
}

func (config *Config) Client() (interface{}, diag.Diagnostics) {

	creds, err := credentials.New(getCredentials(config))
	if err != nil {
		return nil, diag.FromErr(err)
	}

	sess := session.New(creds)
	sdkConfig := &client.Config{}

	client := &common.TwilioClient{
		AccountSid:       config.AccountSid,
		TerraformVersion: config.terraformVersion,

		API:           api.New(sess, sdkConfig),
		Autopilot:     autopilot.New(sess, sdkConfig),
		Chat:          chat.New(sess, sdkConfig),
		Conversations: conversations.New(sess, sdkConfig),
		Flex:          flex.New(sess, sdkConfig),
		Messaging:     messaging.New(sess, sdkConfig),
		Proxy:         proxy.New(sess, sdkConfig),
		Serverless:    serverless.New(sess, sdkConfig),
		SIPTrunking:   trunking.New(sess, sdkConfig),
		Studio:        studio.New(sess, sdkConfig),
		TaskRouter:    taskrouter.New(sess, sdkConfig),
	}
	return client, nil
}

func getCredentials(config *Config) credentials.TwilioCredentials {
	if config.APIKey != "" && config.APISecret != "" {
		return credentials.APIKey{
			Account: config.AccountSid,
			Sid:     config.APIKey,
			Value:   config.APISecret,
		}
	}
	return credentials.Account{
		Sid:       config.AccountSid,
		AuthToken: config.AuthToken,
	}
}
