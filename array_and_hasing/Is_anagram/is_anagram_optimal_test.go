package main

import (
	"fmt"
	"testing"
)

func TestIsAnagramOptimal(t *testing.T) {
	var tests = []struct {
		name  string
		input InputString
		want  bool
	}{}

	for _, testCase := range testCases {
		converted := ConvertTestCase(testCase)
		tests = append(tests, converted)
	}

	for counter, tt := range tests {
		t.Run(fmt.Sprintf("Test: %d", counter), func(t *testing.T) {
			got := IsAnagramOptimal(tt.input)
			if got != tt.want {
				t.Errorf("IsAnagramOptimal(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
