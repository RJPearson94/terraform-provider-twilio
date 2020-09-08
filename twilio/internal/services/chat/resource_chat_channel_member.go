package chat

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/member"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel/members"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceChatChannelMember() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatChannelMemberCreate,
		Read:   resourceChatChannelMemberRead,
		Update: resourceChatChannelMemberUpdate,
		Delete: resourceChatChannelMemberDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Channels/(.*)/Members/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 4 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("channel_sid", match[2])
				d.Set("sid", match[3])
				d.SetId(match[3])
				return []*schema.ResourceData{d}, nil
			},
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"channel_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"attributes": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"identity": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role_sid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceChatChannelMemberCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &members.CreateChannelMemberInput{
		Identity:   d.Get("identity").(string),
		Attributes: utils.OptionalJSONString(d, "attributes"),
		RoleSid:    utils.OptionalString(d, "role_sid"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Members.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat channel member: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceChatChannelMemberRead(d, meta)
}

func resourceChatChannelMemberRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Member(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				d.SetId("")
				return nil
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel member: %s", err)
	}

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

func resourceChatChannelMemberUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &member.UpdateChannelMemberInput{
		Attributes: utils.OptionalJSONString(d, "attributes"),
		RoleSid:    utils.OptionalString(d, "role_sid"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Member(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat channel member: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatChannelMemberRead(d, meta)
}

func resourceChatChannelMemberDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).Channel(d.Get("channel_sid").(string)).Member(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete chat channel member: %s", err.Error())
	}
	d.SetId("")
	return nil
}
