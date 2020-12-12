package sip_trunking

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/internal/services/sip_trunking/helper"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunk/recording"
	"github.com/RJPearson94/twilio-sdk-go/service/trunking/v1/trunks"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSIPTrunkingTrunk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSIPTrunkingTrunkCreate,
		ReadContext:   resourceSIPTrunkingTrunkRead,
		UpdateContext: resourceSIPTrunkingTrunkUpdate,
		DeleteContext: resourceSIPTrunkingTrunkDelete,

		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				format := "/Trunks/(.*)"
				regex := regexp.MustCompile(format)
				match := regex.FindStringSubmatch(d.Id())

				if len(match) != 2 {
					return nil, fmt.Errorf("The imported ID (%s) does not match the format (%s)", d.Id(), format)
				}

				d.Set("sid", match[1])
				d.SetId(match[1])
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
			"cnam_lookup_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"disaster_recovery_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"GET",
					"POST",
				}, false),
			},
			"disaster_recovery_url": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recording": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"do-not-record",
								"record-from-ringing",
								"record-from-answer",
								"record-from-ringing-dual",
								"record-from-answer-dual",
							}, false),
						},
						"trim": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ValidateFunc: validation.StringInSlice([]string{
								"trim-silence",
								"do-not-trim",
							}, false),
						},
					},
				},
			},
			"secure": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"transfer_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type_set": {
				Type:     schema.TypeList,
				Computed: true,
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

func resourceSIPTrunkingTrunkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	createInput := &trunks.CreateTrunkInput{
		CnamLookupEnabled:      utils.OptionalBool(d, "cnam_lookup_enabled"),
		DisasterRecoveryMethod: utils.OptionalString(d, "disaster_recovery_method"),
		DisasterRecoveryURL:    utils.OptionalString(d, "disaster_recovery_url"),
		DomainName:             utils.OptionalString(d, "domain_name"),
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		Secure:                 utils.OptionalBool(d, "secure"),
		TransferMode:           utils.OptionalString(d, "transfer_mode"),
	}

	createResult, err := client.Trunks.CreateWithContext(ctx, createInput)
	if err != nil {
		return diag.Errorf("Failed to create SIP trunk: %s", err.Error())
	}

	d.SetId(createResult.Sid)

	if err := updateRecording(ctx, d, meta); err != nil {
		return err
	}

	return resourceSIPTrunkingTrunkRead(ctx, d, meta)
}

func resourceSIPTrunkingTrunkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	getResponse, err := client.Trunk(d.Id()).FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return diag.Errorf("Failed to read SIP trunk: %s", err.Error())
	}

	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("cnam_lookup_enabled", getResponse.CnamLookupEnabled)
	d.Set("disaster_recovery_method", getResponse.DisasterRecoveryMethod)
	d.Set("disaster_recovery_url", getResponse.DisasterRecoveryURL)
	d.Set("domain_name", getResponse.DomainName)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("recording", helper.FlattenRecording(getResponse.Recording))
	d.Set("secure", getResponse.Secure)
	d.Set("transfer_mode", getResponse.TransferMode)
	d.Set("auth_type", getResponse.AuthType)
	d.Set("auth_type_set", getResponse.AuthTypeSet)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	return nil
}

func resourceSIPTrunkingTrunkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	updateInput := &trunk.UpdateTrunkInput{
		CnamLookupEnabled:      utils.OptionalBool(d, "cnam_lookup_enabled"),
		DisasterRecoveryMethod: utils.OptionalString(d, "disaster_recovery_method"),
		DisasterRecoveryURL:    utils.OptionalString(d, "disaster_recovery_url"),
		DomainName:             utils.OptionalString(d, "domain_name"),
		FriendlyName:           utils.OptionalString(d, "friendly_name"),
		Secure:                 utils.OptionalBool(d, "secure"),
		TransferMode:           utils.OptionalString(d, "transfer_mode"),
	}

	updateResp, err := client.Trunk(d.Id()).UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP trunk: %s", err.Error())
	}

	d.SetId(updateResp.Sid)

	if err := updateRecording(ctx, d, meta); err != nil {
		return err
	}
	return resourceSIPTrunkingTrunkRead(ctx, d, meta)
}

func resourceSIPTrunkingTrunkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).SIPTrunking

	if err := client.Trunk(d.Id()).DeleteWithContext(ctx); err != nil {
		return diag.Errorf("Failed to delete SIP trunk: %s", err.Error())
	}
	d.SetId("")
	return nil
}

func updateRecording(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if !d.HasChanges("recording", "recording.0.trim", "recording.0.mode") {
		return nil
	}

	client := meta.(*common.TwilioClient).SIPTrunking

	updateInput := &recording.UpdateRecordingInput{
		Trim: utils.OptionalString(d, "recording.0.trim"),
		Mode: utils.OptionalString(d, "recording.0.mode"),
	}

	_, err := client.Trunk(d.Id()).Recording().UpdateWithContext(ctx, updateInput)
	if err != nil {
		return diag.Errorf("Failed to update SIP trunk recording: %s", err.Error())
	}

	return nil
}
