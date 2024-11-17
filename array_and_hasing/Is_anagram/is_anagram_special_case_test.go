package main

import (
	"fmt"
	"testing"
)

func TestIsAnagramOptimal_SpecialCase(t *testing.T) {
	var tests = []struct {
		name  string
		input InputString
		want  bool
	}{}

	for _, testCase := range testCases {
		converted := ConvertTestCase(testCase)
		tests = append(tests, converted)
	}

	//var testSpecialCase = IsAnagramTestCase{S1: string(make([]byte, 10000)), S2: string(make([]byte, 10000)), Expected: true} // All-zero long string
	//var testSpecialCase2 = IsAnagramTestCase{S1: "theeyes", S2: "theysee", Expected: true}                                    // All-zero long string
	//
	//var converted = ConvertTestCase(testSpecialCase)
	//var converted2 = ConvertTestCase(testSpecialCase2)
	//tests = append(tests, converted)
	//tests = append(tests, converted2)

	println("Test Length: ", len(tests))
	//for _, testCase := range tests {
	//	t.Run(testCase.name, func(t *testing.T) {
	//		got := IsAnagramOptimal(testCase.input)
	//		if got != testCase.want {
	//			t.Errorf("Abs(-1) = %t; want %t", got, testCase.want)
	//		} else {
	//			println("Hello???")
	//		}
	//	})
	//}
	for i := 0; i < 100; i++ {
		t.Run("Test", func(t *testing.T) {

		})
	}
	for c, _ := range tests {
		t.Run(fmt.Sprintf("%d", c), func(t *testing.T) {

		})
	}
}
