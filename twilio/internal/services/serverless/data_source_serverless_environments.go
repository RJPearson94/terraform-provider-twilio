package serverless

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerlessEnvironments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerlessEnvironmentsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"build_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_suffix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"domain_name": {
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
				},
			},
		},
	}
}

func dataSourceServerlessEnvironmentsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Environments.NewEnvironmentsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No environments were found for serverless service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read serverless environment: %s", err.Error())
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
