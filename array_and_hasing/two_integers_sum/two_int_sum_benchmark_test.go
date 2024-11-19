package main

import (
	"fmt"
	"testing"
)

func BenchmarkTwoSumBruteForce(b *testing.B) {
	var tests = []struct {
		input TwoIntergersInput
		want  [2]int
	}{}

	for _, testCase := range testCases {
		converted := ConvertTestCase(testCase)
		tests = append(tests, converted)
	}
	b.ResetTimer()

	b.Run(fmt.Sprintf("TestCase"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumBruteForce(tt.input)
				var x1 = got[0] != tt.want[0] || got[1] != tt.want[1]
				var y1 = got[0] != tt.want[1] || got[1] != tt.want[0]

				if x1 && y1 {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	for n := 0; n < b.N; n++ {

	}
	b.ReportAllocs()

}
