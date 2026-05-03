package types

// RegistrationMode identifies how the CLI should complete registration.
type RegistrationMode string

const (
	// RegistrationModeAuth means the CLI should complete the auth-service redirect flow.
	RegistrationModeAuth RegistrationMode = "auth"
	// RegistrationModeDirect means the CLI may use direct server registration fallback.
	RegistrationModeDirect RegistrationMode = "direct"
)

// AuthProvider identifies the external authentication provider that verified a user.
type AuthProvider string

const (
	// AuthProviderGoogle identifies identities authenticated by Google OAuth.
	AuthProviderGoogle AuthProvider = "google"
)

const (
	// CodeChallengeMethodS256 is the only supported PKCE challenge method.
	CodeChallengeMethodS256 = "S256"
)

// RegisterStartRequest starts CLI registration with either auth-service handoff or direct fallback.
type RegisterStartRequest struct {
	RedirectURI         string `json:"redirect_uri"`
	State               string `json:"state"`
	Nonce               string `json:"nonce"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
}

// RegisterStartResponse tells the CLI which registration mode to use.
type RegisterStartResponse struct {
	Mode      RegistrationMode `json:"mode"`
	AuthURL   string           `json:"auth_url,omitempty"`
	ExpiresIn int              `json:"expires_in,omitempty"`
}

// RegisterExchangeRequest exchanges a CLI callback authorization code for a Portflare client key.
type RegisterExchangeRequest struct {
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier"`
	RedirectURI  string `json:"redirect_uri"`
}

// RegistrationResponse is returned when the server has created or resolved a Portflare user/key.
type RegistrationResponse struct {
	UserName        string `json:"user_name"`
	PublicUserLabel string `json:"public_user_label"`
	Email           string `json:"email,omitempty"`
	APIKey          string `json:"api_key"`
}

// AuthIdentityClaims are verified identity claims returned by the trusted auth service.
type AuthIdentityClaims struct {
	Provider          AuthProvider `json:"provider"`
	Subject           string       `json:"subject"`
	Email             string       `json:"email,omitempty"`
	EmailVerified     bool         `json:"email_verified"`
	DisplayName       string       `json:"display_name,omitempty"`
	SuggestedUserName string       `json:"suggested_user_name,omitempty"`
}

// InternalVerifyRequest asks the auth service to verify a one-time authorization code.
type InternalVerifyRequest struct {
	Code         string `json:"code"`
	CodeVerifier string `json:"code_verifier,omitempty"`
	RedirectURI  string `json:"redirect_uri"`
}

// InternalVerifyResponse returns verified identity claims to the Portflare server.
type InternalVerifyResponse struct {
	AuthIdentityClaims
}
