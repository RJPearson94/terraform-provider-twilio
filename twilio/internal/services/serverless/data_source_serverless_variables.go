package serverless

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerlessVariables() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerlessVariablesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"variables": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func dataSourceServerlessVariablesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	paginator := client.Service(serviceSid).Environment(environmentSid).Variables.NewVariablesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No variables were found for serverless service with sid (%s) and environment with sid (%s)", serviceSid, environmentSid)
		}
		return fmt.Errorf("[ERROR] Failed to read serverless variable: %s", err.Error())
	}

	d.SetId(serviceSid + "/" + environmentSid)
	d.Set("service_sid", serviceSid)
	d.Set("environment_sid", environmentSid)

	variables := make([]interface{}, 0)

	for _, variable := range paginator.Variables {
		d.Set("account_sid", variable.AccountSid)

		variableMap := make(map[string]interface{})

		variableMap["sid"] = variable.Sid
		variableMap["key"] = variable.Key
		variableMap["value"] = variable.Value
		variableMap["date_created"] = variable.DateCreated.Format(time.RFC3339)

		if variable.DateUpdated != nil {
			variableMap["date_updated"] = variable.DateUpdated.Format(time.RFC3339)
		}

		variableMap["url"] = variable.URL

		variables = append(variables, variableMap)
	}

	d.Set("variables", &variables)

	return nil
}
