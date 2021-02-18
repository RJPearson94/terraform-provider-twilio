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
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugins": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plugin_version_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"phase": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"private": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"unique_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
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
