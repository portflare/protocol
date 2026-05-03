package types

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAuthRegistrationTypesJSONShape(t *testing.T) {
	start := RegisterStartRequest{
		RedirectURI:         "http://127.0.0.1:49152/callback",
		State:               "state-123",
		Nonce:               "nonce-123",
		CodeChallenge:       "challenge-123",
		CodeChallengeMethod: CodeChallengeMethodS256,
	}
	data, err := json.Marshal(start)
	if err != nil {
		t.Fatalf("Marshal(RegisterStartRequest) error = %v", err)
	}
	got := string(data)
	for _, want := range []string{
		`"redirect_uri":"http://127.0.0.1:49152/callback"`,
		`"state":"state-123"`,
		`"nonce":"nonce-123"`,
		`"code_challenge":"challenge-123"`,
		`"code_challenge_method":"S256"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("RegisterStartRequest JSON %s does not contain %s", got, want)
		}
	}

	response := RegisterStartResponse{Mode: RegistrationModeAuth, AuthURL: "https://auth.example.test/cli/start", ExpiresIn: 300}
	data, err = json.Marshal(response)
	if err != nil {
		t.Fatalf("Marshal(RegisterStartResponse) error = %v", err)
	}
	got = string(data)
	for _, want := range []string{`"mode":"auth"`, `"auth_url":"https://auth.example.test/cli/start"`, `"expires_in":300`} {
		if !strings.Contains(got, want) {
			t.Fatalf("RegisterStartResponse JSON %s does not contain %s", got, want)
		}
	}
}

func TestRegisterExchangeAndRegistrationResponseJSONShape(t *testing.T) {
	exchange := RegisterExchangeRequest{Code: "code-123", CodeVerifier: "verifier-123", RedirectURI: "http://127.0.0.1:49152/callback"}
	data, err := json.Marshal(exchange)
	if err != nil {
		t.Fatalf("Marshal(RegisterExchangeRequest) error = %v", err)
	}
	got := string(data)
	for _, want := range []string{`"code":"code-123"`, `"code_verifier":"verifier-123"`, `"redirect_uri":"http://127.0.0.1:49152/callback"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("RegisterExchangeRequest JSON %s does not contain %s", got, want)
		}
	}

	registration := RegistrationResponse{UserName: "alice", PublicUserLabel: "alice", Email: "alice@example.test", APIKey: "pf_secret"}
	data, err = json.Marshal(registration)
	if err != nil {
		t.Fatalf("Marshal(RegistrationResponse) error = %v", err)
	}
	got = string(data)
	for _, want := range []string{`"user_name":"alice"`, `"public_user_label":"alice"`, `"email":"alice@example.test"`, `"api_key":"pf_secret"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("RegistrationResponse JSON %s does not contain %s", got, want)
		}
	}
}

func TestAuthIdentityClaimsAndInternalVerifyJSONShape(t *testing.T) {
	claims := AuthIdentityClaims{
		Provider:          AuthProviderGoogle,
		Subject:           "google-subject",
		Email:             "alice@example.test",
		EmailVerified:     true,
		DisplayName:       "Alice Smith",
		SuggestedUserName: "alice",
	}
	verify := InternalVerifyResponse{AuthIdentityClaims: claims}
	data, err := json.Marshal(verify)
	if err != nil {
		t.Fatalf("Marshal(InternalVerifyResponse) error = %v", err)
	}
	got := string(data)
	for _, want := range []string{
		`"provider":"google"`,
		`"subject":"google-subject"`,
		`"email":"alice@example.test"`,
		`"email_verified":true`,
		`"display_name":"Alice Smith"`,
		`"suggested_user_name":"alice"`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("InternalVerifyResponse JSON %s does not contain %s", got, want)
		}
	}

	request := InternalVerifyRequest{Code: "code-123", CodeVerifier: "verifier-123", RedirectURI: "http://127.0.0.1:49152/callback"}
	data, err = json.Marshal(request)
	if err != nil {
		t.Fatalf("Marshal(InternalVerifyRequest) error = %v", err)
	}
	got = string(data)
	for _, want := range []string{`"code":"code-123"`, `"code_verifier":"verifier-123"`, `"redirect_uri":"http://127.0.0.1:49152/callback"`} {
		if !strings.Contains(got, want) {
			t.Fatalf("InternalVerifyRequest JSON %s does not contain %s", got, want)
		}
	}
}

func TestAuthConstants(t *testing.T) {
	if RegistrationModeAuth != "auth" || RegistrationModeDirect != "direct" {
		t.Fatalf("unexpected registration modes: auth=%q direct=%q", RegistrationModeAuth, RegistrationModeDirect)
	}
	if AuthProviderGoogle != "google" {
		t.Fatalf("unexpected google provider constant: %q", AuthProviderGoogle)
	}
	if CodeChallengeMethodS256 != "S256" {
		t.Fatalf("unexpected code challenge method: %q", CodeChallengeMethodS256)
	}
}
