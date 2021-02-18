package tests

import (
	"fmt"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var domainDataSourceName = "twilio_sip_domain"

func TestAccDataSourceTwilioSIPDomain_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("%s.domain", domainDataSourceName)

	testData := acceptance.TestAccData
	domainName := acctest.RandString(10) + ".sip.twilio.com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioSIPDomain_basic(testData, domainName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttr(stateDataSourceName, "domain_name", domainName),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "auth_type", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "byoc_trunk_sid", ""),
					resource.TestCheckResourceAttr(stateDataSourceName, "emergency.#", "1"),
					resource.TestCheckResourceAttr(stateDataSourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "secure"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "sip_registration"),
					resource.TestCheckResourceAttr(stateDataSourceName, "voice.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccDataSourceTwilioSIPDomain_basic(testData *acceptance.TestData, domainName string) string {
	return fmt.Sprintf(`
resource "twilio_sip_domain" "domain" {
  account_sid = "%s"
  domain_name = "%s"
}

data "twilio_sip_domain" "domain" {
  account_sid = twilio_sip_domain.domain.account_sid
  sid         = twilio_sip_domain.domain.sid
}
`, testData.AccountSid, domainName)
}
