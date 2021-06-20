package flex

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configurations"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_releases"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlexPluginRelease() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlexPluginReleaseCreate,
		ReadContext:   resourceFlexPluginReleaseRead,
		DeleteContext: resourceFlexPluginReleaseDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/PluginService/Releases/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
				return []*schema.ResourceData{d}, nil
			},
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
			"configuration_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.FlexPluginConfigurationSidValidation(),
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

func resourceFlexPluginReleaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createResult, err := createRelease(ctx, d, meta, d.Get("configuration_sid").(string))
	if err != nil {
		return diag.Errorf("Failed to create flex plugin release: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexPluginReleaseRead(ctx, d, meta)
}

func resourceFlexPluginReleaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	getResponse, err := client.PluginRelease(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read flex plugin release: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("configuration_sid", getResponse.ConfigurationSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("url", getResponse.URL)

	return nil
}

func resourceFlexPluginReleaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	releasesPaginator := client.PluginReleases.NewReleasesPaginatorWithOptions(&plugin_releases.ReleasesPageOptions{
		PageSize: sdkUtils.Int(5),
	})

	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	releasesPaginator.Next()

	if releasesPaginator.Error() != nil {
		return diag.Errorf("Failed to read flex plugin releases: %s", releasesPaginator.Error().Error())
	}

	isCurrentRelease := d.Id() == releasesPaginator.Releases[0].Sid
	if isCurrentRelease {
		log.Printf("[INFO] Flex plugin release is current so a new default configuration will be created with a new release")

		defaultConfigResp, err := createDefaultConfiguration(ctx, d, meta)
		if err != nil {
			return diag.Errorf("Failed to create default configuration during release deletion: %s", err.Error())
		}

		if _, err := createRelease(ctx, d, meta, defaultConfigResp.Sid); err != nil {
			return diag.Errorf("Failed to create new flex plugin release deletion: %s", err.Error())
		}
	}

	d.SetId("")
	return nil
}

func createDefaultConfiguration(ctx context.Context, d *schema.ResourceData, meta interface{}) (*plugin_configurations.CreateConfigurationResponse, error) {
	client := meta.(*common.TwilioClient).Flex

	createInput := &plugin_configurations.CreateConfigurationInput{
		Name: fmt.Sprintf("Default Configuration to supersede release %s", d.Id()),
	}
	return client.PluginConfigurations.CreateWithContext(ctx, createInput)
}

func createRelease(ctx context.Context, d *schema.ResourceData, meta interface{}, configurationSid string) (*plugin_releases.CreateReleaseResponse, error) {
	client := meta.(*common.TwilioClient).Flex

	createInput := &plugin_releases.CreateReleaseInput{
		ConfigurationId: configurationSid,
	}

	return client.PluginReleases.CreateWithContext(ctx, createInput)
}
