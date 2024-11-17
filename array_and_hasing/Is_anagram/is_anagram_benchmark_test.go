package main

import (
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

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
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

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsAnagramOptimal(tt.input)
			}
			b.ReportMetric(float64(len(tt.input.S1)), "input_length")
		})
	}
}
