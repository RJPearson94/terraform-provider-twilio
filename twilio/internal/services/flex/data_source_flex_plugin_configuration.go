package flex

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/flex/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexPluginConfiguration() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFlexPluginConfigurationRead,

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.FlexPluginConfigurationSidValidation(),
				Description:  "The SID of the Flex plugin configuration to look up",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the plugin configuration",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the plugin configuration",
			},
			"plugins": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of plugins included in this configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plugin_version_sid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SID of the plugin version included in the configuration",
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

func dataSourceFlexPluginConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	sid := d.Get("sid").(string)
	getResponse, err := client.PluginConfiguration(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Flex plugin configuration with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read flex plugin configuration: %s", err.Error())
	}

	paginator := client.PluginConfiguration(sid).Plugins.NewPluginsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	paginatorErr := paginator.Error()
	if paginatorErr != nil {
		if utils.IsNotFoundError(paginatorErr) {
			return diag.Errorf("No flex plugins were found for plugin configuration with sid (%s)", sid)
		}
		return diag.Errorf("Failed to read flex plugin configuration plugins: %s", paginatorErr.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("archived", getResponse.Archived)
	d.Set("name", getResponse.Name)
	d.Set("description", getResponse.Description)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("url", getResponse.URL)
	d.Set("plugins", helper.FlattenPlugins(paginator.Plugins))

	return nil
}
