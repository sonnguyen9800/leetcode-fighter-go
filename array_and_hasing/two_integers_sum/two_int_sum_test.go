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
		t.Run(fmt.Sprintf("BruteForce: %d", counter), func(t *testing.T) {
			got := TwoSumBruteForce(tt.input)
			var x1 = got[0] != tt.want[0] || got[1] != tt.want[1]
			var y1 = got[0] != tt.want[1] || got[1] != tt.want[0]

			if x1 && y1 {
				t.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v", tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
			}

		})
	}
}
func TestTwoSumOnePass(t *testing.T) {
	var tests = []struct {
		input TwoIntergersInput
		want  [2]int
	}{}

	var x = TwoIntegersSumTestCase{Input: TwoIntergersInput{ArrayInt: []int{3, 2, 4}, Target: 6}, Expected: [2]int{1, 2}}

	//for _, testCase := range testCases {
	//	converted := ConvertTestCase(testCase)
	//	tests = append(tests, converted)
	//}

	converted := ConvertTestCase(x)
	tests = append(tests, converted)

	for counter, tt := range tests {
		t.Run(fmt.Sprintf("OnePass: %d", counter), func(t *testing.T) {
			got := TwoSumOnePass(tt.input)
			var x1 = got[0] != tt.want[0] || got[1] != tt.want[1]
			var y1 = got[0] != tt.want[1] || got[1] != tt.want[0]

			if x1 && y1 {
				t.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v", tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
			}

		})
	}
}
