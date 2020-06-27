package serverless

import (
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployments"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessDeploymentCreate,
		Read:   resourceServerlessDeploymentRead,
		Delete: resourceServerlessDeploymentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"build_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"date_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerlessDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	createResult, err := createServerlessDeployment(d, meta, utils.OptionalString(d, "build_sid"))
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless deployment: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceServerlessDeploymentRead(d, meta)
}

func resourceServerlessDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless

	getResponse, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Deployment(d.Id()).Get()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless deployment: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("environment_sid", getResponse.EnvironmentSid)
	d.Set("build_sid", getResponse.BuildSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Serverless deployments cannot be deleted. So a new deployment will be create without a build sid as this will supersede the current deployment")

	if _, err := createServerlessDeployment(d, meta, nil); err != nil {
		return fmt.Errorf("[ERROR] Failed to create deployment without build sid: %s", err)
	}

	d.SetId("")
	return nil
}

func createServerlessDeployment(d *schema.ResourceData, meta interface{}, sid *string) (*deployments.CreateDeploymentOutput, error) {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &deployments.CreateDeploymentInput{
		BuildSid: sid,
	}

	return client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Deployments.Create(createInput)
}
