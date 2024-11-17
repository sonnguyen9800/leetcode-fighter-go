package main

import (
	"testing"
)

func BenchmarkIsAnagram(b *testing.B) {
	tests := []struct {
		name  string
		input InputString
	}{
		{"short strings", InputString{"listen", "silent"}},
		{"medium strings", InputString{"anagram", "nagaram"}},
		{"long strings", InputString{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}},
		{"different lengths", InputString{"hello", "worlds"}},
		{"empty strings", InputString{"", ""}},
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
	tests := []struct {
		name  string
		input InputString
	}{
		{"short strings", InputString{"listen", "silent"}},
		{"medium strings", InputString{"anagram", "nagaram"}},
		{"long strings", InputString{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"}},
		{"different lengths", InputString{"hello", "worlds"}},
		{"empty strings", InputString{"", ""}},
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
