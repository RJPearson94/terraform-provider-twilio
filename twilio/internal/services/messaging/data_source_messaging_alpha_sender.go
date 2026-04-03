package messaging

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMessagingAlphaSender() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceMessagingAlphaSenderRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingAlphaSenderSidValidation(),
				Description:  "The SID of the alpha sender",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.MessagingServiceSidValidation(),
				Description:  "The SID of the messaging service the alpha sender is associated with",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this alpha sender",
			},
			"alpha_sender": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The alphanumeric sender ID string",
			},
			"capabilities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of capabilities for the alpha sender",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the alpha sender was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the alpha sender was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the alpha sender resource",
			},
		},
	}
}

func dataSourceMessagingAlphaSenderRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Messaging

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).AlphaSender(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Alpha sender with sid (%s) was not found for messaging service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read messaging alpha sender: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("capabilities", getResponse.Capabilities)
	d.Set("alpha_sender", getResponse.AlphaSender)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
