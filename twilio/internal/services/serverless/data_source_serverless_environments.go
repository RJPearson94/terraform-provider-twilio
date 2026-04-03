package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessEnvironments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessEnvironmentsRead,

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
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns these environments",
			},
			"environments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of environments belonging to the Serverless service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this environment by Twilio",
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
				},
			},
		},
	}
}

func dataSourceServerlessEnvironmentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Environments.NewEnvironmentsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No environments were found for serverless service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read serverless environment: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	environments := make([]interface{}, 0)

	for _, environment := range paginator.Environments {
		d.Set("account_sid", environment.AccountSid)

		environmentMap := make(map[string]interface{})

		environmentMap["sid"] = environment.Sid
		environmentMap["build_sid"] = environment.BuildSid
		environmentMap["unique_name"] = environment.UniqueName
		environmentMap["domain_suffix"] = environment.DomainSuffix
		environmentMap["domain_name"] = environment.DomainName
		environmentMap["date_created"] = environment.DateCreated.Format(time.RFC3339)

		if environment.DateUpdated != nil {
			environmentMap["date_updated"] = environment.DateUpdated.Format(time.RFC3339)
		}

		environmentMap["url"] = environment.URL

		environments = append(environments, environmentMap)
	}

	d.Set("environments", &environments)

	return nil
}
