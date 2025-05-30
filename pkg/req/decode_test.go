package req_test

import (
	"io"
	"strings"
	"testing"

	"quoteList/internal/payload"
	"quoteList/pkg/req"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantQuote payload.Quote
		wantErr   bool
	}{
		{
			name:      "valid json",
			input:     `{"quote":"Hello","author":"World"}`,
			wantQuote: payload.Quote{Quote: "Hello", Author: "World"},
			wantErr:   false,
		},
		{
			name:      "invalid json",
			input:     `{"quote":"Hello","author":}`,
			wantQuote: payload.Quote{},
			wantErr:   true,
		},
		{
			name:      "empty body",
			input:     ``,
			wantQuote: payload.Quote{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := io.NopCloser(strings.NewReader(tt.input))

			got, err := req.Decode(body)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.wantQuote {
				t.Errorf("Decode() = %+v, want %+v", got, tt.wantQuote)
			}
		})
	}
}
