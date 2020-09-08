package serverless

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerlessService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerlessServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"include_credentials": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ui_editable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"date_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceServerlessServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.Service(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Serverless service with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read serverless service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("include_credentials", getResponse.IncludeCredentials)
	d.Set("ui_editable", getResponse.UiEditable)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
