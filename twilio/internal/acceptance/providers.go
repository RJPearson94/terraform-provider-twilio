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
	AccountSid             string
	PurchasablePhoneNumber string // TODO: Temp hack this needs to be looked up
	PhoneNumberSid         string
	PhoneNumber            string
	FlexChannelServiceSid  string
	PublicKey              string
	AWSAccessKeyID         string
	AWSSecretAccessKey     string
	CustomerName           string
	Address                AddressDetails
}

var TestAccProvider *schema.Provider
var TestAccProviderFactories map[string]func() (*schema.Provider, error)
var TestAccData *TestData
var once sync.Once

func init() {
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
			AccountSid:             os.Getenv("TWILIO_ACCOUNT_SID"),
			PurchasablePhoneNumber: os.Getenv("TWILIO_PURCHASABLE_PHONE_NUMBER"), // TODO: Temp hack this needs to be looked up
			PhoneNumberSid:         os.Getenv("TWILIO_PHONE_NUMBER_SID"),
			PhoneNumber:            os.Getenv("TWILIO_PHONE_NUMBER"),
			FlexChannelServiceSid:  os.Getenv("TWILIO_FLEX_CHANNEL_SERVICE_SID"),
			PublicKey:              os.Getenv("TWILIO_PUBLIC_KEY"),
			AWSAccessKeyID:         os.Getenv("TWILIO_AWS_ACCESS_KEY_ID"),     // I am using TWILIO_AWS_ACCESS_KEY_ID to ensure only the correct AWS credentials are sent to Twilio
			AWSSecretAccessKey:     os.Getenv("TWILIO_AWS_SECRET_ACCESS_KEY"), // I am using TWILIO_AWS_SECRET_ACCESS_KEY to ensure only the correct AWS credentials are sent to Twilio
			CustomerName:           os.Getenv("TWILIO_CUSTOMER_NAME"),
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
