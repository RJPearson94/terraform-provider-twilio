package flex

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	flex "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin/versions"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugins"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFlexPlugin() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlexPluginCreate,
		ReadContext:   resourceFlexPluginRead,
		UpdateContext: resourceFlexPluginUpdate,
		DeleteContext: resourceFlexPluginDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/PluginService/Plugins/(.*)"
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
			Update: schema.DefaultTimeout(10 * time.Minute),
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
			"unique_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"changelog": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plugin_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"private": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"version_archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"latest_version_sid": {
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

		CustomizeDiff: customdiff.All(
			customdiff.ComputedIf("latest_version_sid", func(_ context.Context, d *schema.ResourceDiff, meta interface{}) bool {
				for _, key := range []string{"changelog", "version", "plugin_url", "private"} {
					if d.HasChange(key) {
						return true
					}
				}
				return false
			}),
		),
	}
}

func resourceFlexPluginCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	createInput := &plugins.CreatePluginInput{
		UniqueName:   d.Get("unique_name").(string),
		Description:  utils.OptionalString(d, "description"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	createResult, err := client.Plugins.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create flex plugin: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if err := createPluginVersion(ctx, d, client); err != nil {
		return err
	}

	return resourceFlexPluginRead(ctx, d, meta)
}

func resourceFlexPluginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	getResponse, err := client.Plugin(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read flex plugin: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("archived", getResponse.Archived)
	d.Set("description", getResponse.Description)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	versionsPaginator := client.Plugin(d.Id()).Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	versionsPaginator.Next()

	if versionsPaginator.Error() != nil {
		return diag.Errorf("Failed to read flex plugin versions: %s", versionsPaginator.Error().Error())
	}

	if len(versionsPaginator.Versions) > 0 {
		latestVersion := versionsPaginator.Versions[0]

		d.Set("latest_version_sid", latestVersion.Sid)
		d.Set("changelog", latestVersion.Changelog)
		d.Set("version", latestVersion.Version)
		d.Set("plugin_url", latestVersion.PluginURL)
		d.Set("private", latestVersion.Private)
		d.Set("version_archived", latestVersion.Archived)
	} else {
		log.Printf("[INFO] No flex plugin versions found for plugin (%s)", getResponse.Sid)
	}

	return nil
}

func resourceFlexPluginUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	if d.HasChanges("description", "friendly_name") {
		updateInput := &plugin.UpdatePluginInput{
			Description:  utils.OptionalString(d, "description"),
			FriendlyName: utils.OptionalString(d, "friendly_name"),
		}

		updateResp, err := client.Plugin(d.Id()).UpdateWithContext(ctx, updateInput)
		if err != nil {
			return diag.Errorf("Failed to update flex plugin: %s", err.Error())
		}

		d.SetId(updateResp.Sid)
	}

	if d.HasChanges("changelog", "version", "plugin_url", "private") {
		if err := createPluginVersion(ctx, d, client); err != nil {
			return err
		}
	}

	return resourceFlexPluginRead(ctx, d, meta)
}

func resourceFlexPluginDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Flex plugin cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}

func createPluginVersion(ctx context.Context, d *schema.ResourceData, client *flex.Flex) diag.Diagnostics {
	createInput := &versions.CreateVersionInput{
		Changelog: utils.OptionalString(d, "changelog"),
		PluginURL: d.Get("plugin_url").(string),
		Private:   utils.OptionalBool(d, "private"),
		Version:   d.Get("version").(string),
	}

	if _, err := client.Plugin(d.Id()).Versions.CreateWithContext(ctx, createInput); err != nil {
		return diag.Errorf("Failed to create flex plugin version: %s", err.Error())
	}

	return nil

}
