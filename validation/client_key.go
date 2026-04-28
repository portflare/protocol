package validation

import "strings"

// ClientKeyPrefix is the required prefix for all Portflare client API keys.
const ClientKeyPrefix = "pf_"

// IsValidClientKey reports whether v matches the minimal shared client key format.
func IsValidClientKey(v string) bool {
	return strings.HasPrefix(v, ClientKeyPrefix) && len(v) > len(ClientKeyPrefix)
}
