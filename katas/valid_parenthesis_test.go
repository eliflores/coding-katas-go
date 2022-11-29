package katas

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s     string
		valid bool
	}{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"{[]}",
			true,
		},
	}

	for _, test := range tests {
		valid := isValid(test.s)

		if valid != test.valid {
			t.Errorf("length = %t; want %t", valid, test.valid)
		}
	}
}
