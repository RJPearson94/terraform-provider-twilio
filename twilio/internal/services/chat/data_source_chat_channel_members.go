package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatChannelMembers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatChannelMembersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
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
			"members": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
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
				},
			},
		},
	}
}

func dataSourceChatChannelMembersRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	channelSid := d.Get("channel_sid").(string)
	paginator := client.Service(serviceSid).Channel(channelSid).Members.NewChannelMembersPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] No channel members were found for chat service with sid (%s) and channel with sid (%s)", serviceSid, channelSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to list chat channel members: %s", err)
	}

	d.SetId(serviceSid + "/" + channelSid)
	d.Set("service_sid", serviceSid)
	d.Set("channel_sid", channelSid)

	members := make([]interface{}, 0)

	for _, member := range paginator.Members {
		d.Set("account_sid", member.AccountSid)

		memberMap := make(map[string]interface{})

		memberMap["sid"] = member.Sid
		memberMap["attributes"] = member.Attributes
		memberMap["identity"] = member.Identity
		memberMap["role_sid"] = member.RoleSid
		memberMap["last_consumed_message_index"] = member.LastConsumedMessageIndex
		memberMap["last_consumption_timestamp"] = member.LastConsumedTimestamp
		memberMap["date_created"] = member.DateCreated.Format(time.RFC3339)

		if member.DateUpdated != nil {
			memberMap["date_updated"] = member.DateUpdated.Format(time.RFC3339)
		}

		memberMap["url"] = member.URL

		members = append(members, memberMap)
	}

	d.Set("members", &members)

	return nil
}
