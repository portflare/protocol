package types

import "time"

const (
	// MessageTypeRegister asks the server to register or refresh an app on a client connection.
	MessageTypeRegister = "register"
	// MessageTypeRegisterAck acknowledges a register request.
	MessageTypeRegisterAck = "register-ack"
	// MessageTypeRequest carries a proxied HTTP request from server to client.
	MessageTypeRequest = "request"
	// MessageTypeResponse carries a proxied HTTP response from client to server.
	MessageTypeResponse = "response"
	// MessageTypeError carries a protocol-level error.
	MessageTypeError = "error"
)

// AppRegistration is the shared wire-level representation of an app known to the client.
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

// ConnectMessage is the common websocket message envelope exchanged between client and server.
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

// TunnelRequest is the request shape forwarded from the server to a connected client.
type TunnelRequest struct {
	Type       string              `json:"type"`
	RequestID  string              `json:"request_id,omitempty"`
	AppName    string              `json:"app_name,omitempty"`
	Method     string              `json:"method,omitempty"`
	URL        string              `json:"url,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	BodyBase64 string              `json:"body_base64,omitempty"`
}

// TunnelResponse is the response shape sent from a client back to the server for a proxied request.
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
