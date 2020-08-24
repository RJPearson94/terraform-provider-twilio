package flex

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/flex/v1/channels"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFlexChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlexChannelCreate,
		Read:   resourceFlexChannelRead,
		Delete: resourceFlexChannelDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"chat_friendly_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"chat_unique_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"chat_user_friendly_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flex_flow_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"long_lived": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"pre_engagement_data": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"task_attributes": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"task_sid": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"user_sid": {
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
	}
}

func resourceFlexChannelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex

	createInput := &channels.CreateChannelInput{
		ChatFriendlyName:     d.Get("chat_friendly_name").(string),
		ChatUniqueName:       utils.OptionalString(d, "chat_unique_name"),
		ChatUserFriendlyName: d.Get("chat_user_friendly_name").(string),
		FlexFlowSid:          d.Get("flex_flow_sid").(string),
		Identity:             d.Get("identity").(string),
		LongLived:            utils.OptionalBool(d, "long_lived"),
		PreEngagementData:    utils.OptionalString(d, "pre_engagement_data"),
		Target:               utils.OptionalString(d, "target"),
		TaskAttributes:       utils.OptionalString(d, "task_attributes"),
		TaskSid:              utils.OptionalString(d, "task_sid"),
	}

	createResult, err := client.Channels.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create flex channel: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceFlexChannelRead(d, meta)
}

func resourceFlexChannelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex

	getResponse, err := client.Channel(d.Id()).Fetch()
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("[ERROR] Failed to read flex channel: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("chat_friendly_name", d.Get("chat_friendly_name").(string))
	d.Set("chat_unique_name", d.Get("chat_unique_name"))
	d.Set("chat_user_friendly_name", d.Get("chat_user_friendly_name").(string))
	d.Set("flex_flow_sid", getResponse.FlexFlowSid)
	d.Set("identity", d.Get("identity").(string))
	d.Set("long_lived", d.Get("long_lived").(bool))
	d.Set("pre_engagement_data", d.Get("pre_engagement_data"))
	d.Set("target", d.Get("target"))
	d.Set("task_attributes", d.Get("task_attributes"))
	d.Set("task_sid", getResponse.TaskSid)
	d.Set("user_sid", getResponse.UserSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceFlexChannelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Flex

	if err := client.Channel(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete flex channel: %s", err.Error())
	}
	d.SetId("")
	return nil
}
