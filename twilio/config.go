package twilio

import twilio "github.com/kevinburke/twilio-go"

type Config struct {
	AccountSid       string
	AuthToken        string
	terraformVersion string
}

type TwilioClient struct {
	accountSid       string
	terraformVersion string
	twilio           *twilio.Client
}

func (config *Config) Client() (interface{}, error) {
	twilioClient := twilio.NewClient(config.AccountSid, config.AuthToken, nil)

	client := &TwilioClient{
		accountSid:       config.AccountSid,
		terraformVersion: config.terraformVersion,
		twilio:           twilioClient,
	}
	return client, nil
}
