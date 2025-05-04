package utils

import (
	"testing"
)

func Test_BuildString(t *testing.T) {
	for _, test := range []struct {
		name      string
		in        []string
		separator string
		expected  string
	}{
		{
			in:        []string{"a", "b", "c"},
			separator: " ",
			expected:  "a b c",
		},
		{
			in:        []string{"a", "b", "c"},
			separator: ": ",
			expected:  "a: b: c",
		},
		{
			separator: ": ",
			expected:  "",
		},
	} {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if got, want := BuildString(test.separator, test.in...), test.expected; got != want {
				t.Fatalf("got: %s, want: %s", got, want)
			}
		})
	}
}
