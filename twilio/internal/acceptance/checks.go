package acceptance

import (
	"os"
	"testing"
)

func PreCheck(t *testing.T) {
	InitialiseProviders()

	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_PHONE_NUMBER_SID",
		"TWILIO_FLEX_CHANNEL_SERVICE_SID",
		"TWILIO_PUBLIC_KEY",
		"TWILIO_CUSTOMER_NAME",
		"TWILIO_ADDRESS_STREET",
		"TWILIO_ADDRESS_STREET_SECONDARY",
		"TWILIO_ADDRESS_CITY",
		"TWILIO_ADDRESS_REGION",
		"TWILIO_ADDRESS_POSTAL_CODE",
		"TWILIO_ADDRESS_ISO_COUNTRY",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			t.Fatalf("`%s` are required for running acceptance tests", variable)
		}
	}
}
