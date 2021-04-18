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

// Application

func ApplicationSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AP[0-9a-fA-F]{32}$"), "")
}

// Autopilot

func AutopilotAssistantSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UA[0-9a-fA-F]{32}$"), "")
}

func AutopilotFieldTypeSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UB[0-9a-fA-F]{32}$"), "")
}

func AutopilotFieldValueSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UC[0-9a-fA-F]{32}$"), "")
}

func AutopilotModelBuildSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UG[0-9a-fA-F]{32}$"), "")
}

func AutopilotTaskFieldSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UE[0-9a-fA-F]{32}$"), "")
}

func AutopilotTaskSampleSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UF[0-9a-fA-F]{32}$"), "")
}

func AutopilotTaskSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UD[0-9a-fA-F]{32}$"), "")
}

func AutopilotWebhookSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^UM[0-9a-fA-F]{32}$"), "")
}

// Bundle

func BundleSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^BU[0-9a-fA-F]{32}$"), "")
}

// BYOC

func ByocSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^BY[0-9a-fA-F]{32}$"), "")
}

// Chat (near duplicate of conversations so all references can be removed when chat is removed from the provider)

func ChatServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^IS[0-9a-fA-F]{32}$"), "")
}

func ChatChannelMemberSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^MB[0-9a-fA-F]{32}$"), "")
}

func ChatChannelSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^CH[0-9a-fA-F]{32}$"), "")
}

func ChatChannelWebhookSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^WH[0-9a-fA-F]{32}$"), "")
}

func ChatRoleSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^RL[0-9a-fA-F]{32}$"), "")
}

func ChatUserSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^US[0-9a-fA-F]{32}$"), "")
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

func FlexFlowSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FO[0-9a-fA-F]{32}$"), "")
}

func FlexPluginSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FP[0-9a-fA-F]{32}$"), "")
}

func FlexPluginReleaseSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FK[0-9a-fA-F]{32}$"), "")
}

func FlexPluginConfigurationSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FJ[0-9a-fA-F]{32}$"), "")
}

func FlexPluginVersionSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^FV[0-9a-fA-F]{32}$"), "")
}

// Identity

func IdentitySidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^RI[0-9a-fA-F]{32}$"), "")
}

// Messaging

func MessagingAlphaSenderSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^AI[0-9a-fA-F]{32}$"), "")
}

func MessagingServiceSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^MG[0-9a-fA-F]{32}$"), "")
}

// Phone Number

func PhoneNumberSidValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile("^PN[0-9a-fA-F]{32}$"), "")
}

func PhoneNumberValidation() schema.SchemaValidateFunc {
	return validation.StringMatch(regexp.MustCompile(`^\+[1-9]\d{1,14}$`), "")
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
