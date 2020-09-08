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
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/versions"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/functions"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/mitchellh/go-homedir"
)

func resourceServerlessFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessFunctionCreate,
		Read:   resourceServerlessFunctionRead,
		Update: resourceServerlessFunctionUpdate,
		Delete: resourceServerlessFunctionDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Functions/(.*)"
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
				Required: true,
				ForceNew: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"latest_version_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
			},
			"source_hash": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
			},
			"content": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"source"},
			},
			"content_file_name": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
			},
			"content_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"protected",
					"private",
				}, false),
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
	}
}

func resourceServerlessFunctionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &functions.CreateFunctionInput{
		FriendlyName: d.Get("friendly_name").(string),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Functions.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless function: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if err := createFunctionVersion(ctx, d, client); err != nil {
		return err
	}

	return resourceServerlessFunctionRead(d, meta)
}

func resourceServerlessFunctionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	functionClient := client.Service(d.Get("service_sid").(string)).Function(d.Id())

	getResponse, err := functionClient.FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless function: %s", err.Error())
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

	versionsPaginator := functionClient.Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	versionsPaginator.Next()

	if versionsPaginator.Error() != nil {
		return fmt.Errorf("[ERROR] Failed to read serverless function versions: %s", versionsPaginator.Error().Error())
	}

	if len(versionsPaginator.Versions) > 0 {
		latestVersion := versionsPaginator.Versions[0]

		d.Set("latest_version_sid", latestVersion.Sid)
		d.Set("path", latestVersion.Path)
		d.Set("visibility", latestVersion.Visibility)

		contentGetResponse, contentErr := functionClient.Version(latestVersion.Sid).Content().FetchWithContext(ctx)
		if contentErr != nil {
			if utils.IsNotFoundError(contentErr) {
				d.SetId("")
				return nil
			}
			return fmt.Errorf("[ERROR] Failed to read serverless function version content: %s", err.Error())
		}

		d.Set("content", contentGetResponse.Content)
	} else {
		log.Printf("[INFO] No serverless function versions found for function (%s)", getResponse.Sid)
	}

	return nil
}

func resourceServerlessFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	if d.HasChange("friendly_name") {
		updateInput := &function.UpdateFunctionInput{
			FriendlyName: d.Get("friendly_name").(string),
		}

		updateResp, err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).UpdateWithContext(ctx, updateInput)
		if err != nil {
			return fmt.Errorf("Failed to update serverless function: %s", err.Error())
		}

		d.SetId(updateResp.Sid)
	}

	if d.HasChanges("source", "source_hash", "content", "content_file_name", "content_type", "path", "visibility") {
		if err := createFunctionVersion(ctx, d, client); err != nil {
			return err
		}
	}

	return resourceServerlessFunctionRead(d, meta)
}

func resourceServerlessFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete serverless function: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func createFunctionVersion(ctx context.Context, d *schema.ResourceData, client *serverless.Serverless) error {
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
			return fmt.Errorf("[ERROR] Error expanding homedir: %s", err)
		}
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("Error opening source: %s", err)
		}

		body = file
		fileName = file.Name()

		defer func() {
			err := file.Close()
			if err != nil {
				log.Printf("[WARN] Error closing source: %s", err)
			}
		}()
	}

	if body == nil || fileName == "" || contentType == "" {
		return fmt.Errorf("[ERROR] body (%v), file name (%v) and content type (%v) are all required", body, fileName, contentType)
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

	if _, err := client.Service(d.Get("service_sid").(string)).Function(d.Id()).Versions.CreateWithContext(ctx, createInput); err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless function version: %s", err.Error())
	}

	return nil
}
