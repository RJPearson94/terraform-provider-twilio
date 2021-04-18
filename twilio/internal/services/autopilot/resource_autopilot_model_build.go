package autopilot

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_build"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_builds"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAutopilotModelBuild() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAutopilotModelBuildCreate,
		ReadContext:   resourceAutopilotModelBuildRead,
		UpdateContext: resourceAutopilotModelBuildUpdate,
		DeleteContext: resourceAutopilotModelBuildDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Assistants/(.*)/ModelBuilds/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("assistant_sid", match[1])
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
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"unique_name_prefix": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_callback": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"triggers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"build_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_code": {
				Type:     schema.TypeInt,
				Computed: true,
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
							Default:  24,
						},
						"delay_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5000,
						},
					},
				},
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

func resourceAutopilotModelBuildCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	var uniqueName *string = nil

	if v, ok := d.GetOk("unique_name_prefix"); ok {
		uniqueName = sdkUtils.String(v.(string) + resource.UniqueId())
	}

	createInput := &model_builds.CreateModelBuildInput{
		UniqueName:     uniqueName,
		StatusCallback: utils.OptionalStringWithEmptyStringOnChange(d, "status_callback"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuilds.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create autopilot model build: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	pollings := d.Get("polling").([]interface{})
	if len(pollings) == 1 {
		if err := poll(ctx, d, meta.(*common.TwilioClient), pollings[0].(map[string]interface{})); err != nil {
			return err
		}
	}

	return resourceAutopilotModelBuildRead(ctx, d, meta)
}

func resourceAutopilotModelBuildRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read autopilot model build: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("build_duration", getResponse.BuildDuration)
	d.Set("status", getResponse.Status)
	d.Set("error_code", getResponse.ErrorCode)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)
	return nil
}

func resourceAutopilotModelBuildUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("unique_name_prefix") {
		return nil
	}

	client := meta.(*common.TwilioClient).Autopilot

	var uniqueName *string = nil

	if v, ok := d.GetOk("unique_name_prefix"); ok {
		uniqueName = sdkUtils.String(v.(string) + resource.UniqueId())
	}

	updateInput := &model_build.UpdateModelBuildInput{
		UniqueName: uniqueName,
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update autopilot model build: %s", err.Error())
	}

	d.SetId(updateResp.Sid)

	// Unique Name changes do not require the model to be rebuild. So polling will not occur
	return resourceAutopilotModelBuildRead(ctx, d, meta)
}

func resourceAutopilotModelBuildDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	if err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete autopilot model build: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func poll(ctx context.Context, d *schema.ResourceData, client *common.TwilioClient, pollingConfig map[string]interface{}) diag.Diagnostics {
	if pollingConfig["enabled"].(bool) {
		for i := 0; i < pollingConfig["max_attempts"].(int); i++ {
			log.Printf("[INFO] Build Polling attempt # %v", i+1)

			getResponse, err := client.Autopilot.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).FetchWithContext(ctx)
			if err != nil {
				return diag.Errorf("Failed to poll autopilot model build: %s", err.Error())
			}

			if getResponse.Status == "failed" {
				return diag.Errorf("Autopilot model build failed")
			}
			if getResponse.Status == "completed" {
				return nil
			}
			time.Sleep(time.Duration(pollingConfig["delay_in_ms"].(int)) * time.Millisecond)
		}
		return diag.Errorf("Reached max polling attempts without a completed model build")
	}
	return nil
}
