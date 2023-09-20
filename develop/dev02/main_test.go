package main

import "testing"

func TestUnpackString(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"a9b3", "aaaaaaaaabbb"},
		{"a2b2c", "aabbc"},
		{"", ""},
	}

	for _, test := range tests {
		result, _ := unpackString(test.input)
		if result != test.expectedOutput {
			t.Errorf("Expected unpackString(%q) to be %q, but got %q", test.input, test.expectedOutput, result)
		}
	}
}
