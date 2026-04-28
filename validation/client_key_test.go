package validation

import "testing"

func TestIsValidClientKey(t *testing.T) {
	cases := []struct {
		name string
		key  string
		want bool
	}{
		{name: "valid", key: "pf_abc123", want: true},
		{name: "missing prefix", key: "abc123", want: false},
		{name: "prefix only", key: "pf_", want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsValidClientKey(tc.key); got != tc.want {
				t.Fatalf("IsValidClientKey(%q) = %v, want %v", tc.key, got, tc.want)
			}
		})
	}
}
