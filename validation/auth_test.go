package validation

import "testing"

func TestPKCEVerifierAndS256ChallengeValidation(t *testing.T) {
	verifier := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~abcde"
	challenge, err := S256CodeChallenge(verifier)
	if err != nil {
		t.Fatalf("S256CodeChallenge() error = %v", err)
	}
	if !IsValidPKCECodeVerifier(verifier) {
		t.Fatalf("expected verifier to be valid")
	}
	if !IsValidS256CodeChallenge(challenge) {
		t.Fatalf("expected generated challenge %q to be valid", challenge)
	}
	if len(challenge) != 43 {
		t.Fatalf("expected S256 challenge length 43, got %d", len(challenge))
	}

	invalidVerifiers := []string{
		"short",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~x",
		"contains space abcdefghijklmnopqrstuvwxyz0123456789abcdef",
		"contains=paddingabcdefghijklmnopqrstuvwxyz0123456789abcdef",
	}
	for _, v := range invalidVerifiers {
		if IsValidPKCECodeVerifier(v) {
			t.Fatalf("expected verifier %q to be invalid", v)
		}
		if _, err := S256CodeChallenge(v); err == nil {
			t.Fatalf("expected S256CodeChallenge(%q) to fail", v)
		}
	}

	for _, challenge := range []string{
		"short",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO=",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNO+",
	} {
		if IsValidS256CodeChallenge(challenge) {
			t.Fatalf("expected challenge %q to be invalid", challenge)
		}
	}
}

func TestLoopbackCLIRedirectURIValidation(t *testing.T) {
	valid := []string{
		"http://127.0.0.1:49152/callback",
		"http://localhost:49152/callback",
		"http://[::1]:49152/callback",
	}
	for _, uri := range valid {
		if !IsLoopbackCLIRedirectURI(uri) {
			t.Fatalf("expected %q to be a valid loopback CLI redirect URI", uri)
		}
	}

	invalid := []string{
		"https://127.0.0.1:49152/callback",
		"http://127.0.0.1/callback",
		"http://127.0.0.1:0/callback",
		"http://127.0.0.1:65536/callback",
		"http://127.0.0.1:49152/other",
		"http://0.0.0.0:49152/callback",
		"http://192.168.1.10:49152/callback",
		"http://evil.test:49152/callback",
		"not a url",
	}
	for _, uri := range invalid {
		if IsLoopbackCLIRedirectURI(uri) {
			t.Fatalf("expected %q to be invalid", uri)
		}
	}
}

func TestSafeBrowserReturnPathValidation(t *testing.T) {
	valid := []string{"/", "/me", "/admin?tab=apps", "/r/alice/web"}
	for _, path := range valid {
		if !IsSafeBrowserReturnPath(path) {
			t.Fatalf("expected %q to be safe", path)
		}
	}

	invalid := []string{
		"",
		"me",
		"//evil.test/path",
		"/\\evil",
		"/%5cevil",
		"https://evil.test/me",
		"/auth/callback\nSet-Cookie:bad=1",
		"/%0d%0aSet-Cookie:bad=1",
	}
	for _, path := range invalid {
		if IsSafeBrowserReturnPath(path) {
			t.Fatalf("expected %q to be unsafe", path)
		}
	}
}

func TestAuthIdentityClaimsValidation(t *testing.T) {
	valid := AuthIdentityClaimsInput{Provider: "google", Subject: "sub-123", Email: "alice@example.test", EmailVerified: true}
	if !IsValidAuthIdentityClaims(valid) {
		t.Fatalf("expected claims to be valid")
	}

	invalid := []AuthIdentityClaimsInput{
		{Provider: "", Subject: "sub-123", Email: "alice@example.test", EmailVerified: true},
		{Provider: "google", Subject: "", Email: "alice@example.test", EmailVerified: true},
		{Provider: "google", Subject: "sub-123", Email: "alice@example.test", EmailVerified: false},
		{Provider: "google", Subject: "sub-123", Email: "not-an-email", EmailVerified: true},
		{Provider: "google", Subject: "sub-123", Email: "Alice <alice@example.test>", EmailVerified: true},
	}
	for _, claims := range invalid {
		if IsValidAuthIdentityClaims(claims) {
			t.Fatalf("expected claims to be invalid: %#v", claims)
		}
	}
}
