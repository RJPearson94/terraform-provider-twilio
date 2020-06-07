package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/iam"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/taskrouter"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		iam.Registration{},
		studio.Registration{},
		taskrouter.Registration{},
	}
}
