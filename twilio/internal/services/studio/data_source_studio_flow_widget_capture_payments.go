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

func dataSourceStudioFlowWidgetCapturePayments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStudioFlowWidgetCapturePaymentsRead,

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
						"hangup": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"max_failed_attempts": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"pay_interrupted": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"provider_error": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"success": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"validation_error": {
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
			"payment_token_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"one-time",
					"reusable",
				}, false),
			},
			"payment_connector": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_amount": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"language": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"min_postal_code_length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_attempts": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_code": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"currency": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"postal_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"payment_method": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ACH_DEBIT",
					"CREDIT_CARD",
				}, false),
			},
			"bank_account_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"COMMERCIAL_CHECKING",
					"COMMERCIAL_SAVINGS",
					"CONSUMER_CHECKING",
					"CONSUMER_SAVINGS",
				}, false),
			},
			"valid_card_types": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"amex",
						"diners-club",
						"discover",
						"enroute",
						"jcb",
						"maestro",
						"master-card",
						"optima",
						"visa",
					}, false),
				},
			},
			"parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"value": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},
		},
	}
}

func dataSourceStudioFlowWidgetCapturePaymentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)

	nextTransitions := widgets.CapturePaymentsNextTransitions{}
	if _, ok := d.GetOk("transitions"); ok {
		nextTransitions.Hangup = utils.OptionalString(d, "transitions.0.hangup")
		nextTransitions.MaxFailedAttempts = utils.OptionalString(d, "transitions.0.max_failed_attempts")
		nextTransitions.PayInterrupted = utils.OptionalString(d, "transitions.0.pay_interrupted")
		nextTransitions.ProviderError = utils.OptionalString(d, "transitions.0.provider_error")
		nextTransitions.Success = utils.OptionalString(d, "transitions.0.success")
		nextTransitions.ValidationError = utils.OptionalString(d, "transitions.0.validation_error")
	}

	var offset *properties.Offset
	if _, ok := d.GetOk("offset"); ok {
		offset = &properties.Offset{
			X: d.Get("offset.0.x").(int),
			Y: d.Get("offset.0.y").(int),
		}
	}

	var paymentParameters *[]widgets.CapturePaymentsParameter
	if v, ok := d.GetOk("parameters"); ok {
		parameters := []widgets.CapturePaymentsParameter{}

		for _, parameter := range v.([]interface{}) {
			parameterMap := parameter.(map[string]interface{})
			parameters = append(parameters, widgets.CapturePaymentsParameter{
				Key:   parameterMap["key"].(string),
				Value: parameterMap["value"].(string),
			})
		}
		paymentParameters = &parameters
	}

	widget := &widgets.CapturePayments{
		Name:            name,
		NextTransitions: nextTransitions,
		Properties: widgets.CapturePaymentsProperties{
			Currency:            utils.OptionalString(d, "currency"),
			Description:         utils.OptionalString(d, "description"),
			Language:            utils.OptionalString(d, "language"),
			MaxAttempts:         utils.OptionalInt(d, "max_attempts"),
			MinPostalCodeLength: utils.OptionalInt(d, "min_postal_code_length"),
			Offset:              offset,
			Parameters:          paymentParameters,
			PaymentAmount:       utils.OptionalString(d, "payment_amount"),
			PaymentConnector:    utils.OptionalString(d, "payment_connector"),
			PaymentMethod:       utils.OptionalString(d, "payment_method"),
			PaymentTokenType:    utils.OptionalString(d, "payment_token_type"),
			PostalCode:          utils.OptionalString(d, "postal_code"),
			SecurityCode:        utils.OptionalBool(d, "security_code"),
			Timeout:             utils.OptionalInt(d, "timeout"),
			ValidCardTypes:      utils.OptionalStringSlice(d, "valid_card_types"),
		},
	}

	if err := widget.Validate(); err != nil {
		return diag.Errorf("Capture payments widget failed validation: %s", err.Error())
	}

	state, err := widget.ToState()
	if err != nil {
		return diag.Errorf("Failed to create capture payments widget: %s", err.Error())
	}

	json, jsonErr := state.ToString()
	if jsonErr != nil {
		return diag.Errorf("Failed to marshal capture payments widget to JSON: %s", jsonErr.Error())
	}

	d.SetId(name)
	d.Set("json", json)

	return nil
}
