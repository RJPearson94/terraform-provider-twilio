package flex

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/flex/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFlexFlow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFlexFlowRead,

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
			"channel_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"chat_service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_identity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"integration": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"channel": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"creation_on_message": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"flow_sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"retry_count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"workflow_sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"workspace_sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"integration_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"janitor_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"long_lived": &schema.Schema{
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

func dataSourceFlexFlowRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	sid := d.Get("sid").(string)
	getResponse, err := client.FlexFlow(sid).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] Flex flow with sid (%s) was not found", sid)
		}
		return fmt.Errorf("[ERROR] Failed to read flex channel: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("channel_type", getResponse.ChannelType)
	d.Set("chat_service_sid", getResponse.ChatServiceSid)
	d.Set("contact_identity", getResponse.ContactIdentity)
	d.Set("enabled", getResponse.Enabled)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("integration", helper.FlattenIntegration(getResponse.Integration))
	d.Set("integration_type", getResponse.IntegrationType)
	d.Set("janitor_enabled", getResponse.JanitorEnabled)
	d.Set("long_lived", getResponse.LongLived)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
