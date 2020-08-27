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
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/function/versions"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/mitchellh/go-homedir"
)

func resourceServerlessFunctionVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerlessFunctionVersionCreate,
		Read:   resourceServerlessFunctionVersionRead,
		Delete: resourceServerlessFunctionVersionDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Functions/(.*)/Versions/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("function_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
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
			"function_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
				ForceNew:      true,
			},
			"source_hash": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"content"},
				ForceNew:      true,
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
				ForceNew:      true,
			},
			"content_file_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"source"},
				ForceNew:      true,
			},
			"content_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerlessFunctionVersionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

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

	createResult, err := client.Service(d.Get("service_sid").(string)).Function(d.Get("function_sid").(string)).Versions.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create serverless function: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceServerlessFunctionVersionRead(d, meta)
}

func resourceServerlessFunctionVersionRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	functionVersionClient := client.Service(d.Get("service_sid").(string)).Function(d.Get("function_sid").(string)).Version(d.Id())

	getResponse, err := functionVersionClient.FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless function version: %s", err.Error())
	}

	contentGetResponse, contentErr := functionVersionClient.Content().FetchWithContext(ctx)
	if contentErr != nil {
		if utils.IsNotFoundError(contentErr) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read serverless function version content: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("function_sid", getResponse.FunctionSid)
	d.Set("content_type", d.Get("content_type").(string))
	d.Set("file_name", d.Get("content_file_name").(string))
	d.Set("content", contentGetResponse.Content)
	d.Set("source_hash", d.Get("source_hash").(string))
	d.Set("path", getResponse.Path)
	d.Set("visibility", getResponse.Visibility)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("url", getResponse.URL)

	return nil
}

func resourceServerlessFunctionVersionDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[INFO] Serverless function versions cannot be deleted. So the resource will remain until the function resource has been removed")

	d.SetId("")
	return nil
}
