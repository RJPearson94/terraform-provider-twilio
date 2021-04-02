package helper

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func WorkspaceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WS[0-9a-fA-F]{32}$"), "")
}

func ActivitySidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WA[0-9a-fA-F]{32}$"), "")
}

func TaskChannelSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^TC[0-9a-fA-F]{32}$"), "")
}

func TaskQueueSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WQ[0-9a-fA-F]{32}$"), "")
}

func WorkerSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WK[0-9a-fA-F]{32}$"), "")
}

func WorkflowSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WW[0-9a-fA-F]{32}$"), "")
}
