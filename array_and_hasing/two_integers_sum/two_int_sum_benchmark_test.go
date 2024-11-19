package main

import (
	"fmt"
	"testing"
)

var TEST_CASES_NUMBER = 100
var MIN_ARRAY_SIZE = 100
var MAX_ARRAY_SIZE = 100
var MIN_NUMBER = 1
var MAX_NUMBER = 10000

func BenchmarkTwoSum(b *testing.B) {
	var tests = []struct {
		input TwoIntergersInput
		want  [2]int
	}{}

	all_test_case := GenerateFuzzyTestCases(TEST_CASES_NUMBER,
		[2]int{MIN_ARRAY_SIZE, MAX_ARRAY_SIZE},
		[2]int{MIN_NUMBER, MAX_NUMBER})

	for _, testCase := range all_test_case {
		converted := ConvertTestCase(testCase)
		tests = append(tests, converted)
	}
	b.ResetTimer()

	b.Run(fmt.Sprintf("Bruteforce"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumBruteForce(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	b.ResetTimer()

	b.Run(fmt.Sprintf("OnePassV1"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumOnePass(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	b.ResetTimer()

	b.Run(fmt.Sprintf("OnePassV2"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumOnePassVer2(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	b.ResetTimer()

	b.Run(fmt.Sprintf("OnePassV3"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumOnePassVer3(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	b.ResetTimer()

	b.Run(fmt.Sprintf("OnePassV4"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumOnePassVer4(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
	b.ResetTimer()

	b.Run(fmt.Sprintf("OnePassV5"), func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, tt := range tests {
				got := TwoSumOnePassNeetcode(tt.input)
				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]

				if sum_expected != sum_wanted {
					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
				}
			}

		}
	})
}

//func BenchmarkTwoSumOnePass(b *testing.B) {
//	var tests = []struct {
//		input TwoIntergersInput
//		want  [2]int
//	}{}
//	all_test_case := GenerateFuzzyTestCases(TEST_CASES_NUMBER,
//		[2]int{MIN_ARRAY_SIZE, MAX_ARRAY_SIZE},
//		[2]int{MIN_NUMBER, MAX_NUMBER})
//
//	for _, testCase := range all_test_case {
//		converted := ConvertTestCase(testCase)
//		tests = append(tests, converted)
//	}
//	b.ResetTimer()
//
//	b.Run(fmt.Sprintf("OnePassV2"), func(b *testing.B) {
//		for n := 0; n < b.N; n++ {
//			for _, tt := range tests {
//				got := TwoSumOnePass(tt.input)
//				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
//				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]
//
//				if sum_expected != sum_wanted {
//					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
//						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
//				}
//			}
//
//		}
//	})
//
//}
//
//func BenchmarkTwoSumOnePassVer2(b *testing.B) {
//	var tests = []struct {
//		input TwoIntergersInput
//		want  [2]int
//	}{}
//	all_test_case := GenerateFuzzyTestCases(TEST_CASES_NUMBER,
//		[2]int{MIN_ARRAY_SIZE, MAX_ARRAY_SIZE},
//		[2]int{MIN_NUMBER, MAX_NUMBER})
//
//	for _, testCase := range all_test_case {
//		converted := ConvertTestCase(testCase)
//		tests = append(tests, converted)
//	}
//	b.ResetTimer()
//
//	b.Run(fmt.Sprintf("Size %v", len(tests)), func(b *testing.B) {
//		for n := 0; n < b.N; n++ {
//			for _, tt := range tests {
//				got := TwoSumOnePassVer2(tt.input)
//				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
//				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]
//
//				if sum_expected != sum_wanted {
//					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
//						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
//				}
//			}
//
//		}
//	})
//
//}
//func BenchmarkTwoSumOnePassVer3(b *testing.B) {
//	var tests = []struct {
//		input TwoIntergersInput
//		want  [2]int
//	}{}
//	all_test_case := GenerateFuzzyTestCases(TEST_CASES_NUMBER,
//		[2]int{MIN_ARRAY_SIZE, MAX_ARRAY_SIZE},
//		[2]int{MIN_NUMBER, MAX_NUMBER})
//
//	for _, testCase := range all_test_case {
//		converted := ConvertTestCase(testCase)
//		tests = append(tests, converted)
//	}
//	b.ResetTimer()
//
//	b.Run(fmt.Sprintf("Size %v", len(tests)), func(b *testing.B) {
//		for n := 0; n < b.N; n++ {
//			for _, tt := range tests {
//				got := TwoSumOnePassVer3(tt.input)
//				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
//				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]
//
//				if sum_expected != sum_wanted {
//					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
//						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
//				}
//			}
//
//		}
//	})
//
//	b.Run(fmt.Sprintf("OnePassV4"), func(b *testing.B) {
//		for n := 0; n < b.N; n++ {
//			for _, tt := range tests {
//				got := TwoSumOnePassVer4(tt.input)
//				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
//				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]
//
//				if sum_expected != sum_wanted {
//					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
//						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
//				}
//			}
//
//		}
//	})
//
//}
//func BenchmarkTwoSumOnePassVer4(b *testing.B) {
//	var tests = []struct {
//		input TwoIntergersInput
//		want  [2]int
//	}{}
//	all_test_case := GenerateFuzzyTestCases(TEST_CASES_NUMBER,
//		[2]int{MIN_ARRAY_SIZE, MAX_ARRAY_SIZE},
//		[2]int{MIN_NUMBER, MAX_NUMBER})
//
//	for _, testCase := range all_test_case {
//		converted := ConvertTestCase(testCase)
//		tests = append(tests, converted)
//	}
//	b.ResetTimer()
//
//	b.Run(fmt.Sprintf("Size %v", len(tests)), func(b *testing.B) {
//		for n := 0; n < b.N; n++ {
//			for _, tt := range tests {
//				got := TwoSumOnePassVer4(tt.input)
//				var sum_expected = tt.input.ArrayInt[got[0]] + tt.input.ArrayInt[got[1]]
//				var sum_wanted = tt.input.ArrayInt[tt.want[0]] + tt.input.ArrayInt[tt.want[1]]
//
//				if sum_expected != sum_wanted {
//					b.Errorf("Array: %v, Target: %v Wanted: %v and %v, Got: %v %v",
//						tt.input.ArrayInt, tt.input.Target, tt.want[0], tt.want[1], got[0], got[1])
//				}
//			}
//
//		}
//	})
//
//}
