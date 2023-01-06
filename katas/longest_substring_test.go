package katas

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s      string
		length int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	for _, test := range tests {
		length := lengthOfLongestSubstring(test.s)

		if length != test.length {
			t.Errorf("length = %d; want %d", length, test.length)
		}
	}
}
