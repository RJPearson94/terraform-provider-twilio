package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessDeployments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessDeploymentsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
				Description:  "The SID of the Serverless service",
			},
			"environment_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
				Description:  "The SID of the Serverless environment",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these deployments",
			},
			"deployments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of deployments belonging to the Serverless environment",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this deployment by Twilio",
						},
						"build_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the build deployed in this deployment",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the deployment was created, in RFC 3339 format",
						},
						"date_updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the deployment was last updated, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the deployment resource",
						},
					},
				},
			},
		},
	}
}

func dataSourceServerlessDeploymentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	paginator := client.Service(serviceSid).Environment(environmentSid).Deployments.NewDeploymentsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No deployments were found for serverless service with sid (%s) and environment with sid (%s)", serviceSid, environmentSid)
		}
		return diag.Errorf("Failed to read serverless deployment: %s", err.Error())
	}

	d.SetId(serviceSid + "/" + environmentSid)
	d.Set("service_sid", serviceSid)
	d.Set("environment_sid", environmentSid)

	deployments := make([]interface{}, 0)

	for _, deployment := range paginator.Deployments {
		d.Set("account_sid", deployment.AccountSid)

		deploymentMap := make(map[string]interface{})

		deploymentMap["sid"] = deployment.Sid
		deploymentMap["build_sid"] = deployment.BuildSid
		deploymentMap["date_created"] = deployment.DateCreated.Format(time.RFC3339)

		if deployment.DateUpdated != nil {
			deploymentMap["date_updated"] = deployment.DateUpdated.Format(time.RFC3339)
		}

		deploymentMap["url"] = deployment.URL

		deployments = append(deployments, deploymentMap)
	}

	d.Set("deployments", &deployments)

	return nil
}
