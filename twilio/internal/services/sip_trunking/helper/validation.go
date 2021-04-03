package helper

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func TrunkSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^TK[0-9a-fA-F]{32}$"), "")
}

func OriginationURLValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^OU[0-9a-fA-F]{32}$"), "")
}
