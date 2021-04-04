package serverless

import (
	"context"
	"log"
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/common"
	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/serverless/v1/service/asset/versions"
	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceServerlessAsset() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServerlessAssetRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessAssetSidValidation(),
			},
			"service_sid": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: utils.ServerlessServiceSidValidation(),
			},
			"account_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"friendly_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_version_sid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visibility": {
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
	}
}

func dataSourceServerlessAssetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*common.TwilioClient).Serverless

	serviceSid := d.Get("service_sid").(string)
	sid := d.Get("sid").(string)
	assetClient := client.Service(serviceSid).Asset(sid)

	getResponse, err := assetClient.FetchWithContext(ctx)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return diag.Errorf("Asset with sid (%s) was not found for serverless service with sid (%s)", sid, serviceSid)
		}
		return diag.Errorf("Failed to read serverless asset: %s", err.Error())
	}

	d.SetId(getResponse.Sid)
	d.Set("sid", getResponse.Sid)
	d.Set("account_sid", getResponse.AccountSid)
	d.Set("service_sid", getResponse.ServiceSid)
	d.Set("friendly_name", getResponse.FriendlyName)
	d.Set("date_created", getResponse.DateCreated.Format(time.RFC3339))

	if getResponse.DateUpdated != nil {
		d.Set("date_updated", getResponse.DateUpdated.Format(time.RFC3339))
	}

	d.Set("url", getResponse.URL)

	versionsPaginator := assetClient.Versions.NewVersionsPaginatorWithOptions(&versions.VersionsPageOptions{
		PageSize: sdkUtils.Int(5),
	})
	// The twilio api return the latest version as the first element in the array.
	// So there is no need to loop to retrieve all records
	versionsPaginator.Next()

	if versionsPaginator.Error() != nil {
		return diag.Errorf("Failed to read serverless asset versions: %s", versionsPaginator.Error().Error())
	}

	if len(versionsPaginator.Versions) > 0 {
		latestVersion := versionsPaginator.Versions[0]

		d.Set("latest_version_sid", latestVersion.Sid)
		d.Set("path", latestVersion.Path)
		d.Set("visibility", latestVersion.Visibility)
	} else {
		log.Printf("[INFO] No serverless asset versions found for asset (%s)", getResponse.Sid)
	}

	return nil
}
