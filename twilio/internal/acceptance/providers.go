package acceptance

import (
	"os"

	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

type AddressDetails struct {
	Street          string
	StreetSecondary string
	City            string
	Region          string
	PostalCode      string
	IsoCountry      string
}

type TestData struct {
	AccountSid            string
	PhoneNumberSid        string
	FlexChannelServiceSid string
	CustomerName          string
	Address               AddressDetails
}

var TestAccProviders map[string]terraform.ResourceProvider
var TestAccProviderFactories func() map[string]terraform.ResourceProviderFactory
var TestAccProvider *schema.Provider
var TestAccData *TestData

func init() {
	TestAccProvider = twilio.Provider().(*schema.Provider)
	TestAccProviders = map[string]terraform.ResourceProvider{
		"twilio": TestAccProvider,
	}
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
		CustomerName:          os.Getenv("TWILIO_CUSTOMER_NAME"),
		Address: AddressDetails{
			Street:          os.Getenv("TWILIO_ADDRESS_STREET"),
			StreetSecondary: os.Getenv("TWILIO_ADDRESS_STREET_SECONDARY"),
			City:            os.Getenv("TWILIO_ADDRESS_CITY"),
			Region:          os.Getenv("TWILIO_ADDRESS_REGION"),
			PostalCode:      os.Getenv("TWILIO_ADDRESS_POSTAL_CODE"),
			IsoCountry:      os.Getenv("TWILIO_ADDRESS_ISO_COUNTRY"),
		},
	}
}
