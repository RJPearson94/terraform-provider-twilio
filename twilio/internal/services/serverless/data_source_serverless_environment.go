package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessEnvironment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessEnvironmentRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessEnvironmentSidValidation(),
				Description:  "The SID of the Serverless environment",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
				Description:  "The SID of the Serverless service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this environment",
			},
			"build_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the build currently deployed to this environment",
			},
			"unique_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique, developer-assigned name for the environment",
			},
			"domain_suffix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A URL-friendly suffix appended to the environment's domain name",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The fully qualified domain name for this environment",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the environment was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the environment was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the environment resource",
			},
		},
	}
}

func dataSourceServerlessEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Environment(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Environment with sid (%s) was not found for serverless service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read serverless environment: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("build_sid", getResponse.BuildSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("domain_suffix", getResponse.DomainSuffix)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
