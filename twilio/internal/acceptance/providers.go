package acceptance

import (
	"os"

	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type TestData struct {
	AccountSid            string
	PhoneNumberSid        string
	FlexChannelServiceSid string
}

var TestAccProviderFactories func() map[string]terraform.ResourceProviderFactory
var TestAccProvider *schema.Provider
var TestAccData *TestData

func init() {
	TestAccProvider = twilio.Provider().(*schema.Provider)
	TestAccProviderFactories = func() map[string]terraform.ResourceProviderFactory {
		factories := make(map[string]terraform.ResourceProviderFactory, 1)
		factories["twilio"] = func() (terraform.ResourceProvider, error) {
			return TestAccProvider, nil
		}

		return factories
	}
	TestAccData = &TestData{
		AccountSid:            os.Getenv("TWILIO_ACCOUNT_SID"),
		PhoneNumberSid:        os.Getenv("TWILIO_PHONE_NUMBER_SID"),
		FlexChannelServiceSid: os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
	}
}
