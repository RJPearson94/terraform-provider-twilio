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

// Chat

func ChatInstanceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^IS[0-9a-fA-F]{32}$"), "")
}

// Conversations

func ConversationServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^IS[0-9a-fA-F]{32}$"), "")
}

func ConversationRoleSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^RL[0-9a-fA-F]{32}$"), "")
}

func ConversationSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CH[0-9a-fA-F]{32}$"), "")
}

func ConversationWebhookSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WH[0-9a-fA-F]{32}$"), "")
}

func ConversationUserSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^US[0-9a-fA-F]{32}$"), "")
}

// Flex

func FlowSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FW[0-9a-fA-F]{32}$"), "")
}

// Messaging

func MessagingServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^MG[0-9a-fA-F]{32}$"), "")
}

// Phone Number

func PhoneNumberSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^PN[0-9a-fA-F]{32}$"), "")
}

// Proxy

func ProxyServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^KS[0-9a-fA-F]{32}$"), "")
}

// Serverless

func ServerlessServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZS[0-9a-fA-F]{32}$"), "")
}

func ServerlessAssetSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZH[0-9a-fA-F]{32}$"), "")
}

func ServerlessAssetVersionSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZN[0-9a-fA-F]{32}$"), "")
}

func ServerlessFunctionSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZH[0-9a-fA-F]{32}$"), "")
}

func ServerlessFunctionVersionSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZN[0-9a-fA-F]{32}$"), "")
}

func ServerlessEnvironmentSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZE[0-9a-fA-F]{32}$"), "")
}

func ServerlessBuildSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZB[0-9a-fA-F]{32}$"), "")
}

func ServerlessDeploymentSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZD[0-9a-fA-F]{32}$"), "")
}

func ServerlessVariableSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^ZV[0-9a-fA-F]{32}$"), "")
}

// Short Code

func ShortCodeSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^SC[0-9a-fA-F]{32}$"), "")
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
