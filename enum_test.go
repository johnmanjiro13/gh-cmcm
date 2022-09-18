package cmcm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputFormat_String(t *testing.T) {
	tests := map[string]struct {
		format OutputFormat
		want   string
	}{
		"plain": {Plain, "plain"},
		"json":  {JSON, "json"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.format.String())
		})
	}
}

func TestOutputFormat_Valid(t *testing.T) {
	tests := map[string]struct {
		format OutputFormat
		want   error
	}{
		"plain":   {Plain, nil},
		"json":    {JSON, nil},
		"invalid": {OutputFormat("invalid"), ErrInvalidOutputFormat},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if tt.want != nil {
				assert.ErrorIs(t, tt.format.Valid(), tt.want)
			} else {
				assert.NoError(t, tt.format.Valid())
			}
		})
	}
}

func TestParseOutputFormat(t *testing.T) {
	tests := map[string]struct {
		s       string
		want    OutputFormat
		wantErr error
	}{
		"plain":   {"plain", Plain, nil},
		"json":    {"json", JSON, nil},
		"invalid": {"invalid", "", ErrInvalidOutputFormat},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ParseOutputFormat(tt.s)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
