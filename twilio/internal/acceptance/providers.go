package acceptance

import (
	"github.com/RJPearson94/terraform-provider-twilio/twilio"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TestAccProviders map[string]terraform.ResourceProvider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = twilio.Provider().(*schema.Provider)
	TestAccProviders = map[string]terraform.ResourceProvider{
		"twilio": TestAccProvider,
	}
}
