package serverless

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceServerlessAssets() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceServerlessAssetsRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_sid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"account_sid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assets": &schema.Schema{
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
						"latest_version_sid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"visibility": &schema.Schema{
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

func dataSourceServerlessAssetsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*common.TwilioClient).Serverless
	ctx, cancel := context.WithTimeout(meta.(*common.TwilioClient).StopContext, d.Timeout(schema.TimeoutRead))
	defer cancel()

	serviceSid := d.Get("service_sid").(string)
	paginator := client.Service(serviceSid).Assets.NewAssetsPaginator()
	for paginator.NextWithContext(ctx) {
	}

	err := paginator.Error()
	if err != nil {
		if utils.IsNotFoundError(err) {
			return fmt.Errorf("[ERROR] No assets were found for serverless service with sid (%s)", serviceSid)
		}
		return fmt.Errorf("[ERROR] Failed to read serverless asset: %s", err.Error())
	}

	d.SetId(serviceSid)
	d.Set("service_sid", serviceSid)

	assets := make([]interface{}, 0)

	for _, asset := range paginator.Assets {
		d.Set("account_sid", asset.AccountSid)

		assetMap := make(map[string]interface{})

		assetMap["sid"] = asset.Sid
		assetMap["friendly_name"] = asset.FriendlyName
		assetMap["date_created"] = asset.DateCreated.Format(time.RFC3339)

		if asset.DateUpdated != nil {
			assetMap["date_updated"] = asset.DateUpdated.Format(time.RFC3339)
		}

		assetMap["url"] = asset.URL

		versionsPaginator := client.Service(serviceSid).Asset(asset.Sid).Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
			PageSize: sdkUtils.Int(5),
		})
		// The twilio api return the latest version as the first element in the array.
		// So there is no need to loop to retrieve all records
		versionsPaginator.Next()

		if versionsPaginator.Error() != nil {
			return fmt.Errorf("[ERROR] Failed to read serverless asset versions: %s", versionsPaginator.Error().Error())
		}

		if len(versionsPaginator.Versions) > 0 {
			latestVersion := versionsPaginator.Versions[0]

			assetMap["latest_version_sid"] = latestVersion.Sid
			assetMap["path"] = latestVersion.Path
			assetMap["visibility"] = latestVersion.Visibility
		} else {
			log.Printf("[INFO] No serverless asset versions found for asset (%s)", asset.Sid)
		}

		assets = append(assets, assetMap)
	}

	d.Set("assets", &assets)

	return nil
}
