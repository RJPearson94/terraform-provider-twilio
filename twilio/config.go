package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	studio "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
	twilio "github.com/kevinburke/twilio-go"
)

type Config struct {
	AccountSid       string
	AuthToken        string
	terraformVersion string
}

func (config *Config) Client() (interface{}, error) {
	creds, err := credentials.New(credentials.Account{
		Sid:       config.AccountSid,
		AuthToken: config.AuthToken,
	})
	if err != nil {
		return nil, err
	}

	client := &common.TwilioClient{
		AccountSid:       config.AccountSid,
		TerraformVersion: config.terraformVersion,
		Twilio:           twilio.NewClient(config.AccountSid, config.AuthToken, nil),
		Studio:           studio.NewWithCredentials(creds),
	}
	return client, nil
}
