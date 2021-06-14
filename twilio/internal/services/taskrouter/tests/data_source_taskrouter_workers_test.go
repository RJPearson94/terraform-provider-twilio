package tests

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const workersDataSourceName = "twilio_taskrouter_workers"

func TestAccDataSourceTwilioTaskRouterWorkers_basic(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_basic(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "2"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withActivityName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withActivityName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withActivitySid(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withActivitySid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withAvailable(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withAvailable(friendlyName, true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
			{
				// Needs to be set to unavailable so the worker can be deleted
				Config: testAccDataSourceTwilioTaskRouterWorkers_withAvailable(friendlyName, false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.#"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withFriendlyName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withFriendlyName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withTargetWorkerExpression(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withTargetWorkerExpression(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{\"skills\":[\"tester\"]}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withTaskQueueName(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withTaskQueueName(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{\"skills\":[\"tester\"]}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_withTaskQueueSid(t *testing.T) {
	stateDataSourceName := fmt.Sprintf("data.%s.workers", workersDataSourceName)
	friendlyName := acctest.RandString(10)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTwilioTaskRouterWorkers_withTaskQueueSid(friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(stateDataSourceName, "id"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "account_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workspace_sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.#", "1"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.sid"),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.friendly_name", friendlyName),
					resource.TestCheckResourceAttr(stateDataSourceName, "workers.0.attributes", "{\"skills\":[\"tester\"]}"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_created"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_updated"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.date_status_changed"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.available"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_name"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.activity_sid"),
					resource.TestCheckResourceAttrSet(stateDataSourceName, "workers.0.url"),
				),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_invalidWorkspaceSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkers_invalidWorkflowSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of workspace_sid to match regular expression "\^WS\[0-9a-fA-F\]\{32\}\$", got workspace_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_invalidActivitySid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkers_invalidActivitySid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of activity_sid to match regular expression "\^WA\[0-9a-fA-F\]\{32\}\$", got activity_sid`),
			},
		},
	})
}

func TestAccDataSourceTwilioTaskRouterWorkers_invalidTaskQueueSid(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.PreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceTwilioTaskRouterWorkers_invalidTaskQueueSid(),
				ExpectError: regexp.MustCompile(`(?s)expected value of task_queue_sid to match regular expression "\^WQ\[0-9a-fA-F\]\{32\}\$", got task_queue_sid`),
			},
		},
	})
}

func testAccDataSourceTwilioTaskRouterWorkers_basic(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withActivityName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	activity_sid = twilio_taskrouter_activity.activity.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	activity_name = twilio_taskrouter_activity.activity.friendly_name

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withActivitySid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	activity_sid = twilio_taskrouter_activity.activity.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	activity_sid = twilio_taskrouter_activity.activity.sid

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withAvailable(friendlyName string, available bool) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_activity" "activity" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	available = true
}

resource "twilio_taskrouter_activity" "activity_offline" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
	available = false
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	activity_sid = %[2]v ? twilio_taskrouter_activity.activity.sid : twilio_taskrouter_activity.activity_offline.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	available = twilio_taskrouter_activity.activity.available

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName, available)
}

func testAccDataSourceTwilioTaskRouterWorkers_withFriendlyName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	friendly_name = "%[1]s"

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withTaskQueueName(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	target_workers = "skills HAS 'tester'"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	attributes = jsonencode({
		"skills":["tester"]
	})
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	task_queue_name = twilio_taskrouter_task_queue.task_queue.friendly_name

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withTaskQueueSid(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	target_workers = "skills HAS 'tester'"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	attributes = jsonencode({
		"skills":["tester"]
	})
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	task_queue_sid = twilio_taskrouter_task_queue.task_queue.sid

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_withTargetWorkerExpression(friendlyName string) string {
	return fmt.Sprintf(`
resource "twilio_taskrouter_workspace" "workspace" {
  friendly_name          = "%[1]s"
  multi_task_enabled     = true
  prioritize_queue_order = "FIFO"
}

resource "twilio_taskrouter_task_queue" "task_queue" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	target_workers = "skills HAS 'tester'"
}

resource "twilio_taskrouter_worker" "worker" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s"
	attributes = jsonencode({
		"skills":["tester"]
	})
}

resource "twilio_taskrouter_worker" "worker_2" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
  friendly_name = "%[1]s-2"
}

data "twilio_taskrouter_workers" "workers" {
  workspace_sid = twilio_taskrouter_workspace.workspace.sid
	target_workers_expression = twilio_taskrouter_task_queue.task_queue.target_workers

	depends_on = [
		twilio_taskrouter_worker.worker,
		twilio_taskrouter_worker.worker_2,
	]
}
`, friendlyName)
}

func testAccDataSourceTwilioTaskRouterWorkers_invalidWorkflowSid() string {
	return `
data "twilio_taskrouter_workers" "workers" {
  workspace_sid = "workspace_sid"
}
`
}

func testAccDataSourceTwilioTaskRouterWorkers_invalidActivitySid() string {
	return `
data "twilio_taskrouter_workers" "workers" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	activity_sid = "activity_sid"
}
`
}

func testAccDataSourceTwilioTaskRouterWorkers_invalidTaskQueueSid() string {
	return `
data "twilio_taskrouter_workers" "workers" {
  workspace_sid = "WSaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	task_queue_sid = "task_queue_sid"
}
`
}
