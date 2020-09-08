package serverless

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	serverless "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/assets"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/mitchellh/go-homedir"
)

func resourceServerlessAsset() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServerlessAssetCreate,
		ReadContext:   resourceServerlessAssetRead,
		UpdateContext: resourceServerlessAssetUpdate,
		DeleteContext: resourceServerlessAssetDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Assets/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
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
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"latest_version_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
			},
			"source_hash": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
			},
			"content_file_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"protected",
					"private",
				}, false),
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

		CustomizeDiff: customdiff.All(
			customdiff.ComputedIf("latest_version_sid", func(_ context.Context, d *schema.ResourceDiff, meta interface{}) bool {
				for _, key := range []string{"source", "source_hash", "content", "content_file_name", "content_type", "path", "visibility"} {
					if d.HasChange(key) {
						return true
					}
				}
				return false
			}),
		),
	}
}

func resourceServerlessAssetCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	createInput := &assets.CreateAssetInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Assets.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create serverless asset: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if err := createAssetVersion(ctx, d, client); err != nil {
		return err
	}

	return resourceServerlessAssetRead(ctx, d, meta)
}

func resourceServerlessAssetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	assetClient := client.Service(d.Get("service_sid").(string)).Asset(d.Id())

	getResponse, err := assetClient.FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read serverless asset: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	versionsPaginator := assetClient.Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	versionsPaginator.Next()

	if versionsPaginator.Error() != nil {
		return diag.Errorf("Failed to read serverless asset versions: %s", versionsPaginator.Error().Error())
	}

	if len(versionsPaginator.Versions) > 0 {
		latestVersion := versionsPaginator.Versions[0]

		d.Set("latest_version_sid", latestVersion.Sid)
		d.Set("path", latestVersion.Path)
		d.Set("visibility", latestVersion.Visibility)
	} else {
		log.Printf("[INFO] No serverless asset versions found for asset (%s)", getResponse.Sid)
	}

	return nil
}

func resourceServerlessAssetUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	if d.HasChange("friendly_name") {
		updateInput := &asset.UpdateAssetInput{
			FriendlyName: d.Get("friendly_name").(string),
		}

		updateResp, err := client.Service(d.Get("service_sid").(string)).Asset(d.Id()).UpdateWithContext(ctx, updateInput)
		if err != nil {
			return diag.Errorf("Failed to update serverless asset: %s", err.Error())
		}

		d.SetId(updateResp.Sid)
	}

	if d.HasChanges("source", "source_hash", "content", "content_file_name", "content_type", "path", "visibility") {
		if err := createAssetVersion(ctx, d, client); err != nil {
			return err
		}
	}

	return resourceServerlessAssetRead(ctx, d, meta)
}

func resourceServerlessAssetDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	if err := client.Service(d.Get("service_sid").(string)).Asset(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete serverless asset: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func createAssetVersion(ctx context.Context, d *schema.ResourceData, client *serverless.Serverless) diag.Diagnostics {
	var body io.ReadSeeker
	var fileName string
	var contentType = d.Get("content_type").(string)

	if value, ok := d.GetOk("content"); ok {
		body = strings.NewReader(value.(string))
		fileName = d.Get("content_file_name").(string)
	}

	if value, ok := d.GetOk("source"); ok {
		path, err := homedir.Expand(value.(string))
		if err != nil {
			return diag.Errorf("Error expanding homedir: %s", err.Error())
		}
		file, err := os.Open(path)
		if err != nil {
			return diag.Errorf("Error opening source: %s", err.Error())
		}

		body = file
		fileName = file.Name()

		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("[WARN] Error closing source: %s", err.Error())
			}
		}()
	}

	if body == nil || fileName == "" || contentType == "" {
		return diag.Errorf("body (%v), file name (%v) and content type (%v) are all required", body, fileName, contentType)
	}

	createInput := &versions.CreateVersionInput{
		Content: versions.CreateContentDetails{
			Body:        body,
			ContentType: contentType,
			FileName:    fileName,
		},
		Path:       d.Get("path").(string),
		Visibility: d.Get("visibility").(string),
	}

	if _, err := client.Service(d.Get("service_sid").(string)).Asset(d.Id()).Versions.CreateWithContext(ctx, createInput); err != nil {
		return diag.Errorf("Failed to create serverless asset version: %s", err.Error())
	}

	return nil
}
