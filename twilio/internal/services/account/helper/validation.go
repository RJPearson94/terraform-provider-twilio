package helper

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func AddressSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AD[0-9a-fA-F]{32}$"), "")
}
