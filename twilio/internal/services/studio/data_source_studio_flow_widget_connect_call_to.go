package studio

import (
	"context"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/studio/properties"
	"github.com/RJPearson94/twilio-sdk-go/studio/widgets"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceStudioFlowWidgetConnectCallTo() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetConnectCallToRead,

		Schema: map[string]*schema.Schema{
			"json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A JSON string representation of the widget state, for use as an entry in the `states` list of a `twilio_studio_flow_definition` data source",
			},
			"transitions": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The next widget(s) to transition to after this widget",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"call_completed": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the outbound call completes",
						},
						"hangup": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the next widget when the caller hangs up",
						},
					},
				},
			},
			"offset": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "The position of this widget in the Studio visual editor",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The x-axis position. Defaults to 0",
						},
						"y": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     0,
							Description: "The y-axis position. Defaults to 0",
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "The unique name of this widget within the flow, used to reference it in transitions",
			},
			"noun": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"client",
					"conference",
					"number",
					"number-multi",
					"sim",
					"sip",
				}, false),
				Description: "The TwiML noun that determines the call destination type. Valid values: `client`, `conference`, `number`, `number-multi`, `sim`, `sip`",
			},
			"caller_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{{contact.channel.address}}",
				Description: "The caller ID to display on the outbound call. Defaults to `{{contact.channel.address}}`",
			},
			"record": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to record the outbound call",
			},
			"sip_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SIP URI to dial when `noun` is `sip` (e.g. `sip:user@domain.com`)",
			},
			"sip_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The password for SIP authentication. Sensitive — will not be shown in logs or plans",
			},
			"sip_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The username for SIP authentication",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of seconds to wait for the call to be answered before timing out",
			},
			"to": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The destination to call. Format depends on `noun` (e.g. a phone number in E.164 format, a client name, or a conference name)",
			},
		},
	}
}

func dataSourceStudioFlowWidgetConnectCallToRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.ConnectCallToNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Hangup = utils.OptionalString(d, "transitions.0.hangup")
		nextTransitions.CallCompleted = utils.OptionalString(d, "transitions.0.call_completed")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	widget := &widgets.ConnectCallTo{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.ConnectCallToProperties{
			CallerID:    d.Get("caller_id").(string),
			Noun:        d.Get("noun").(string),
			Offset:      offset,
			Record:      utils.OptionalBool(d, "record"),
			SipEndpoint: utils.OptionalString(d, "sip_endpoint"),
			SipPassword: utils.OptionalString(d, "sip_password"),
			SipUsername: utils.OptionalString(d, "sip_username"),
			Timeout:     utils.OptionalInt(d, "timeout"),
			To:          utils.OptionalString(d, "to"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Connect call to widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create connect call to widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal connect call to widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
