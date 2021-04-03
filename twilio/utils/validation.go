package utils

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Account

func AccountSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AC[0-9a-fA-F]{32}$"), "")
}

// Address

func AddressSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AD[0-9a-fA-F]{32}$"), "")
}

// BYOC

func ByocSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^BY[0-9a-fA-F]{32}$"), "")
}

// SIP

func SIPDomainSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^SD[0-9a-fA-F]{32}$"), "")
}

func SIPIPAddressSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^IP[0-9a-fA-F]{32}$"), "")
}

func SIPCredentialSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CR[0-9a-fA-F]{32}$"), "")
}

func SIPIPAccessControlListSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AL[0-9a-fA-F]{32}$"), "")
}

func SIPCredentialListSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CL[0-9a-fA-F]{32}$"), "")
}

func SIPTrunkSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^TK[0-9a-fA-F]{32}$"), "")
}

func SIPOriginationURLValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^OU[0-9a-fA-F]{32}$"), "")
}

// Phone Number

func PhoneNumberSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^PN[0-9a-fA-F]{32}$"), "")
}

// Studio

func StudioFlowSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FW[0-9a-fA-F]{32}$"), "")
}

// TaskRouter

func TaskRouterWorkspaceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WS[0-9a-fA-F]{32}$"), "")
}

func TaskRouterActivitySidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WA[0-9a-fA-F]{32}$"), "")
}

func TaskRouterTaskChannelSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^TC[0-9a-fA-F]{32}$"), "")
}

func TaskRouterTaskQueueSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WQ[0-9a-fA-F]{32}$"), "")
}

func TaskRouterWorkerSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WK[0-9a-fA-F]{32}$"), "")
}

func TaskRouterWorkflowSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WW[0-9a-fA-F]{32}$"), "")
}

// Video

func VideoCompositionHookSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^HK[0-9a-fA-F]{32}$"), "")
}

// Voice

func VoiceQueueSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^QU[0-9a-fA-F]{32}$"), "")
}
