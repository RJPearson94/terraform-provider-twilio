package helper

import (
	"time"

	"github.com/RJPearson94/terraform-provider-twilio/twilio/utils"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/configuration/notification"
	"github.com/RJPearson94/twilio-sdk-go/service/conversations/v1/service/conversation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func FlattenTimers(d *schema.ResourceData, timers conversation.FetchConversationTimersResponse) *[]interface{} {
	timerMap := make(map[string]interface{})

	if _, ok := d.GetOk("timers"); ok {
		timerMap["closed"] = utils.OptionalStringWithEmptyStringOnChange(d, "timers.0.closed")
		timerMap["inactive"] = utils.OptionalStringWithEmptyStringOnChange(d, "timers.0.inactive")
	}

	if timers.DateClosed != nil {
		timerMap["date_closed"] = timers.DateClosed.Format(time.RFC3339)
	}

	if timers.DateInactive != nil {
		timerMap["date_inactive"] = timers.DateInactive.Format(time.RFC3339)
	}

	return &[]interface{}{
		timerMap,
	}
}

func FlattenNotificationsAction(input notification.FetchNotificationConversationActionResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"enabled":  input.Enabled,
			"template": input.Template,
			"sound":    input.Sound,
		},
	}
}

func FlattenNotificationsNewMessage(input notification.FetchNotificationNewMessageResponse) *[]interface{} {
	return &[]interface{}{
		map[string]interface{}{
			"enabled":             input.Enabled,
			"template":            input.Template,
			"sound":               input.Sound,
			"badge_count_enabled": input.BadgeCountEnabled,
		},
	}
}
