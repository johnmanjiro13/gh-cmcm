package cmcm

import (
	"errors"
	"fmt"
)

var ErrInvalidOutputFormat = errors.New("invalid output format")

type OutputFormat string

const (
	Plain OutputFormat = "plain"
	JSON  OutputFormat = "json"
)

func (o OutputFormat) String() string {
	return string(o)
}

func (o OutputFormat) Valid() error {
	switch o {
	case Plain, JSON:
		return nil
	default:
		return fmt.Errorf("%w: got %s", ErrInvalidOutputFormat, o)
	}
}

func ParseOutputFormat(s string) (OutputFormat, error) {
	o := OutputFormat(s)
	if err := o.Valid(); err != nil {
		return "", err
	}
	return o, nil
}
