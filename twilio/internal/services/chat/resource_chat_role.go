package chat

import (
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/role"
	"github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service/roles"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceChatRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceChatRoleCreate,
		Read:   resourceChatRoleRead,
		Update: resourceChatRoleUpdate,
		Delete: resourceChatRoleDelete,
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
			"service_sid": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"channel",
					"deployment",
				}, false),
			},
			"permissions": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func resourceChatRoleCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	createInput := &roles.CreateRoleInput{
		FriendlyName: d.Get("friendly_name").(string),
		Type:         d.Get("type").(string),
		Permission:   utils.ConvertToStringSlice(d.Get("permissions").([]interface{})),
	}

	createResult, err := client.Service(d.Get("service_sid").(string)).Roles.Create(createInput)
	if err != nil {
		return fmt.Errorf("[ERROR] Failed to create chat role: %s", err)
	}

	d.SetId(createResult.Sid)
	return resourceChatRoleRead(d, meta)
}

func resourceChatRoleRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	getResponse, err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).Fetch()
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
		return fmt.Errorf("[ERROR] Failed to read chat role: %s", err)
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("type", getResponse.Type)
	d.Set("permissions", getResponse.Permissions)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceChatRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	updateInput := &role.UpdateRoleInput{
		Permission: utils.ConvertToStringSlice(d.Get("permissions").([]interface{})),
	}

	updateResp, err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).Update(updateInput)
	if err != nil {
		return fmt.Errorf("Failed to update chat role: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceChatRoleRead(d, meta)
}

func resourceChatRoleDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Chat

	if err := client.Service(d.Get("service_sid").(string)).Role(d.Id()).Delete(); err != nil {
		return fmt.Errorf("Failed to delete chat role: %s", err.Error())
	}
	d.SetId("")
	return nil
}
