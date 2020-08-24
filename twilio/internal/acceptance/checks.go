package acceptance

import (
	"os"
	"testing"
)

func PreCheck(t *testing.T) {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_PHONE_NUMBER_SID",
		"TWILIO_FLEX_CHANNEL_SERVICE_SID",
	}

	for _, variable := range variables {
		if value := os.Getenv(variable); value == "" {
			t.Fatalf("`%s` are required for running acceptance tests", variable)
		}
	}
}
