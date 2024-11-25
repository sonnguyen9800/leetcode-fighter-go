package is_anagram

import (
	"fmt"
	"testing"
)

func BenchmarkIsAnagram(b *testing.B) {

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
		b.Run(fmt.Sprintf("Test: %d", counter), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsAnagram(tt.input)
			}
			b.ReportMetric(float64(len(tt.input.S1)), "input_length")
		})
	}
}

func BenchmarkIsAnagramOptimal(b *testing.B) {
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
		b.Run(fmt.Sprintf("Test: %d", counter), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsAnagramOptimal(tt.input)
			}
			b.ReportMetric(float64(len(tt.input.S1)), "input_length")
		})
	}
}
