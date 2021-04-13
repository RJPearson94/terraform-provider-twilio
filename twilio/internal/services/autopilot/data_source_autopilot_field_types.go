package autopilot

import (
	"context"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAutopilotFieldTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAutopilotFieldTypesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.AutopilotAssistantSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_name": {
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

func dataSourceAutopilotFieldTypesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Autopilot

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).FieldTypes.NewFieldTypesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("No field types were found for assistant with sid (%s)", assistantSid)
		}
		return diag.Errorf("Failed to list autopilot field types: %s", err.Error())
	}

	d.SetId(assistantSid)
	d.Set("assistant_sid", assistantSid)

	fieldTypes := make([]interface{}, 0)

	for _, fieldType := range paginator.FieldTypes {
		d.Set("account_sid", fieldType.AccountSid)

		fieldTypeMap := make(map[string]interface{})

		fieldTypeMap["sid"] = fieldType.Sid
		fieldTypeMap["unique_name"] = fieldType.UniqueName
		fieldTypeMap["friendly_name"] = fieldType.FriendlyName
		fieldTypeMap["date_created"] = fieldType.DateCreated.Format(time.RFC3339)

		if fieldType.DateUpdated != nil {
			fieldTypeMap["date_updated"] = fieldType.DateUpdated.Format(time.RFC3339)
		}

		fieldTypeMap["url"] = fieldType.URL

		fieldTypes = append(fieldTypes, fieldTypeMap)
	}

	d.Set("field_types", &fieldTypes)

	return nil
}
