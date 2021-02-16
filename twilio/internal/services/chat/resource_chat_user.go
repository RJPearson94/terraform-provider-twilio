package chat

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/user"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/users"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceChatUser() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		CreateContext: resourceChatUserCreate,
		ReadContext:   resourceChatUserRead,
		UpdateContext: resourceChatUserUpdate,
		DeleteContext: resourceChatUserDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Services/(.*)/Users/(.*)"
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
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     validation.StringIsJSON,
				DiffSuppressFunc: structure.SuppressJsonDiff,
			},
			"identity": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_notifiable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_online": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"joined_channels_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"role_sid": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceChatUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	createInput := &users.CreateUserInput{
		Attributes:   utils.OptionalJSONString(d, "attributes"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		Identity:     d.Get("identity").(string),
		RoleSid:      utils.OptionalString(d, "role_sid"),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Users.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create chat user: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceChatUserRead(ctx, d, meta)
}

func resourceChatUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	getResponse, err := client.Service(d.Get("service_sid").(string)).User(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				d.SetId("")
				return nil
			}
		}
		return diag.Errorf("Failed to read chat user: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("attributes", getResponse.Attributes)
	d.Set("identity", getResponse.Identity)
	d.Set("is_notifiable", getResponse.IsNotifiable)
	d.Set("is_online", getResponse.IsOnline)
	d.Set("joined_channels_count", getResponse.JoinedChannelsCount)
	d.Set("role_sid", getResponse.RoleSid)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceChatUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	updateInput := &user.UpdateUserInput{
		Attributes:   utils.OptionalJSONString(d, "attributes"),
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		RoleSid:      utils.OptionalString(d, "role_sid"),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).User(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update chat user: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatUserRead(ctx, d, meta)
}

func resourceChatUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	if err := client.Service(d.Get("service_sid").(string)).User(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete chat user: %s", err.Error())
	}
	d.SetId("")
	return nil
}
