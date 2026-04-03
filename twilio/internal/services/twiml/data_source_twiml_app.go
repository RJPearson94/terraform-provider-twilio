package twiml

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/twiml/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTwimlApp() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTwimlAppRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ApplicationSidValidation(),
				Description:  "The SID of the TwiML application to look up",
			},
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
				Description:  "The SID of the account that owns this TwiML application",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the TwiML application",
			},
			"messaging": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The messaging settings for the TwiML application",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status_callback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL called for messaging status callback events",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming messages",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the messaging fallback URL",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when the application receives an incoming message",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the messaging URL",
						},
					},
				},
			},
			"voice": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The voice settings for the TwiML application",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"caller_id_lookup": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether caller ID lookup is enabled for incoming voice calls",
						},
						"fallback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when an error occurs retrieving or executing the TwiML for incoming voice calls",
						},
						"fallback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice fallback URL",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL to call when the application receives an incoming voice call",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice URL",
						},
						"status_callback_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL called for voice status callback events",
						},
						"status_callback_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP method used to call the voice status callback URL",
						},
					},
				},
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the TwiML application was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the TwiML application was last updated, in RFC 3339 format",
			},
		},
	}
}

func dataSourceTwimlAppRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Application(d.Get("sid").(string)).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read application: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("messaging", helper.FlattenMessaging(getResponse))
	d.Set("voice", helper.FlattenVoice(getResponse))
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	return nil
}
