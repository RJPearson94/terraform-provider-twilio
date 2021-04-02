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

const queueResourceName = "twilio_voice_queue"

func TestAccTwilioAccountQueue_basic(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "100"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				ResourceName:      stateResourceName,
				ImportState:       true,
				ImportStateIdFunc: testAccTwilioAccountQueueImportStateIdFunc(stateResourceName),
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTwilioAccountQueue_update(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	newFriendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "100"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
			{
				Config: testAccTwilioAccountQueue_basic(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
					resource.TestCheckResourceAttrSet(stateResourceName, "id"),
					resource.TestCheckResourceAttrSet(stateResourceName, "sid"),
					resource.TestCheckResourceAttrSet(stateResourceName, "account_sid"),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "100"),
					resource.TestCheckResourceAttrSet(stateResourceName, "average_wait_time"),
					resource.TestCheckResourceAttrSet(stateResourceName, "current_size"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_created"),
					resource.TestCheckResourceAttrSet(stateResourceName, "date_updated"),
				),
			},
		},
	})
}

func TestAccTwilioAccountQueue_maxSize(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	maxSize := 1
	newMaxSize := 5000

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "100"),
				),
			},
			{
				Config: testAccTwilioAccountQueue_maxSize(testData, friendlyName, maxSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "1"),
				),
			},
			{
				Config: testAccTwilioAccountQueue_maxSize(testData, friendlyName, newMaxSize),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "5000"),
				),
			},
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "max_size", "100"),
				),
			},
		},
	})
}

func TestAccTwilioAccountQueue_invalidMaxSizeOf0(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	maxSize := 0

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountQueue_maxSize(testData, friendlyName, maxSize),
				ExpectError: regexp.MustCompile(`(?s)expected max_size to be in the range \(1 - 5000\), got 0`),
			},
		},
	})
}

func TestAccTwilioAccountQueue_invalidMaxSizeOf10(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(10)
	maxSize := 5001

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountQueue_maxSize(testData, friendlyName, maxSize),
				ExpectError: regexp.MustCompile(`(?s)expected max_size to be in the range \(1 - 5000\), got 5001`),
			},
		},
	})
}

func TestAccTwilioAccountQueue_friendlyName(t *testing.T) {
	stateResourceName := fmt.Sprintf("%s.queue", queueResourceName)

	testData := acceptance.TestAccData
	friendlyName := acctest.RandString(1)
	newFriendlyName := acctest.RandString(64)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      testAccCheckTwilioAccountQueueDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTwilioAccountQueue_basic(testData, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", friendlyName),
				),
			},
			{
				Config: testAccTwilioAccountQueue_basic(testData, newFriendlyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTwilioAccountQueueExists(stateResourceName),
					resource.TestCheckResourceAttr(stateResourceName, "friendly_name", newFriendlyName),
				),
			},
		},
	})
}

func TestAccTwilioAccountQueue_invalidFriendlyNameWithLengthOf0(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountQueue_basic(testData, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got `),
			},
		},
	})
}

func TestAccTwilioAccountQueue_invalidFriendlyNameWithLengthOf65(t *testing.T) {
	testData := acceptance.TestAccData
	friendlyName := "7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountQueue_basic(testData, friendlyName),
				ExpectError: regexp.MustCompile(`(?s)expected length of friendly_name to be in the range \(1 - 64\), got 7y80krlx0npe98jtdhahyvx8jvfz09x21x226uxj8gowkun6dgl2p1xj819qjzgtt`),
			},
		},
	})
}

func TestAccTwilioAccountQueue_invalidAccountSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccTwilioAccountQueue_invalidAccountSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of account_sid to match regular expression "\^AC\[0-9a-fA-F\]\{32\}\$", got account_sid`),
			},
		},
	})
}

func testAccCheckTwilioAccountQueueDestroy(s *terraform.State) error {
	client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

	for _, rs := range s.RootModule().Resources {
		if rs.Type != queueResourceName {
			continue
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Queue(rs.Primary.ID).Fetch(); err != nil {
			if utils.IsNotFoundError(err) {
				return nil
			}
			return fmt.Errorf("Error occurred when retrieving queue information %s", err)
		}
	}

	return nil
}

func testAccCheckTwilioAccountQueueExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.TestAccProvider.Meta().(*common.TwilioClient).API

		// Ensure we have enough information in state to look up in API
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("Not found: %s", name)
		}

		if _, err := client.Account(rs.Primary.Attributes["account_sid"]).Queue(rs.Primary.ID).Fetch(); err != nil {
			return fmt.Errorf("Error occurred when retrieving queue information %s", err)
		}

		return nil
	}
}

func testAccTwilioAccountQueueImportStateIdFunc(name string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return "", fmt.Errorf("Not found: %s", name)
		}

		return fmt.Sprintf("/Accounts/%s/Queues/%s", rs.Primary.Attributes["account_sid"], rs.Primary.Attributes["sid"]), nil
	}
}

func testAccTwilioAccountQueue_basic(testData *acceptance.TestData, friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_voice_queue" "queue" {
  account_sid   = "%s"
  friendly_name = "%s"
}
`, testData.AccountSid, friendlyName)
}

func testAccTwilioAccountQueue_maxSize(testData *acceptance.TestData, friendlyName string, maxSize int) string {
	return fmt.Sprintf(`
resource "twilio_voice_queue" "queue" {
  account_sid   = "%s"
  friendly_name = "%s"
  max_size      = "%d"
}
`, testData.AccountSid, friendlyName, maxSize)
}

func testAccTwilioAccountQueue_invalidAccountSid() string {
	return `
resource "twilio_voice_queue" "queue" {
  account_sid   = "account_sid"
  friendly_name = "invalid_account_sid"
}
`
}
