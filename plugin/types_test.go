package plugin

import (
	"reflect"
	"testing"
)

func TestSplitWithEscaping(t *testing.T) {
	tests := []struct {
		Input  string
		Output []string
	}{
		{"", []string{}},
		{"a,b", []string{"a", "b"}},
		{",,,", []string{"", "", "", ""}},
		{",a\\,", []string{"", "a,"}},
		{"a,b\\,c\\\\d,e", []string{"a", "b,c\\\\d", "e"}},
	}

	for _, test := range tests {
		strings := splitWithEscaping(test.Input, ",", "\\")
		got, want := strings, test.Output

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got tag %v, want %v", got, want)
		}
	}
}
