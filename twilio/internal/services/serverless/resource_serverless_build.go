package serverless

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/build"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/builds"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServerlessBuild() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessBuildCreate,
		Read:   resourceServerlessBuildRead,
		Update: resourceServerlessBuildUpdate,
		Delete: resourceServerlessBuildDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

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
				Required: true,
				ForceNew: true,
			},
			"asset_version": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Required: true,
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
			"function_version": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Required: true,
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
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"polling": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"max_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  30,
						},
						"delay_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1000,
						},
					},
				},
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
	}
}

func resourceServerlessBuildCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	dependencyArray := make([]builds.CreateDependency, 0)
	for key, value := range d.Get("dependencies").(map[string]interface{}) {
		dependencyArray = append(dependencyArray, builds.CreateDependency{
			Name:    key,
			Version: value.(string),
		})
	}

	dependencies, err := json.Marshal(dependencyArray)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to marshal dependencies: %s", err)
	}

	createInput := &builds.CreateBuildInput{
		AssetVersions:    expandVersionSids(d.Get("asset_version").([]interface{})),
		FunctionVersions: expandVersionSids(d.Get("function_version").([]interface{})),
		Dependencies:     sdkUtils.String(string(dependencies)),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Builds.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless build: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	pollings := d.Get("polling").([]interface{})
	if len(pollings) == 1 {
		if err := poll(d, meta.(*common.TwilioClient), pollings[0].(map[string]interface{})); err != nil {
			return err
		}
	}

	return resourceServerlessBuildRead(d, meta)
}

func resourceServerlessBuildRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Get("service_sid").(string)).Build(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless build: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("asset_version", flatternAssetVersions(getResponse.AssetVersions))
	d.Set("function_version", flatternFunctionVersions(getResponse.FunctionVersions))
	d.Set("dependencies", flatternDependencies(getResponse.Dependencies))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("status", getResponse.Status)
	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessBuildUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Serverless deployments cannot be updated. So only polling config can be updated without a new resource being created")

	return nil
}

func resourceServerlessBuildDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).Build(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete serverless build: %s", err.Error())
	}

	d.SetId("")
	return nil
}

func expandVersionSids(input []interface{}) *[]string {
	versionSids := make([]string, 0)
	for _, version := range input {
		versionMap := version.(map[string]interface{})
		versionSids = append(versionSids, versionMap["sid"].(string))
	}
	return &versionSids
}

func flatternAssetVersions(input *[]build.FetchAssetVersion) *[]interface{} {
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

func flatternFunctionVersions(input *[]build.FetchFunctionVersion) *[]interface{} {
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

func flatternDependencies(input *[]build.FetchDependency) *[]interface{} {
	if input == nil {
		return nil
	}

	results := make([]interface{}, 0)

	for _, prop := range *input {
		result := make(map[string]interface{})
		result["name"] = prop.Name
		result["version"] = prop.Version

		results = append(results, result)
	}

	return &results
}

func poll(d *schema.ResourceData, client *common.TwilioClient, pollingConfig map[string]interface{}) error {
	if pollingConfig["enabled"].(bool) {
		for i := 0; i < pollingConfig["max_attempts"].(int); i++ {
			log.Printf("[INFO] Build Polling attempt # %v", i+1)
			ctx, cancel := context.WithTimeout(client.StopContext, d.Timeout(schema.TimeoutRead))
			defer cancel()

			getResponse, err := client.Serverless.Service(d.Get("service_sid").(string)).Build(d.Id()).FetchWithContext(ctx)
			if err != nil {
				return fmt.Errorf("[ERROR] Failed to poll serverless build: %s", err.Error())
			}

			if getResponse.Status == "failed" {
				return fmt.Errorf("[ERROR] Serverless build failed")
			}
			if getResponse.Status == "completed" {
				return nil
			}
			time.Sleep(time.Duration(pollingConfig["delay_in_ms"].(int)) * time.Millisecond)
		}
		return fmt.Errorf("[ERROR] Reached max polling attempts without a completed build")
	}
	return nil
}
