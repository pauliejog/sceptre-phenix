package web

import "testing"

func TestNormalizeFileServerEndpoint(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		endpoint string
		wantErr  bool
	}{
		{name: "empty disabled", value: ""},
		{name: "zero disabled", value: "0"},
		{name: "port only", value: "8080", endpoint: "127.0.0.1:8080"},
		{name: "ipv4", value: "127.0.0.1:8080", endpoint: "127.0.0.1:8080"},
		{name: "hostname", value: "localhost:8080", endpoint: "localhost:8080"},
		{name: "all interfaces", value: "0.0.0.0:8080", endpoint: "0.0.0.0:8080"},
		{name: "empty host", value: ":8080", endpoint: ":8080"},
		{name: "ipv6", value: "[::1]:8080", endpoint: "[::1]:8080"},
		{name: "invalid bare host", value: "abc", wantErr: true},
		{name: "missing port", value: "localhost", wantErr: true},
		{name: "invalid port", value: "127.0.0.1:notaport", wantErr: true},
		{name: "port out of range", value: "99999", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpoint, err := normalizeFileServerEndpoint(tt.value)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if endpoint != tt.endpoint {
				t.Fatalf("endpoint = %q, want %q", endpoint, tt.endpoint)
			}
		})
	}
}
