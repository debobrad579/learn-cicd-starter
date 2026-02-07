package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectErr   bool
		expectedErr error
	}{
		{
			name: "valid ApiKey header",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-key"},
			},
			expectedKey: "my-secret-key",
			expectErr:   false,
		},
		{
			name:        "missing Authorization header",
			headers:     http.Header{},
			expectErr:   true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - wrong scheme",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-key"},
			},
			expectErr: true,
		},
		{
			name: "malformed header - no key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectErr: true,
		},
		{
			name: "malformed header - empty value",
			headers: http.Header{
				"Authorization": []string{""},
			},
			expectErr:   true,
			expectedErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("did not expect error, got %v", err)
			}

			if key != tt.expectedKey {
				t.Fatalf("expected key %q, got %q", tt.expectedKey, key)
			}
		})
	}
}
