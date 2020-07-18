package utils

import (
	"strings"

	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
)

func OptionalString(d *schema.ResourceData, key string) *string {
	if !d.HasChange(key) {
		return nil
	}
	return sdkUtils.String(d.Get(key).(string))
}

func OptionalJSONString(d *schema.ResourceData, key string) *string {
	if !d.HasChange(key) {
		return nil
	}

	// error not handled as it is assumed stringIsJSON validation is applied to the resource
	normalizedJSON, _ := structure.NormalizeJsonString(d.Get(key).(string))
	return sdkUtils.String(normalizedJSON)
}

func OptionalSeperatedString(d *schema.ResourceData, key string, seperator string) *string {
	if !d.HasChange(key) {
		return nil
	}

	retrievedKey := d.Get(key)
	if retrievedKey == nil {
		return nil
	}

	stringSlice := ConvertToStringSlice(retrievedKey.([]interface{}))
	return sdkUtils.String(strings.Join(stringSlice[:], seperator))
}

func OptionalStringSlice(d *schema.ResourceData, key string) *[]string {
	if !d.HasChange(key) {
		return nil
	}

	retrievedKey := d.Get(key)
	if retrievedKey == nil {
		return nil
	}

	stringSlice := ConvertToStringSlice(retrievedKey.([]interface{}))
	return &stringSlice
}

func OptionalInt(d *schema.ResourceData, key string) *int {
	if !d.HasChange(key) {
		return nil
	}
	return sdkUtils.Int(d.Get(key).(int))
}

func OptionalBool(d *schema.ResourceData, key string) *bool {
	if !d.HasChange(key) {
		return nil
	}
	return sdkUtils.Bool(d.Get(key).(bool))
}
