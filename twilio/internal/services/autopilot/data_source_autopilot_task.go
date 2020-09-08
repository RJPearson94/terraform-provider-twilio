package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
)

func dataSourceAutopilotTask() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotTaskRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"assistant_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"actions_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"actions": &schema.Schema{
				Type:     schema.TypeString,
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

func dataSourceAutopilotTaskRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Assistant(assistantSid).Task(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Task with sid (%s) was not found for assistant with sid (%s)", sid, assistantSid)
		}
		return fmt.Errorf("[ERROR] Failed to read autopilot task: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("assistant_sid", getResponse.AssistantSid)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("actions_url", getResponse.ActionsURL)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	getActionsResponse, err := client.Assistant(getResponse.AssistantSid).Task(getResponse.Sid).Actions().FetchWithContext(ctx)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to read autopilot task actions: %s", err.Error())
	}
	actionsJSONString, err := structure.FlattenJsonToString(getActionsResponse.Data)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to flatten actions json to string: %s", err.Error())
	}
	d.Set("actions", actionsJSONString)

	return nil
}
