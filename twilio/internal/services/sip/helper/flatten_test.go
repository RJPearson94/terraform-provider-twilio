package helper

import (
	"testing"

	sdkUtils "github.com/RJPearson94/twilio-sdk-go/utils"

	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domain"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/domains"
)

func TestFlattenVoice(t *testing.T) {
	url := "https://example.com/voice"
	method := "POST"

	resp := &domain.FetchDomainResponse{
		VoiceURL:    &url,
		VoiceMethod: &method,
	}

	result := FlattenVoice(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	if got := m["url"].(*string); *got != url {
		t.Errorf("url: got %q, want %q", *got, url)
	}
	if got := m["method"].(*string); *got != method {
		t.Errorf("method: got %q, want %q", *got, method)
	}
}

func TestFlattenVoice_NilFields(t *testing.T) {
	resp := &domain.FetchDomainResponse{}

	result := FlattenVoice(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	if m["url"] != (*string)(nil) {
		t.Errorf("url: expected nil *string, got %v", m["url"])
	}
	if m["method"] != (*string)(nil) {
		t.Errorf("method: expected nil *string, got %v", m["method"])
	}
}

func TestFlattenEmergency(t *testing.T) {
	callerSid := "PNaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resp := &domain.FetchDomainResponse{
		EmergencyCallerSid:      &callerSid,
		EmergencyCallingEnabled: true,
	}

	result := FlattenEmergency(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	if got := m["caller_sid"].(*string); *got != callerSid {
		t.Errorf("caller_sid: got %q, want %q", *got, callerSid)
	}
	if got := m["calling_enabled"].(bool); !got {
		t.Errorf("calling_enabled: got false, want true")
	}
}

func TestFlattenVoiceFromCreate(t *testing.T) {
	url := "https://example.com/voice"
	method := "POST"
	fallbackURL := "https://example.com/fallback"
	fallbackMethod := "GET"
	statusCallbackURL := "https://example.com/status"
	statusCallbackMethod := "POST"

	resp := &domains.CreateDomainResponse{
		VoiceURL:                  &url,
		VoiceMethod:               &method,
		VoiceFallbackURL:          &fallbackURL,
		VoiceFallbackMethod:       &fallbackMethod,
		VoiceStatusCallbackURL:    &statusCallbackURL,
		VoiceStatusCallbackMethod: &statusCallbackMethod,
	}

	result := FlattenVoiceFromCreate(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})

	cases := []struct {
		key  string
		want string
	}{
		{"url", url},
		{"method", method},
		{"fallback_url", fallbackURL},
		{"fallback_method", fallbackMethod},
		{"status_callback_url", statusCallbackURL},
		{"status_callback_method", statusCallbackMethod},
	}
	for _, c := range cases {
		got := m[c.key].(*string)
		if *got != c.want {
			t.Errorf("%s: got %q, want %q", c.key, *got, c.want)
		}
	}
}

func TestFlattenVoiceFromCreate_NilFields(t *testing.T) {
	resp := &domains.CreateDomainResponse{}

	result := FlattenVoiceFromCreate(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	for _, key := range []string{"url", "method", "fallback_url", "fallback_method", "status_callback_url", "status_callback_method"} {
		if m[key] != (*string)(nil) {
			t.Errorf("%s: expected nil *string, got %v", key, m[key])
		}
	}
}

func TestFlattenEmergencyFromCreate(t *testing.T) {
	callerSid := "PNaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	resp := &domains.CreateDomainResponse{
		EmergencyCallerSid:      sdkUtils.String(callerSid),
		EmergencyCallingEnabled: true,
	}

	result := FlattenEmergencyFromCreate(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	if got := m["caller_sid"].(*string); *got != callerSid {
		t.Errorf("caller_sid: got %q, want %q", *got, callerSid)
	}
	if got := m["calling_enabled"].(bool); !got {
		t.Errorf("calling_enabled: got false, want true")
	}
}

func TestFlattenEmergencyFromCreate_DisabledWithNilCallerSid(t *testing.T) {
	resp := &domains.CreateDomainResponse{
		EmergencyCallerSid:      nil,
		EmergencyCallingEnabled: false,
	}

	result := FlattenEmergencyFromCreate(resp)

	if result == nil {
		t.Fatal("expected non-nil result")
	}
	items := *result
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	m := items[0].(map[string]interface{})
	if m["caller_sid"] != (*string)(nil) {
		t.Errorf("caller_sid: expected nil *string, got %v", m["caller_sid"])
	}
	if got := m["calling_enabled"].(bool); got {
		t.Errorf("calling_enabled: got true, want false")
	}
}
