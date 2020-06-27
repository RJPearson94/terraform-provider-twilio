package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/chat"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/iam"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/proxy"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/serverless"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/taskrouter"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		chat.Registration{},
		iam.Registration{},
		proxy.Registration{},
		serverless.Registration{},
		studio.Registration{},
		taskrouter.Registration{},
	}
}
