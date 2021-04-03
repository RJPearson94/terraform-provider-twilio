package helper

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DomainSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^SD[0-9a-fA-F]{32}$"), "")
}

func ByocSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^BY[0-9a-fA-F]{32}$"), "")
}

func IPAddressSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^IP[0-9a-fA-F]{32}$"), "")
}

func CredentialSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CR[0-9a-fA-F]{32}$"), "")
}
