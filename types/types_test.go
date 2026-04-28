package types

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestConnectMessageJSONShape(t *testing.T) {
	msg := ConnectMessage{
		Type:       MessageTypeRegisterAck,
		RequestID:  "req-123",
		AppName:    "web",
		PublicPort: 13000,
		Approved:   true,
		UserName:   "alice-smith",
		Message:    "ok",
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	got := string(data)
	for _, want := range []string{
		`"type":"register-ack"`,
		`"request_id":"req-123"`,
		`"app_name":"web"`,
		`"public_port":13000`,
		`"approved":true`,
		`"user_name":"alice-smith"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("marshaled JSON %s does not contain %s", got, want)
		}
	}
}

func TestAppRegistrationRoundTrip(t *testing.T) {
	now := time.Date(2026, 4, 22, 6, 0, 0, 0, time.UTC)
	in := AppRegistration{
		AppName:        "web",
		TargetURL:      "http://127.0.0.1:3000",
		PublicPort:     13000,
		Approved:       true,
		Source:         "manual",
		DiscoveredPort: 3000,
		LastSeenAt:     now,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	data, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	var out AppRegistration
	if err := json.Unmarshal(data, &out); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	if out.AppName != in.AppName || out.TargetURL != in.TargetURL || out.PublicPort != in.PublicPort || out.Source != in.Source {
		t.Fatalf("round trip mismatch: got %+v want %+v", out, in)
	}
}

func TestMessageTypeConstants(t *testing.T) {
	cases := map[string]string{
		"register":     MessageTypeRegister,
		"register-ack": MessageTypeRegisterAck,
		"request":      MessageTypeRequest,
		"response":     MessageTypeResponse,
		"error":        MessageTypeError,
	}

	for want, got := range cases {
		if got != want {
			t.Fatalf("constant mismatch: got %q want %q", got, want)
		}
	}
}

