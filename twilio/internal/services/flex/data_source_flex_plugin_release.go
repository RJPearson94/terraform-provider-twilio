package flex

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlexPluginRelease() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFlexPluginReleaseRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.FlexPluginReleaseSidValidation(),
				Description:  "The SID of the Flex plugin release to look up",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this plugin release",
			},
			"configuration_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the plugin configuration associated with this release",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the plugin release was created, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the plugin release resource",
			},
		},
	}
}

func dataSourceFlexPluginReleaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Flex

	sid := d.Get("sid").(string)
	getResponse, err := client.PluginRelease(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Flex plugin release with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read flex plugin release: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("configuration_sid", getResponse.ConfigurationSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))
	d.Set("url", getResponse.URL)

	return nil
}
