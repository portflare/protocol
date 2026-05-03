package validation

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

const (
	// CodeChallengeMethodS256 is the only supported PKCE challenge method.
	CodeChallengeMethodS256 = "S256"
	minPKCEVerifierLength   = 43
	maxPKCEVerifierLength   = 128
	s256ChallengeLength     = 43
)

// AuthIdentityClaimsInput is the minimal identity-claim shape accepted by validation helpers.
type AuthIdentityClaimsInput struct {
	Provider      string
	Subject       string
	Email         string
	EmailVerified bool
}

// IsValidPKCECodeVerifier reports whether verifier matches RFC 7636 verifier constraints.
func IsValidPKCECodeVerifier(verifier string) bool {
	if len(verifier) < minPKCEVerifierLength || len(verifier) > maxPKCEVerifierLength {
		return false
	}
	for _, r := range verifier {
		if r > unicode.MaxASCII {
			return false
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			continue
		}
		switch r {
		case '-', '.', '_', '~':
			continue
		default:
			return false
		}
	}
	return true
}

// S256CodeChallenge computes the base64url-encoded SHA-256 PKCE challenge for verifier.
func S256CodeChallenge(verifier string) (string, error) {
	if !IsValidPKCECodeVerifier(verifier) {
		return "", errors.New("invalid PKCE code verifier")
	}
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:]), nil
}

// IsValidS256CodeChallenge reports whether challenge has the expected S256 base64url shape.
func IsValidS256CodeChallenge(challenge string) bool {
	if len(challenge) != s256ChallengeLength {
		return false
	}
	for _, r := range challenge {
		if r > unicode.MaxASCII {
			return false
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			continue
		}
		return false
	}
	return true
}

// IsLoopbackCLIRedirectURI reports whether raw is a safe localhost callback URI for CLI handoff.
func IsLoopbackCLIRedirectURI(raw string) bool {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "http" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
		return false
	}
	if u.Path != "/callback" {
		return false
	}
	host := u.Hostname()
	port := u.Port()
	if host == "" || port == "" {
		return false
	}
	portNum, err := strconv.Atoi(port)
	if err != nil || portNum < 1 || portNum > 65535 {
		return false
	}
	if host == "localhost" {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

// IsSafeBrowserReturnPath reports whether path is a same-origin relative browser return path.
func IsSafeBrowserReturnPath(path string) bool {
	if path == "" || !strings.HasPrefix(path, "/") || strings.HasPrefix(path, "//") {
		return false
	}
	if strings.Contains(path, "\\") || strings.ContainsAny(path, "\r\n\t") {
		return false
	}
	u, err := url.Parse(path)
	if err != nil {
		return false
	}
	if strings.Contains(u.Path, "\\") || strings.ContainsAny(u.Path, "\r\n\t") {
		return false
	}
	return u.Scheme == "" && u.Host == "" && strings.HasPrefix(u.Path, "/")
}

// IsValidAuthIdentityClaims reports whether claims contain the minimum trusted identity data.
func IsValidAuthIdentityClaims(claims AuthIdentityClaimsInput) bool {
	if strings.TrimSpace(claims.Provider) == "" || strings.TrimSpace(claims.Subject) == "" {
		return false
	}
	if strings.TrimSpace(claims.Email) == "" || !claims.EmailVerified {
		return false
	}
	email := strings.TrimSpace(claims.Email)
	addr, err := mail.ParseAddress(email)
	return err == nil && addr.Address == email
}
