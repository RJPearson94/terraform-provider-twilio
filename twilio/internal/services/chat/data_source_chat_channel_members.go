package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatChannelMembers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceChatChannelMembersRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"channel_sid": {
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attributes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role_sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_consumed_message_index": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"last_consumption_timestamp": {
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
				},
			},
		},
	}
}

func dataSourceChatChannelMembersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

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
				return diag.Errorf("No channel members were found for chat service with sid (%s) and channel with sid (%s)", serviceSid, channelSid)
			}
		}
		return diag.Errorf("Failed to list chat channel members: %s", err.Error())
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
		if member.LastConsumedTimestamp != nil {
			memberMap["last_consumption_timestamp"] = member.LastConsumedTimestamp.Format(time.RFC3339)
		}
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
