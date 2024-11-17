package main

import "fmt"

type InputString struct {
	S1, S2 string
}
type IsAnagramTestCase struct {
	S1, S2   string
	Expected bool
}

func ConvertTestCase(testCase IsAnagramTestCase) struct {
	name  string
	input InputString
	want  bool
} {

	return struct {
		name  string
		input InputString
		want  bool
	}{
		name: fmt.Sprintf("Input %s and %s", testCase.S1, testCase.S2),
		input: InputString{
			S1: testCase.S1,
			S2: testCase.S2,
		},
		want: testCase.Expected}

}
