package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var pushCredentialsFCMResourceName = "twilio_conversations_push_credential_fcm"

func TestAccTwilioConversationsPushCredentialsFCM_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.push_credential_fcm", pushCredentialsFCMResourceName)
	friendlyName := acctest.RandString(10)
	secret := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsPushCredentialsFCMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "secret", secret),
					resource.TestCheckResourceAttr(stateResourceName, "type", "fcm"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				ResourceName:            stateResourceName,
				ImportState:             true,
				ImportStateIdFunc:       testAccTwilioConversationsPushCredentialsFCMImportStateIdFunc(stateResourceName),
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"secret"},
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.push_credential_fcm", pushCredentialsFCMResourceName)
	friendlyName := acctest.RandString(10)
	secret := acctest.RandString(10)
	newSecret := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsPushCredentialsFCMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "secret", secret),
					resource.TestCheckResourceAttr(stateResourceName, "type", "fcm"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, newSecret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateResourceName, "secret", newSecret),
					resource.TestCheckResourceAttr(stateResourceName, "type", "fcm"),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
					resource.TestCheckResourceAttrSet(stateResourceName, "url"),
				),
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.push_credential_fcm", pushCredentialsFCMResourceName)
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(64)
	secret := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsPushCredentialsFCMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(newFriendlyName, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_invalidFriendlyNameWith0Characters(t *testing.T) {
	friendlyName := ""
	secret := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_invalidFriendlyNameWith65Characters(t *testing.T) {
	friendlyName := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	secret := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`),
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_secret(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.push_credential_fcm", pushCredentialsFCMResourceName)
	friendlyName := acctest.RandString(1)
	secret := acctest.RandString(1)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioConversationsPushCredentialsFCMDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioConversationsPushCredentialsFCMExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "secret", secret),
				),
			},
		},
	})
}

func TestAccTwilioConversationsPushCredentialsFCM_blankSecret(t *testing.T) {
	friendlyName := "test"
	secret := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName, secret),
				ExpectError: regexp.MustCompile(`(?s)expected \"secret\" to not be an empty string, got `),
			},
		},
	})
}

func testAccCheckTwilioConversationsPushCredentialsFCMDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

	for _, rs := range s.RootModule().Resources {
		if rs.Type != pushCredentialsFCMResourceName {
			continue
		}

		if _, err := client.Credential(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving push credentials information %s", err.Error())
		}
	}

	return nil
}

func testAccCheckTwilioConversationsPushCredentialsFCMExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).Conversations

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Credential(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving push credentials information %s", err.Error())
		}

		return nil
	}
}

func testAccTwilioConversationsPushCredentialsFCMImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Credentials/%s", rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioConversationsPushCredentialsFCM_basic(friendlyName string, secret string) string {
	return fmt.Sprintf(`
resource "twilio_conversations_push_credential_fcm" "push_credential_fcm" {
  friendly_name = "%s"
  secret        = "%s"
}
`, friendlyName, secret)
}
