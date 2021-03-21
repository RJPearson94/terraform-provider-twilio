package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/account"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/autopilot"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/chat"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/conversations"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/flex"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/iam"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/messaging"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/phone_number"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/serverless"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/taskrouter"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/video"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/voice"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		account.Registration{},
		autopilot.Registration{},
		chat.Registration{},
		conversations.Registration{},
		flex.Registration{},
		iam.Registration{},
		messaging.Registration{},
		phone_number.Registration{},
		proxy.Registration{},
		serverless.Registration{},
		studio.Registration{},
		sip.Registration{},
		sip_trunking.Registration{},
		taskrouter.Registration{},
		video.Registration{},
		voice.Registration{},
	}
}
