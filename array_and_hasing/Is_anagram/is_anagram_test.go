package main

import (
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.input)
			if got != tt.want {
				t.Errorf("IsAnagram(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}

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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagramOptimal(tt.input)
			if got != tt.want {
				t.Errorf("IsAnagramOptimal(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
