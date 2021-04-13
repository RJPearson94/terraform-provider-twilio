package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotFieldValues() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotFieldValuesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"field_type_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotFieldTypeSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"language": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"synonym_of": {
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
				},
			},
		},
	}
}

func dataSourceAutopilotFieldValuesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	fieldTypeSid := d.Get("field_type_sid").(string)
	paginator := client.Assistant(assistantSid).FieldType(fieldTypeSid).FieldValues.NewFieldValuesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No field values were found for assistant with sid (%s) and field type with sid (%s)", assistantSid, fieldTypeSid)
		}
		return diag.Errorf("Failed to list autopilot field values: %s", err.Error())
	}

	d.SetId(fieldTypeSid + "/" + fieldTypeSid)
	d.Set("assistant_sid", assistantSid)
	d.Set("field_type_sid", fieldTypeSid)

	values := make([]interface{}, 0)

	for _, value := range paginator.FieldValues {
		d.Set("account_sid", value.AccountSid)

		valueMap := make(map[string]interface{})

		valueMap["sid"] = value.Sid
		valueMap["language"] = value.Language
		valueMap["value"] = value.Value
		valueMap["synonym_of"] = value.SynonymOf
		valueMap["date_created"] = value.DateCreated.Format(time.RFC3339)

		if value.DateUpdated != nil {
			valueMap["date_updated"] = value.DateUpdated.Format(time.RFC3339)
		}

		valueMap["url"] = value.URL

		values = append(values, valueMap)
	}

	d.Set("field_values", &values)

	return nil
}
