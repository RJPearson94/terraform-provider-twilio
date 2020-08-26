package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channel"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/channels"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceChatChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatChannelCreate,
		Read:   resourceChatChannelRead,
		Update: resourceChatChannelUpdate,
		Delete: resourceChatChannelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"public",
					"private",
				}, false),
				ForceNew: true,
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

func resourceChatChannelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutCreate))
	defer cancel()

	createInput := &channels.CreateChannelInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		UniqueName:   utils.OptionalString(d, "unique_name"),
		Attributes:   utils.OptionalString(d, "attributes"),
		Type:         utils.OptionalString(d, "type"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Channels.CreateWithContext(ctx, createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat channel: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceChatChannelRead(d, meta)
}

func resourceChatChannelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	getResponse, err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if twilioError.Status == 403 && twilioError.Message == "Service instance not found" {
				d.SetId("")
				return nil
			}
			if twilioError.IsNotFoundError() {
				d.SetId("")
				return nil
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat channel: %s", err)
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

func resourceChatChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutUpdate))
	defer cancel()

	updateInput := &channel.UpdateChannelInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		UniqueName:   utils.OptionalString(d, "unique_name"),
		Attributes:   utils.OptionalString(d, "attributes"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat channel: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatChannelRead(d, meta)
}

func resourceChatChannelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutDelete))
	defer cancel()

	if err := client.Service(d.Get("service_sid").(string)).Channel(d.Id()).DeleteWithContext(ctx); err != nil {
		return fmt.Errorf("Failed to delete chat channel: %s", err.Error())
	}
	d.SetId("")
	return nil
}
