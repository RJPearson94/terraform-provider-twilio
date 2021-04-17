package chat

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channels"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChatChannel() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		CreateContext: resourceChatChannelCreate,
		ReadContext:   resourceChatChannelRead,
		UpdateContext: resourceChatChannelUpdate,
		DeleteContext: resourceChatChannelDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Channels/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("service_sid", match[1])
				d.Set("sid", match[2])
				d.SetId(match[2])
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
			"sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unique_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{}",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "public", // Replicating the API default to prevent breaking existing apps but I would usually set the default to private
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"private",
				}, false),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"members_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"messages_count": {
				Type:     schema.TypeInt,
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
	}
}

func resourceChatChannelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	createInput := &channels.CreateChannelInput{
		FriendlyName: utils.OptionalStringWithEmptyStringDefault(d, "friendly_name"),
		UniqueName:   utils.OptionalStringWithEmptyStringDefault(d, "unique_name"),
		Attributes:   utils.OptionalString(d, "attributes"),
		Type:         utils.OptionalString(d, "type"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Channels.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create chat channel: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceChatChannelRead(ctx, d, meta)
}

func resourceChatChannelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	getResponse, err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				d.SetId("")
				return nil
			}
		}
		return diag.Errorf("Failed to read chat channel: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("unique_name", getResponse.UniqueName)
	d.Set("attributes", getResponse.Attributes)
	d.Set("type", getResponse.Type)
	d.Set("created_by", getResponse.CreatedBy)
	d.Set("members_count", getResponse.MembersCount)
	d.Set("messages_count", getResponse.MessagesCount)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceChatChannelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	updateInput := &channel.UpdateChannelInput{
		FriendlyName: utils.OptionalStringWithEmptyStringDefault(d, "friendly_name"),
		UniqueName:   utils.OptionalStringWithEmptyStringDefault(d, "unique_name"),
		Attributes:   utils.OptionalString(d, "attributes"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update chat channel: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatChannelRead(ctx, d, meta)
}

func resourceChatChannelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	if err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete chat channel: %s", err.Error())
	}
	d.SetId("")
	return nil
}
