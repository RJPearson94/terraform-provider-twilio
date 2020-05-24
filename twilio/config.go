package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	twilio "github.com/kevinburke/twilio-go"
)

type Config struct {
	AccountSid       string
	AuthToken        string
	terraformVersion string
}

func (config *Config) Client() (interface{}, error) {
	twilioClient := twilio.NewClient(config.AccountSid, config.AuthToken, nil)

	client := &common.TwilioClient{
		AccountSid:       config.AccountSid,
		TerraformVersion: config.terraformVersion,
		Twilio:           twilioClient,
	}
	return client, nil
}
