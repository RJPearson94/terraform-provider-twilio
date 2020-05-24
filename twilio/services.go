package twilio

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/iam"
)

func SupportedServices() []common.ServiceRegistration {
	return []common.ServiceRegistration{
		iam.Registration{},
	}
}
