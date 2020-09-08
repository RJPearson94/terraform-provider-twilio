package serverless

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/versions"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessFunctions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessFunctionsRead,

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
			"functions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"latest_version_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"content": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"visibility": {
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

func dataSourceServerlessFunctionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Functions.NewFunctionsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No functions were found for serverless service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read serverless function: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	functions := make([]interface{}, 0)

	for _, function := range paginator.Functions {
		d.Set("account_sid", function.AccountSid)

		functionMap := make(map[string]interface{})

		functionMap["sid"] = function.Sid
		functionMap["friendly_name"] = function.FriendlyName
		functionMap["date_created"] = function.DateCreated.Format(time.RFC3339)

		if function.DateUpdated != nil {
			functionMap["date_updated"] = function.DateUpdated.Format(time.RFC3339)
		}

		functionMap["url"] = function.URL

		versionsPaginator := client.Service(serviceSid).Function(function.Sid).Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
			PageSize: sdkUtils.Int(5),
		})
		// The twilio api return the latest version as the first element in the array.
		// So there is no need to loop to retrieve all records
		versionsPaginator.Next()

		if versionsPaginator.Error() != nil {
			return diag.Errorf("Failed to read serverless function versions: %s", versionsPaginator.Error().Error())
		}

		if len(versionsPaginator.Versions) > 0 {
			latestVersion := versionsPaginator.Versions[0]

			functionMap["latest_version_sid"] = latestVersion.Sid
			functionMap["path"] = latestVersion.Path
			functionMap["visibility"] = latestVersion.Visibility

			contentGetResponse, contentErr := client.Service(serviceSid).Function(function.Sid).Version(latestVersion.Sid).Content().FetchWithContext(ctx)
			if contentErr != nil {
				if utils.IsNotFoundError(contentErr) {
					return diag.Errorf("Function version with sid (%s) was not found for serverless service with sid (%s) and function with sid (%s)", latestVersion.Sid, serviceSid, function.Sid)
				}
				return diag.Errorf("Failed to read serverless function version content: %s", err.Error())
			}

			functionMap["content"] = contentGetResponse.Content
		} else {
			log.Printf("[INFO] No serverless function versions found for function (%s)", function.Sid)
		}

		functions = append(functions, functionMap)
	}

	d.Set("functions", &functions)

	return nil
}
