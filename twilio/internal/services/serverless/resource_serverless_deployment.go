package serverless

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/environment/deployments"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessDeployment() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessDeploymentCreate,
		Read:   resourceServerlessDeploymentRead,
		Delete: resourceServerlessDeploymentDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Environments/(.*)/Deployments/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("environment_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"environment_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"build_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_latest_deployment": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerlessDeploymentCreate(d *schema.ResourceData, meta interface{}) error {
	createResult, err := createServerlessDeployment(d, meta, utils.OptionalString(d, "build_sid"), schema.TimeoutCreate)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless deployment: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessDeploymentRead(d, meta)
}

func resourceServerlessDeploymentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	environmentsClient := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string))

	getResponse, err := environmentsClient.Deployment(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless deployment: %s", err.Error())
	}

	deploymentsPaginator := environmentsClient.Deployments.NewDeploymentsPaginatorWithOptions(&deployments.DeploymentsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	deploymentsPaginator.Next()

	if deploymentsPaginator.Error() != nil {
		return fmt.Errorf("[ERROR] Failed to read serverless deployments: %s", deploymentsPaginator.Error().Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("environment_sid", getResponse.EnvironmentSid)
	d.Set("build_sid", getResponse.BuildSid)
	d.Set("is_latest_deployment", deploymentsPaginator.Deployments[0].Sid == getResponse.Sid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessDeploymentDelete(d *schema.ResourceData, meta interface{}) error {
	if d.Get("is_latest_deployment").(bool) {
		if d.Get("build_sid") != nil {
			log.Printf("[INFO] Serverless deployments cannot be deleted. So a new deployment will be create without a build sid as this will supersede the current deployment")

			if _, err := createServerlessDeployment(d, meta, nil, schema.TimeoutDelete); err != nil {
				return fmt.Errorf("[ERROR] Failed to create deployment without build sid: %s", err.Error())
			}
		} else {
			log.Printf("[INFO] Serverless deployment build sid is already nil, so removing this resource from Terraform state as you cannot currently delete a serverless deployment")
		}
	} else {
		log.Printf("[INFO] Serverless deployment is not the latest deployments, so skipping creating a new deployment without a build sid")
	}

	d.SetId("")
	return nil
}

func createServerlessDeployment(d *schema.ResourceData, meta interface{}, sid *string, timeoutKey string) (*deployments.CreateDeploymentResponse, error) {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(timeoutKey))
	defer cancel()

	createInput := &deployments.CreateDeploymentInput{
		BuildSid: sid,
	}

	return client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Deployments.CreateWithContext(ctx, createInput)
}
