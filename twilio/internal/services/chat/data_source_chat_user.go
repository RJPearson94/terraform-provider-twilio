package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceChatUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceChatUserRead,

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
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": &schema.Schema{
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
			"is_notifiable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_online": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"joined_channels_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"role_sid": &schema.Schema{
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

func dataSourceChatUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	getResponse, err := client.Service(serviceSid).User(sid).FetchWithContext(ctx)
	if err != nil {
		if twilioError, ok := err.(*sdkUtils.TwilioError); ok {
			// currently programmable chat returns a 403 if the service instance does not exist
			if (twilioError.Status == 403 && twilioError.Message == "Service instance not found") || twilioError.IsNotFoundError() {
				return fmt.Errorf("[ERROR] User with sid (%s) was not found for chat service with sid (%s)", sid, serviceSid)
			}
		}
		return fmt.Errorf("[ERROR] Failed to read chat user: %s", err)
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
