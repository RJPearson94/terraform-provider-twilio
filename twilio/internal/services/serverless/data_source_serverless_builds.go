package serverless

import (
	"context"
	"sort"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/builds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessBuilds() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessBuildsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"builds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
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
										Type:     schema.TypeString,
										Computed: true,
									},
									"asset_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_created": {
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
								},
							},
						},
						"function_versions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
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
										Type:     schema.TypeString,
										Computed: true,
									},
									"function_sid": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_created": {
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
								},
							},
						},
						"dependencies": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"runtime": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
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

func dataSourceServerlessBuildsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Builds.NewBuildsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No builds were found for serverless service with sid (%s)", serviceSid)
		}
		return diag.Errorf("Failed to read serverless build: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	builds := make([]interface{}, 0)

	for _, build := range paginator.Builds {
		d.Set("account_sid", build.AccountSid)

		buildMap := make(map[string]interface{})

		buildMap["sid"] = build.Sid
		buildMap["status"] = build.Status
		buildMap["runtime"] = build.Runtime
		buildMap["asset_versions"] = flattenAssetVersions(build.AssetVersions)
		buildMap["function_versions"] = flattenPageFunctionVersions(build.FunctionVersions)
		buildMap["dependencies"] = flattenPageDependencies(build.Dependencies)
		buildMap["date_created"] = build.DateCreated.Format(time.RFC3339)

		if build.DateUpdated != nil {
			buildMap["date_updated"] = build.DateUpdated.Format(time.RFC3339)
		}

		buildMap["url"] = build.URL

		builds = append(builds, buildMap)
	}

	d.Set("builds", &builds)

	return nil
}

func flattenAssetVersions(input *[]builds.PageAssetVersion) *[]interface{} {
	if input == nil {
		return nil
	}

	assetVersions := *input

	// Sort array in SID order due to values being returned in a random order if 2 or more resources are created at the same time
	sort.Slice(assetVersions[:], func(i, j int) bool {
		return assetVersions[i].Sid < assetVersions[j].Sid
	})

	results := make([]interface{}, 0)

	for _, prop := range assetVersions {
		result := make(map[string]interface{})
		result["sid"] = prop.Sid
		result["account_sid"] = prop.AccountSid
		result["service_sid"] = prop.ServiceSid
		result["asset_sid"] = prop.AssetSid
		result["date_created"] = prop.DateCreated.Format(time.RFC3339)
		result["path"] = prop.Path
		result["visibility"] = prop.Visibility

		results = append(results, result)
	}

	return &results
}

func flattenPageFunctionVersions(input *[]builds.PageFunctionVersion) *[]interface{} {
	if input == nil {
		return nil
	}

	functionVersions := *input

	// Sort array in SID order due to values being returned in a random order if 2 or more resources are created at the same time
	sort.Slice(functionVersions[:], func(i, j int) bool {
		return functionVersions[i].Sid < functionVersions[j].Sid
	})

	results := make([]interface{}, 0)

	for _, prop := range functionVersions {
		result := make(map[string]interface{})
		result["sid"] = prop.Sid
		result["account_sid"] = prop.AccountSid
		result["service_sid"] = prop.ServiceSid
		result["function_sid"] = prop.FunctionSid
		result["date_created"] = prop.DateCreated.Format(time.RFC3339)
		result["path"] = prop.Path
		result["visibility"] = prop.Visibility

		results = append(results, result)
	}

	return &results
}

func flattenPageDependencies(input *[]builds.PageDependency) map[string]string {
	if input == nil {
		return nil
	}

	results := make(map[string]string, 0)

	for _, prop := range *input {
		results[prop.Name] = prop.Version
	}

	return results
}
