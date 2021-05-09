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
				Type:     schema.TypeString,
				Computed: true,
			},
			"transitions": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"call_completed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"hangup": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"offset": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"x": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"y": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
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
			},
			"caller_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "{{contact.channel.address}}",
			},
			"record": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"sip_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sip_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"sip_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"to": {
				Type:     schema.TypeString,
				Optional: true,
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
