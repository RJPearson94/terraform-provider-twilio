package tests

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var resourceName = "twilio_iam_api_key"

func TestAccTwilioAPIKey_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioAPIKey_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioAPIKey_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.api_key", resourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.TestAccProviders,
		CheckDestroy: testAccCheckTwilioAPIKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAPIKey_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", ""),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAPIKey_friendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAPIKeyExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func testAccCheckTwilioAPIKeyDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Twilio
	context := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != resourceName {
			continue
		}

		keyResponse, err := client.Keys.Get(context, rs.Primary.ID)

		if err != nil {
			if strings.Contains(err.Error(), fmt.Sprintf("%s.json was not found", rs.Primary.ID)) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving api key information %s", err)
		}
		if keyResponse != nil {
			return fmt.Errorf("API Key still exists")
		}

	}

	return nil
}

func testAccCheckTwilioAPIKeyExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Twilio
		context := context.Background()

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		_, err := client.Keys.Get(context, rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Error occurred when retrieving api key information %s", err)
		}

		return nil
	}
}

func testAccTwilioAPIKey_basic() string {
	return `
resource "twilio_iam_api_key" "api_key" {}
`
}

func testAccTwilioAPIKey_friendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_iam_api_key" "api_key" {
	friendly_name = "%s"
}
`, friendlyName)
}
