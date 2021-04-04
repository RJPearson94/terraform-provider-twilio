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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerlessDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerlessDeploymentCreate,
		ReadContext:   resourceServerlessDeploymentRead,
		DeleteContext: resourceServerlessDeploymentDelete,

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
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
			},
			"environment_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
			},
			"build_sid": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessBuildSidValidation(),
			},
			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"is_latest_deployment": {
				Type:     schema.TypeBool,
				Computed: true,
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

func resourceServerlessDeploymentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createResult, err := createServerlessDeployment(ctx, d, meta, utils.OptionalString(d, "build_sid"))
	if err != nil {
		return diag.Errorf("Failed to create serverless deployment: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessDeploymentRead(ctx, d, meta)
}

func resourceServerlessDeploymentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	environmentsClient := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string))

	getResponse, err := environmentsClient.Deployment(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read serverless deployment: %s", err.Error())
	}

	deploymentsPaginator := environmentsClient.Deployments.NewDeploymentsPaginatorWithOptions(&deployments.DeploymentsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	deploymentsPaginator.Next()

	if deploymentsPaginator.Error() != nil {
		return diag.Errorf("Failed to read serverless deployments: %s", deploymentsPaginator.Error().Error())
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

func resourceServerlessDeploymentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	buildNeedsRemoving, err := doesBuildNeedsRemoving(ctx, d, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if *buildNeedsRemoving {
		log.Printf("[INFO] Serverless deployments cannot be deleted. So a new deployment will be created without a build sid as this will supersede the current deployment")

		if _, err := createServerlessDeployment(ctx, d, meta, nil); err != nil {
			return diag.Errorf("Failed to create deployment without build sid: %s", err.Error())
		}
	}

	d.SetId("")
	return nil
}

func createServerlessDeployment(ctx context.Context, d *schema.ResourceData, meta interface{}, sid *string) (*deployments.CreateDeploymentResponse, error) {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &deployments.CreateDeploymentInput{
		BuildSid: sid,
	}

	return client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).Deployments.CreateWithContext(ctx, createInput)
}

func doesBuildNeedsRemoving(ctx context.Context, d *schema.ResourceData, meta interface{}) (*bool, error) {
	client := meta.(*common.TwilioClient).Serverless

	resp, err := client.Service(d.Get("service_sid").(string)).Environment(d.Get("environment_sid").(string)).FetchWithContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to read serverless environment during deletion of the deployment: %s", err.Error())
	}

	return sdkUtils.Bool(resp.BuildSid != nil && *resp.BuildSid == d.Get("build_sid").(string)), nil
}
