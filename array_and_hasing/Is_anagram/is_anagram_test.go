package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
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
			got := IsAnagram(tt.input)
			if got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
