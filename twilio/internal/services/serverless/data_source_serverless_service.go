package serverless

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessService() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessServiceRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
				ExactlyOneOf: []string{"sid", "unique_name"},
			},
			"unique_name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"sid", "unique_name"},
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"include_credentials": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ui_editable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_base": {
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
	}
}

func dataSourceServerlessServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	var identifier string

	if v, ok := d.GetOk("sid"); ok {
		identifier = v.(string)
	} else {
		identifier = d.Get("unique_name").(string)
	}

	getResponse, err := client.Service(identifier).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Serverless service with sid/ unique name (%s) was not found", identifier)
		}
		return diag.Errorf("Failed to read serverless service: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("include_credentials", getResponse.IncludeCredentials)
	d.Set("ui_editable", getResponse.UiEditable)
	d.Set("domain_base", getResponse.DomainBase)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
