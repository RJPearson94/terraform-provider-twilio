package autopilot

import (
	"context"
	"fmt"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAutopilotFieldTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAutopilotFieldTypesRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"assistant_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_types": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"friendly_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"unique_name": &schema.Schema{
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
				},
			},
		},
	}
}

func dataSourceAutopilotFieldTypesRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Autopilot
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	assistantSid := d.Get("assistant_sid").(string)
	paginator := client.Assistant(assistantSid).FieldTypes.NewFieldTypesPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No field types were found for assistant with sid (%s)", assistantSid)
		}
		return fmt.Errorf("[ERROR] Failed to list autopilot field types: %s", err.Error())
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
