package flex

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/flex/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/plugin_configurations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFlexPluginConfiguration() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlexPluginConfigurationCreate,
		ReadContext:   resourceFlexPluginConfigurationRead,
		DeleteContext: resourceFlexPluginConfigurationDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/PluginService/Configurations/(.*)"
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique SID assigned to this plugin configuration by Twilio",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this plugin configuration",
			},
			"archived": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the plugin configuration has been archived",
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The name of the plugin configuration. Changing this forces a new resource",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "A description of the plugin configuration. Changing this forces a new resource",
			},
			"plugins": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "A list of plugin versions to include in this configuration. Changing this forces a new resource",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plugin_version_sid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: utils.FlexPluginVersionSidValidation(),
							Description:  "The SID of the plugin version to include in the configuration. Changing this forces a new resource",
						},
						"plugin_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the plugin",
						},
						"plugin_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hosted URL of the plugin bundle",
						},
						"phase": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The load order phase of the plugin",
						},
						"private": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the plugin version is private",
						},
						"unique_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique name of the plugin",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version string of the plugin",
						},
						"date_created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the plugin was created, in RFC 3339 format",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The absolute URL of the plugin resource",
						},
					},
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the plugin configuration was created, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the plugin configuration resource",
			},
		},
	}
}

func resourceFlexPluginConfigurationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	createInput := &plugin_configurations.CreateConfigurationInput{
		Name:        d.Get("name").(string),
		Description: utils.OptionalString(d, "description"),
	}

	if v, ok := d.GetOk("plugins"); ok {
		plugins := make([]string, 0)
		for index := range v.([]interface{}) {
			plugins = append(plugins, fmt.Sprintf(`{"plugin_version":"%s"}`, d.Get(fmt.Sprintf("plugins.%d.plugin_version_sid", index)).(string)))
		}
		createInput.Plugins = &plugins
	}

	createResult, err := client.PluginConfigurations.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create flex plugin configuration: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexPluginConfigurationRead(ctx, d, meta)
}

func resourceFlexPluginConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	getResponse, err := client.PluginConfiguration(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read flex plugin configuration: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("archived", getResponse.Archived)
	d.Set("name", getResponse.Name)
	d.Set("description", getResponse.Description)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("url", getResponse.URL)

	paginator := client.PluginConfiguration(d.Id()).Plugins.NewPluginsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	paginatorErr := paginator.Error()
	if paginatorErr != nil {
		if utils.IsNotFoundError(paginatorErr) {
			return nil
		}
		return diag.Errorf("Failed to read flex plugin configuration plugins: %s", paginatorErr.Error())
	}

	d.Set("plugins", helper.FlattenPlugins(paginator.Plugins))

	return nil
}

func resourceFlexPluginConfigurationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	if _, err := client.PluginConfiguration(d.Id()).ArchiveWithContext(ctx); err != nil {
		return diag.Errorf("Failed to archive flex plugin configuration: %s", err.Error())
	}
	d.SetId("")
	return nil
}
