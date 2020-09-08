package serverless

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/builds"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerlessBuilds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerlessBuildsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"builds": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"asset_versions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"account_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"asset_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_created": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"path": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"visibility": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"function_versions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"account_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"function_sid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"date_created": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"path": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"visibility": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dependencies": &schema.Schema{
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"status": &schema.Schema{
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

func dataSourceServerlessBuildsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Builds.NewBuildsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No builds were found for serverless service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read serverless build: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	builds := make([]interface{}, 0)

	for _, build := range paginator.Builds {
		d.Set("account_sid", build.AccountSid)

		buildMap := make(map[string]interface{})

		buildMap["sid"] = build.Sid
		buildMap["status"] = build.Status
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

	results := make([]interface{}, 0)

	for _, prop := range *input {
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

	results := make([]interface{}, 0)

	for _, prop := range *input {
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
