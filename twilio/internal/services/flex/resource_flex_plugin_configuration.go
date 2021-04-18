package flex

import (
	"context"
	"fmt"
	"log"
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
			"archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"plugins": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plugin_version_sid": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: utils.FlexPluginVersionSidValidation(),
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
	log.Printf("[INFO] Flex plugin configuration cannot be deleted, so removing from the Terraform state")

	d.SetId("")
	return nil
}
