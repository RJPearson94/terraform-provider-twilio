package helper

// IsVoiceReceiveMode is used since Programmable fax has been disabled on some accounts voice receive mode is no
// longer being returned, so receive mode is not returned it is assumed voice is configured
func IsVoiceReceiveMode(receiveMode *string) bool {
	return receiveMode == nil || *receiveMode == "voice"
}
