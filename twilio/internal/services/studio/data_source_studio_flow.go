package studio

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

func dataSourceStudioFlow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.StudioFlowSidValidation(),
				Description:  "The SID of the Studio flow to read",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this flow",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable label for the flow",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle status of the flow. Either `draft` or `published`",
			},
			"definition": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string defining the flow structure",
			},
			"commit_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the changes made in the current flow revision",
			},
			"revision": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The current revision number of the flow",
			},
			"valid": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the flow definition passed Twilio's validation checks",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the flow was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the flow was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the flow resource",
			},
			"webhook_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The webhook URL used to trigger this flow via the REST API",
			},
		},
	}
}

func dataSourceStudioFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Studio

	sid := d.Get("sid").(string)
	getResponse, err := client.Flow(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Studio flow with sid (%s) was not found", sid)
		}
		return diag.Errorf("Failed to read studio flow: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)

	json, err := structure.FlattenJsonToString(getResponse.Definition)
	if err != nil {
		return diag.Errorf("Unable to flattern definition json to string")
	}
	d.Set("definition", json)
	d.Set("status", getResponse.Status)
	d.Set("revision", getResponse.Revision)
	d.Set("commit_message", getResponse.CommitMessage)
	d.Set("valid", getResponse.Valid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}
	d.Set("url", getResponse.URL)
	d.Set("webhook_url", getResponse.WebhookURL)

	return nil
}
