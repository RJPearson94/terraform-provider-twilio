package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/twilio-sdk-go/client"
	accounts "github.com/RJPearson94/twilio-sdk-go/service/accounts/v1"
	api "github.com/RJPearson94/twilio-sdk-go/service/api/v2010"
	autopilot "github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1"
	chat "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	conversations "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	flex "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	messaging "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	proxy "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	serverless "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	sync "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	taskrouter "github.com/RJPearson94/twilio-sdk-go/service/taskrouter/v1"
	trunking "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	verify "github.com/RJPearson94/twilio-sdk-go/service/verify/v2"
	video "github.com/RJPearson94/twilio-sdk-go/service/video/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	"github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Config struct {
	AccountSid               string
	AuthToken                string
	APIKey                   string
	APISecret                string
	SkipCredentialValidation bool
	RetryAttempts            int
	BackoffInterval          int
	Edge                     string
	Region                   string
	terraformVersion         string
}

func (config *Config) Client() (interface{}, diag.Diagnostics) {

	creds, err := sessionCredentials(config)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	sess := session.New(creds)
	sdkConfig := &client.Config{
		RetryAttempts:   utils.Int(config.RetryAttempts),
		BackoffInterval: utils.Int(config.BackoffInterval),
	}

	if config.Edge != "" {
		sdkConfig.Edge = utils.String(config.Edge)
	}
	if config.Region != "" {
		sdkConfig.Region = utils.String(config.Region)
	}

	client := &common.TwilioClient{
		AccountSid:       config.AccountSid,
		TerraformVersion: config.terraformVersion,

		Accounts:      accounts.New(sess, sdkConfig),
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
		Sync:          sync.New(sess, sdkConfig),
		TaskRouter:    taskrouter.New(sess, sdkConfig),
		Verify:        verify.New(sess, sdkConfig),
		Video:         video.New(sess, sdkConfig),
	}
	return client, nil
}

func sessionCredentials(config *Config) (*credentials.Credentials, error) {
	creds := getCredentials(config)
	if config.SkipCredentialValidation == true {
		return credentials.NewWithNoValidation(creds), nil
	}
	return credentials.New(creds)
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
