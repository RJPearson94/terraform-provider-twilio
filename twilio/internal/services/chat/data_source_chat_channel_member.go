package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatChannelMember() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatChannelMemberRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"attributes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"identity": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_consumed_message_index": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_consumption_timestamp": &schema.Schema{
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

func dataSourceChatChannelMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).Channel(channelSid).Member(sid).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] Channel member with sid (%s) was not found for chat service with sid (%s) and channel with sid (%s)", sid, serviceSid, channelSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel member: %s", err)
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("channel_sid", getResponse.ChannelSid)
	d.Set("attributes", getResponse.Attributes)
	d.Set("identity", getResponse.Identity)
	d.Set("role_sid", getResponse.RoleSid)
	d.Set("last_consumed_message_index", getResponse.LastConsumedMessageIndex)
	d.Set("last_consumption_timestamp", getResponse.LastConsumedTimestamp)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}
