package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/serverless/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessBuild() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessBuildRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessBuildSidValidation(),
				Description:  "The SID of the Serverless build",
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
				Description: "The SID of the account that owns this build",
			},
			"asset_versions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of asset versions included in this build",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this asset version by Twilio",
						},
						"account_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the account that owns this asset version",
						},
						"service_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the Serverless service that owns this asset version",
						},
						"asset_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the asset that this version belongs to",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the asset version was created, in RFC 3339 format",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL path of the asset version",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access control for the asset version",
						},
					},
				},
			},
			"function_versions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of function versions included in this build",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique SID assigned to this function version by Twilio",
						},
						"account_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the account that owns this function version",
						},
						"service_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the Serverless service that owns this function version",
						},
						"function_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the function that this version belongs to",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the function version was created, in RFC 3339 format",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL path of the function version",
						},
						"visibility": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The access control for the function version",
						},
					},
				},
			},
			"dependencies": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A map of npm package names to version ranges included in the build",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"runtime": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Node.js runtime version used for the build",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the build",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the build was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the build was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the build resource",
			},
		},
	}
}

func dataSourceServerlessBuildRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Build(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Build with sid (%s) was not found for serverless service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read serverless build: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("asset_versions", helper.FlattenAssetVersions(getResponse.AssetVersions))
	d.Set("function_versions", helper.FlattenFunctionVersions(getResponse.FunctionVersions))
	d.Set("dependencies", helper.FlattenDependencies(getResponse.Dependencies))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("runtime", getResponse.Runtime)
	d.Set("status", getResponse.Status)
	d.Set("url", getResponse.URL)

	return nil
}
