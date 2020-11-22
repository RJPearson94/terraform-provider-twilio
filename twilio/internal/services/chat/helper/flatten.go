package helper

import "github.com/RJPearson94/twilio-sdk-go/service/chat/v2/service"

func FlattenNotifications(input service.FetchServiceNotificationsResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"added_to_channel":     flattenNotificationsAction(input.AddedToChannel),
			"invited_to_channel":   flattenNotificationsAction(input.InvitedToChannel),
			"removed_from_channel": flattenNotificationsAction(input.RemovedFromChannel),
			"new_message":          flattenNewMesage(input.NewMessage),
			"log_enabled":          input.LogEnabled,
		},
	}
}

func FlattenLimits(input service.FetchServiceLimitsResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"user_channels":   input.UserChannels,
			"channel_members": input.ChannelMembers,
		},
	}
}

func FlattenMedia(input service.FetchServiceMediaResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"compatibility_message": input.CompatibilityMessage,
			"size_limit_mb":         input.SizeLimitMB,
		},
	}
}

func flattenNotificationsAction(input service.FetchServiceNotificationsActionResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"enabled":  input.Enabled,
			"template": input.Template,
			"sound":    input.Sound,
		},
	}
}

func flattenNewMesage(input service.FetchServiceNotificationsNewMessageResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"enabled":             input.Enabled,
			"template":            input.Template,
			"sound":               input.Sound,
			"badge_count_enabled": input.BadgeCountEnabled,
		},
	}
}
