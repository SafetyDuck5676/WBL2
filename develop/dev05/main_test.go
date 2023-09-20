package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	type TestCase struct {
		SearchTerm string
		Options    Options
		Input      string
		Expected   string
	}

	testCases := []TestCase{
		{
			SearchTerm: "example",
			Options:    Options{},
			Input: `This is an example file.
It contains multiple lines.
The keyword "example" is present in this line.`,
			Expected: `This is an example file.
The keyword "example" is present in this line.`,
		},
		{
			SearchTerm: "example",
			Options: Options{
				LineNum: true,
			},
			Input: `This is an example file.
It contains multiple lines.
The keyword "example" is present in this line.`,
			Expected: `1:This is an example file.
3:The keyword "example" is present in this line.`,
		},
	}

	for _, testCase := range testCases {
		reader := strings.NewReader(testCase.Input)
		writer := bytes.NewBufferString("")

		grep(testCase.SearchTerm, testCase.Options, reader, writer)
		result := strings.TrimSpace(writer.String())
		expected := strings.TrimSpace(testCase.Expected)

		if result != expected {
			t.Errorf("SearchTerm: %s, Options: %+v\nExpected:\n%s\n\nGot:\n%s\n\n",
				testCase.SearchTerm, testCase.Options, expected, result)
		}
	}
}
