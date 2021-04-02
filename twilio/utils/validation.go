package utils

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func AccountSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AC[0-9a-fA-F]{32}$"), "")
}
