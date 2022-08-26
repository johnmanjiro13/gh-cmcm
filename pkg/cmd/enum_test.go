package cmd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnmanjiro13/gh-cmcm/pkg/cmd"
)

func TestOutputFormat_String(t *testing.T) {
	tests := map[string]struct {
		format cmd.OutputFormat
		want   string
	}{
		"plain": {cmd.Plain, "plain"},
		"json":  {cmd.JSON, "json"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.format.String())
		})
	}
}

func TestOutputFormat_Valid(t *testing.T) {
	tests := map[string]struct {
		format cmd.OutputFormat
		want   error
	}{
		"plain":   {cmd.Plain, nil},
		"json":    {cmd.JSON, nil},
		"invalid": {cmd.OutputFormat("invalid"), cmd.ErrInvalidOutputFormat},
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
		want    cmd.OutputFormat
		wantErr error
	}{
		"plain":   {"plain", cmd.Plain, nil},
		"json":    {"json", cmd.JSON, nil},
		"invalid": {"invalid", "", cmd.ErrInvalidOutputFormat},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := cmd.ParseOutputFormat(tt.s)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
