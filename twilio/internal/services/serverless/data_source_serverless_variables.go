package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessVariables() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessVariablesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"variables": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
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

func dataSourceServerlessVariablesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	environmentSid := d.Get("environment_sid").(string)
	paginator := client.Service(serviceSid).Environment(environmentSid).Variables.NewVariablesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No variables were found for serverless service with sid (%s) and environment with sid (%s)", serviceSid, environmentSid)
		}
		return diag.Errorf("Failed to read serverless variable: %s", err.Error())
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
