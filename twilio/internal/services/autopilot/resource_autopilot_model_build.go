package autopilot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_build"
	"github.com/RJPearson94/twilio-sdk-go/service/autopilot/v1/assistant/model_builds"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAutopilotModelBuild() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutopilotModelBuildCreate,
		Read:   resourceAutopilotModelBuildRead,
		Update: resourceAutopilotModelBuildUpdate,
		Delete: resourceAutopilotModelBuildDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status_callback": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceAutopilotModelBuildCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &model_builds.CreateModelBuildInput{
		UniqueName:     utils.OptionalString(d, "unique_name"),
		StatusCallback: utils.OptionalString(d, "status_callback"),
	}

	createResult, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuilds.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create autopilot model build: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	pollings := d.Get("polling").([]interface{})
	if len(pollings) == 1 {
		if err := poll(d, meta.(*common.TwilioClient), pollings[0].(map[string]interface{})); err != nil {
			return err
		}
	}

	return resourceAutopilotModelBuildRead(d, meta)
}

func resourceAutopilotModelBuildRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot model build: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("status_callback", d.Get("status_callback").(string))
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

func resourceAutopilotModelBuildUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &model_build.UpdateModelBuildInput{
		UniqueName: utils.OptionalString(d, "unique_name"),
	}

	updateResp, err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update autopilot model build: %s", err.Error())
	}

	d.SetId(updateResp.Sid)

	// Unique Name changes do not require the model to be rebuild. So polling will not occur
	return resourceAutopilotModelBuildRead(d, meta)
}

func resourceAutopilotModelBuildDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete autopilot model build: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func poll(d *schema.ResourceData, client *common.TwilioClient, pollingConfig map[string]interface{}) error {
	if pollingConfig["enabled"].(bool) {
		for i := 0; i < pollingConfig["max_attempts"].(int); i++ {
			log.Printf("[INFO] Build Polling attempt # %v", i+1)

			ctx, cancel := context.WithTimeout(client.StopContext, d.Timeout(schema.TimeoutRead))
			defer cancel()

			getResponse, err := client.Autopilot.Assistant(d.Get("assistant_sid").(string)).ModelBuild(d.Id()).FetchWithContext(ctx)
			if err != nil {
				return fmt.Errorf("[ERROR] Failed to poll autopilot model build: %s", err)
			}

			if getResponse.Status == "failed" {
				return fmt.Errorf("[ERROR] Autopilot model build failed")
			}
			if getResponse.Status == "completed" {
				return nil
			}
			time.Sleep(time.Duration(pollingConfig["delay_in_ms"].(int)) * time.Millisecond)
		}
		return fmt.Errorf("[ERROR] Reached max polling attempts without a completed model build")
	}
	return nil
}
