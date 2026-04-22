package types

import "time"

const (
	MessageTypeRegister    = "register"
	MessageTypeRegisterAck = "register-ack"
	MessageTypeRequest     = "request"
	MessageTypeResponse    = "response"
	MessageTypeError       = "error"
)

// AppRegistration is a shared wire-level representation of a registered app.
type AppRegistration struct {
	AppName        string    `json:"app_name"`
	TargetURL      string    `json:"target_url"`
	PublicPort     int       `json:"public_port,omitempty"`
	Approved       bool      `json:"approved"`
	Source         string    `json:"source,omitempty"`
	DiscoveredPort int       `json:"discovered_port,omitempty"`
	Offline        bool      `json:"offline,omitempty"`
	LastSeenAt     time.Time `json:"last_seen_at,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// ConnectMessage is sent between the client and server over the websocket tunnel.
type ConnectMessage struct {
	Type       string              `json:"type"`
	RequestID  string              `json:"request_id,omitempty"`
	AppName    string              `json:"app_name,omitempty"`
	PublicPort int                 `json:"public_port,omitempty"`
	Method     string              `json:"method,omitempty"`
	URL        string              `json:"url,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	BodyBase64 string              `json:"body_base64,omitempty"`
	StatusCode int                 `json:"status_code,omitempty"`
	Error      string              `json:"error,omitempty"`
	Approved   bool                `json:"approved,omitempty"`
	UserName   string              `json:"user_name,omitempty"`
	Message    string              `json:"message,omitempty"`
}

// TunnelRequest is the server-side request shape forwarded to a connected client.
type TunnelRequest struct {
	Type       string              `json:"type"`
	RequestID  string              `json:"request_id,omitempty"`
	AppName    string              `json:"app_name,omitempty"`
	Method     string              `json:"method,omitempty"`
	URL        string              `json:"url,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	BodyBase64 string              `json:"body_base64,omitempty"`
}

// TunnelResponse is the client-to-server response shape for proxied requests.
type TunnelResponse struct {
	Type       string              `json:"type"`
	RequestID  string              `json:"request_id,omitempty"`
	StatusCode int                 `json:"status_code,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	BodyBase64 string              `json:"body_base64,omitempty"`
	Error      string              `json:"error,omitempty"`
	AppName    string              `json:"app_name,omitempty"`
	PublicPort int                 `json:"public_port,omitempty"`
	Approved   bool                `json:"approved,omitempty"`
	UserName   string              `json:"user_name,omitempty"`
	Message    string              `json:"message,omitempty"`
}
