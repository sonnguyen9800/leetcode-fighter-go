package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name  string
		input InputString
		want  bool
	}{
		{"same words", InputString{"listen", "silent"}, true},
		{"case insensitive", InputString{"Listen", "Silent"}, true},
		{"different lengths", InputString{"listen", "silently"}, false},
		{"non-anagrams", InputString{"hello", "world"}, false},
		{"empty strings", InputString{"", ""}, true},
		{"single characters", InputString{"a", "A"}, true},
		{"special characters", InputString{"@b#c", "c#b@"}, true},
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
	tests := []struct {
		name  string
		input InputString
		want  bool
	}{
		{"same words", InputString{"listen", "silent"}, true},
		{"case insensitive", InputString{"Listen", "Silent"}, true},
		{"different lengths", InputString{"listen", "silently"}, false},
		{"non-anagrams", InputString{"hello", "world"}, false},
		{"empty strings", InputString{"", ""}, true},
		{"single characters", InputString{"a", "A"}, true},
		{"special characters", InputString{"@b#c", "c#b@"}, true},
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
