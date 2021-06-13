package voice

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queue"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/queues"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVoiceQueue() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVoiceQueueCreate,
		ReadContext:   resourceVoiceQueueRead,
		UpdateContext: resourceVoiceQueueUpdate,
		DeleteContext: resourceVoiceQueueDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Accounts/(.*)/Queues/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 3 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("account_sid", match[1])
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: utils.AccountSidValidation(),
			},
			"friendly_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
			},
			"max_size": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      100,
				ValidateFunc: validation.IntBetween(1, 5000),
			},
			"average_wait_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"current_size": {
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
		},
	}
}

func resourceVoiceQueueCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	createInput := &queues.CreateQueueInput{
		FriendlyName: d.Get("friendly_name").(string),
		MaxSize:      utils.OptionalInt(d, "max_size"),
	}

	createResult, err := client.Account(d.Get("account_sid").(string)).Queues.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create queue: %s", err.Error())
	}

	d.SetId(createResult.Sid)
	return resourceVoiceQueueRead(ctx, d, meta)
}

func resourceVoiceQueueRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	getResponse, err := client.Account(d.Get("account_sid").(string)).Queue(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		// If the account sid is incorrect a 401 is returned, a this is a generic error this will not be handled here and an error will be returned
		return diag.Errorf("Failed to read queue: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("max_size", getResponse.MaxSize)
	d.Set("average_wait_time", getResponse.AverageWaitTime)
	d.Set("current_size", getResponse.CurrentSize)
	d.Set("date_created", getResponse.DateCreated.Time.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Time.Format(time.RFC3339))
	}

	return nil
}

func resourceVoiceQueueUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	updateInput := &queue.UpdateQueueInput{
		FriendlyName: utils.OptionalString(d, "friendly_name"),
		MaxSize:      utils.OptionalInt(d, "max_size"),
	}

	updateResp, err := client.Account(d.Get("account_sid").(string)).Queue(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update queue: %s", err.Error())
	}

	d.SetId(updateResp.Sid)
	return resourceVoiceQueueRead(ctx, d, meta)
}

func resourceVoiceQueueDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).API

	if err := client.Account(d.Get("account_sid").(string)).Queue(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete queue: %s", err.Error())
	}

	d.SetId("")
	return nil
}
