package chat

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceChatUser() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "Programmable Chat API will reach the end of life on 25th July 2022 (except for Flex applications), please see https://www.twilio.com/changelog/programmable-chat-end-of-life for more information",

		ReadContext: dataSourceChatUserRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatUserSidValidation(),
				Description:  "The SID of the user",
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ChatServiceSidValidation(),
				Description:  "The SID of the Programmable Chat service",
			},
			"account_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the account that owns this user",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A human-readable label for the user",
			},
			"attributes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string of custom attributes for the user",
			},
			"identity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identity string of the user",
			},
			"is_notifiable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the user has a potentially valid push registration for the chat service",
			},
			"is_online": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the user is actively connected to the chat service",
			},
			"joined_channels_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of channels the user has joined",
			},
			"role_sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The SID of the role assigned to the user",
			},
			"date_created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the user was created, in RFC 3339 format",
			},
			"date_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time the user was last updated, in RFC 3339 format",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The absolute URL of the user resource",
			},
		},
	}
}

func dataSourceChatUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Chat

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).User(sid).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return diag.Errorf("User with sid (%s) was not found for chat service with sid (%s)", sid, serviceSid)
			}
		}
		return diag.Errorf("Failed to read chat user: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
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
