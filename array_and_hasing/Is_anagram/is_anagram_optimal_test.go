package main

import "testing"

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

	println("Test Length: ", len(tests))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IsAnagramOptimal(tt.input)
			//if got != tt.want {
			//	t.Errorf("Result = %v, want %v", got, tt.want)
			//	t.Errorf("IsAnagramOptimal(%v) = %v; want %v", tt.input, got, tt.want)
			//}
		})
	}
}
