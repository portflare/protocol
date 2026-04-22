package validation

import "strings"

const ClientKeyPrefix = "pf_"

func IsValidClientKey(v string) bool {
	return strings.HasPrefix(v, ClientKeyPrefix) && len(v) > len(ClientKeyPrefix)
}
