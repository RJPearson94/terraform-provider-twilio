package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessDeployment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessDeploymentRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessDeploymentSidValidation(),
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
			},
			"environment_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"build_sid": {
				Type:     schema.TypeString,
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

func dataSourceServerlessDeploymentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Environment(environmentSid).Deployment(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Deployment with sid (%s) was not found for serverless service with sid (%s) and environment with sid (%s)", sid, serviceSid, environmentSid)
		}
		return diag.Errorf("Failed to read serverless deployment: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
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
