package acceptance

import (
	"os"
	"testing"
)

func PreCheck(t *testing.T) {
	variables := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` are required for running acceptance tests", variable)
		}
	}
}
