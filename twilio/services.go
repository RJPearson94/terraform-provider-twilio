package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/iam"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/studio"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		iam.Registration{},
		studio.Registration{},
	}
}
