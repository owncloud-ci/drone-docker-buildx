package plugin

import (
	"strings"
)

// StringSliceFlag is a flag type which support comma separated values and escaping to not split at unwanted lines.
type StringSliceFlag struct {
	slice []string
}

func (s *StringSliceFlag) String() string {
	return strings.Join(s.slice, " ")
}

func (s *StringSliceFlag) Set(value string) error {
	s.slice = splitWithEscaping(value, ",", "\\")

	return nil
}

func (s *StringSliceFlag) Get() []string {
	return s.slice
}

func splitWithEscaping(in, separator, escapeString string) []string {
	if len(in) == 0 {
		return []string{}
	}

	out := strings.Split(in, separator)

	//nolint:gomnd
	for i := len(out) - 2; i >= 0; i-- {
		if strings.HasSuffix(out[i], escapeString) {
			out[i] = out[i][:len(out[i])-len(escapeString)] + separator + out[i+1]
			out = append(out[:i+1], out[i+2:]...)
		}
	}

	return out
}
