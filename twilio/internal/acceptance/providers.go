package acceptance

import (
	"os"

	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type TestData struct {
	PhoneNumberSid string
}

var TestAccProviders map[string]terraform.ResourceProvider
var TestAccProvider *schema.Provider
var TestAccData *TestData

func init() {
	TestAccProvider = twilio.Provider().(*schema.Provider)
	TestAccProviders = map[string]terraform.ResourceProvider{
		"twilio": TestAccProvider,
	}
	TestAccData = &TestData{
		PhoneNumberSid: os.Getenv("TWILIO_PHONE_NUMBER_SID"),
	}
}
