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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this asset by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this asset",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
				Description:  "The SID of the Serverless service. Changing this forces a new resource",
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "A human-readable label for the asset",
			},
			"latest_version_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the latest version of the asset",
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
				Description:   "The path to the file containing the asset content. Conflicts with `content`",
			},
			"source_hash": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
				Description:   "A hash of the source file, used to detect changes. Conflicts with `content`",
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
				Description:   "The inline content of the asset. Conflicts with `source`",
			},
			"content_file_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
				Description:   "The file name to use when uploading inline content. Conflicts with `source`",
			},
			"content_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The MIME type of the asset content (e.g. `text/html`)",
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 255),
				Description:  "The URL path at which the asset will be accessible",
			},
			"visibility": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"protected",
					"private",
				}, false),
				Description: "The access control for the asset. Valid values are `public`, `protected`, `private`",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the asset was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the asset was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the asset resource",
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
