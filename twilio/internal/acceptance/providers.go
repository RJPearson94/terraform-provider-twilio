package acceptance

import (
	"os"
	"sync"

	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

var TestAccProvider *schema.Provider
var TestAccProviderFactories map[string]func() (*schema.Provider, error)
var TestAccData *TestData
var once sync.Once

func init() {
	if os.Getenv("TF_ACC") == "" {
		return
	}

	InitialiseProviders()
}

func InitialiseProviders() {
	once.Do(func() {
		TestAccProvider = twilio.Provider()
		TestAccProviderFactories = map[string]func() (*schema.Provider, error){
			"twilio": func() (*schema.Provider, error) {
				return TestAccProvider, nil
			},
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
	})
}
