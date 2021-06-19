package twiml

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/applications"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTwimlApps() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTwimlAppsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"apps": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"messaging": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status_callback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"voice": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"caller_id_lookup": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"fallback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"fallback_method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"method": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_callback_url": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status_callback_method": {
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
						"date_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTwimlAppsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	options := &applications.ApplicationsPageOptions{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
	}

	accountSid := d.Get("account_sid").(string)
	paginator := client.Account(accountSid).Applications.NewApplicationsPaginatorWithOptions(options)
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to list apps: %s", err.Error())
	}

	d.SetId(accountSid)
	d.Set("account_sid", accountSid)

	apps := make([]interface{}, 0)

	for _, app := range paginator.Applications {
		appMap := make(map[string]interface{})

		appMap["sid"] = app.Sid
		appMap["friendly_name"] = app.FriendlyName
		appMap["messaging"] = &[]interface{}{
			map[string]interface{}{
				"status_callback_url": app.MessageStatusCallback,
				"fallback_method":     app.SmsFallbackMethod,
				"fallback_url":        app.SmsFallbackURL,
				"method":              app.SmsMethod,
				"url":                 app.SmsURL,
			},
		}
		appMap["voice"] = &[]interface{}{
			map[string]interface{}{
				"caller_id_lookup":       app.VoiceCallerIDLookup,
				"fallback_method":        app.VoiceFallbackMethod,
				"fallback_url":           app.VoiceFallbackURL,
				"method":                 app.VoiceMethod,
				"url":                    app.VoiceURL,
				"status_callback_method": app.StatusCallbackMethod,
				"status_callback_url":    app.StatusCallback,
			},
		}
		appMap["date_created"] = app.DateCreated.Format(time.RFC3339)

		if app.DateUpdated != nil {
			appMap["date_updated"] = app.DateUpdated.Format(time.RFC3339)
		}

		apps = append(apps, appMap)
	}

	d.Set("apps", &apps)

	return nil
}
