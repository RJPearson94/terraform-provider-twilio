package utils

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func AccountSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AC[0-9a-fA-F]{32}$"), "")
}

func PhoneNumberSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^PN[0-9a-fA-F]{32}$"), "")
}

func IPAccessControlListSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AL[0-9a-fA-F]{32}$"), "")
}

func CredentialListSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CL[0-9a-fA-F]{32}$"), "")
}
