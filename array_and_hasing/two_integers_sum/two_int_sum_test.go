package main

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input TwoIntergersInput
		want  [2]int
	}{}

	for _, testCase := range testCases {
		converted := ConvertTestCase(testCase)
		tests = append(tests, converted)
	}
	
	for counter, tt := range tests {
		t.Run(fmt.Sprintf("Test: %d", counter), func(t *testing.T) {
			got := TwoSumBruteForce(tt.input)
			var x1 = got[0] != tt.want[0] || got[1] != tt.want[1]
			var y1 = got[0] != tt.want[1] || got[1] != tt.want[0]

			if x1 && y1 {
				t.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v", tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
			}

		})
	}
}
